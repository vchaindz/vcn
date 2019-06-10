/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/mattn/go-colorable"

	"gopkg.in/yaml.v2"

	"github.com/dustin/go-humanize"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/meta"
)

type result struct {
	Artifact     *api.ArtifactResponse       `json:"artifact"`
	Verification *api.BlockchainVerification `json:"verification"`
}

func (r result) WriteTo(out io.Writer) (n int64, err error) {
	w := new(tabwriter.Writer)
	w.Init(out, 0, 8, 0, '\t', 0)

	printf := func(format string, a ...interface{}) error {
		m, err := fmt.Fprintf(w, format, a...)
		n += int64(m)
		return err
	}

	if a := r.Artifact; a != nil {

		s := reflect.ValueOf(a).Elem()
		typeOfT := s.Type()

		for i, l := 0, s.NumField(); i < l; i++ {
			f := s.Field(i)
			if key, ok := typeOfT.Field(i).Tag.Lookup("vcn"); ok {
				var value string
				switch key {
				case "Size":
					if size, ok := f.Interface().(uint64); ok && size > 0 {
						value = humanize.Bytes(size)
					}
				case "Metadata":
					if metadata, ok := f.Interface().(api.Metadata); ok {
						for k, v := range metadata {
							if v == "" {
								continue
							}
							if vv, err := json.MarshalIndent(v, "\t", "    "); err == nil {
								value += fmt.Sprintf("\n\t\t%s=%s", k, string(vv))
							}
						}
						value = strings.TrimPrefix(value, "\n")
					}
				case "Signer":
					if f.Interface() != r.Verification.Key() {
						value = fmt.Sprintf("%s", f.Interface())
					}
				default:
					value = fmt.Sprintf("%s", f.Interface())
				}
				if value != "" {
					err = printf("%s:\t%s\n", key, value)
					if err != nil {
						return
					}
				}
			}
		}
	}

	if bv := r.Verification; bv != nil {
		if key := bv.Key(); key != "" {
			err = printf("Key:\t%s\n", bv.Key())
			if err != nil {
				return
			}
		}
		if bv.Level > 0 {
			err = printf("Level:\t%s\n", bv.LevelName())
			if err != nil {
				return
			}
		}
		if date := bv.Date(); date != "" {
			err = printf("Date:\t%s\n", date)
			if err != nil {
				return
			}
		}
		err = printf("Status:\t%s\n", meta.StatusNameStyled(r.Verification.Status))
		if err != nil {
			return
		}
	}

	return n, w.Flush()
}

func Print(output string, a *api.Artifact, artifact *api.ArtifactResponse, verification *api.BlockchainVerification) error {

	r := result{
		Verification: verification,
	}
	if artifact != nil {
		r.Artifact = artifact
	} else if a != nil {
		r.Artifact = &api.ArtifactResponse{
			Name: a.Name,
			Kind: a.Kind,
			Hash: a.Hash,
			Size: a.Size,
		}
	}

	switch output {
	case "":
		r.WriteTo(colorable.NewColorableStdout())
	case "yaml":
		b, err := yaml.Marshal(r)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	case "json":
		b, err := json.MarshalIndent(r, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	default:
		return outputNotSupportedErr(output)
	}
	return nil
}

func PrintList(output string, artifacts []api.ArtifactResponse) error {
	switch output {
	case "":
		for _, a := range artifacts {
			fmt.Print(a)
		}
	case "yaml":
		b, err := yaml.Marshal(artifacts)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	case "json":
		b, err := json.MarshalIndent(artifacts, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	default:
		return outputNotSupportedErr(output)
	}
	return nil
}

func PrintErr(output string, err error) error {
	switch output {
	case "":
		fmt.Printf("Error: %s\n", err)
	case "yaml":
		b, err := yaml.Marshal(map[string]string{
			"error": err.Error(),
		})
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	case "json":
		b, err := json.MarshalIndent(map[string]string{
			"error": err.Error(),
		}, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	default:
		return outputNotSupportedErr(output)
	}
	return nil
}

func outputNotSupportedErr(output string) error {
	return fmt.Errorf("output format not supported: %s", output)
}
