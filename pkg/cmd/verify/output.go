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
	"time"

	"gopkg.in/yaml.v2"

	"github.com/dustin/go-humanize"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/meta"
)

type result struct {
	Artifact     *api.ArtifactResponse       `json:"artifact"`
	Verification *api.BlockchainVerification `json:"verification"`
}

func (r result) WriterTo(out io.Writer) {
	w := new(tabwriter.Writer)
	w.Init(out, 0, 8, 0, '\t', 0)

	if a := r.Artifact; a != nil {

		s := reflect.ValueOf(a).Elem()
		typeOfT := s.Type()

		for i, l := 0, s.NumField(); i < l; i++ {
			f := s.Field(i)
			if key, ok := typeOfT.Field(i).Tag.Lookup("vcn"); ok {
				var value string
				switch true {
				case key == "Size":
					if size, ok := f.Interface().(uint64); ok {
						value = humanize.Bytes(size)
					}
				case key == "Metadata":
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
				default:
					value = fmt.Sprintf("%s", f.Interface())
				}
				if value != "" {
					fmt.Fprintf(w, "%s:\t%s\n", key, value)
				}
			}
		}
	}

	if bv := r.Verification; bv != nil {
		if key := bv.Key(); key != "" {
			fmt.Fprintf(w, "Key:\t%s\n", bv.Key())
		}
		if bv.Level > 0 {
			fmt.Fprintf(w, "Level:\t%s\n", bv.LevelName())
		}
		if bv.Timestamp != time.Unix(0, 0) {
			fmt.Fprintf(w, "Date:\t%s\n", bv.Timestamp.String())
		}
	}

	fmt.Fprintf(w, "Status:\t%s\n", meta.StatusNameStyled(r.Verification.Status))

	w.Flush()
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
		r.WriterTo(os.Stdout)
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
