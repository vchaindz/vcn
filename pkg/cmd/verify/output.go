/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package verify

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"

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
					if size, ok := f.Interface().(uint64); ok {
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
	}

	err = printf("Status:\t%s\n", meta.StatusNameStyled(r.Verification.Status))
	if err != nil {
		return
	}

	return n, w.Flush()
}

func print(output string, a *api.Artifact, artifact *api.ArtifactResponse, verification *api.BlockchainVerification) error {

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
		r.WriteTo(os.Stdout)
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
		return fmt.Errorf("output format not supported: %s", output)
	}
	return nil
}
