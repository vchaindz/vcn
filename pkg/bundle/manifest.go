/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package bundle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"

	// See https://github.com/opencontainers/go-digest#usage
	_ "crypto/sha256"
	_ "crypto/sha512"

	digest "github.com/opencontainers/go-digest"
)

const (
	// ManifestSchemaVersion is the current manifest schema version.
	ManifestSchemaVersion = 1

	// ManifestFilename is the default filename for manifest when stored.
	ManifestFilename = ".vcn.manifest.json"
)

// Manifest provides bundle structure when marshalled to JSON.
type Manifest struct {
	// SchemaVersion is the manifest schema that this bundle follows
	SchemaVersion uint `json:"schemaVersion"`

	// Items is an ordered list of items referenced by the manifest.
	Items []Descriptor `json:"items"`
}

// MarshalJSON implements the json.Marshaler interface.
func (m *Manifest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("cannot marshal nil manifest")
	}
	m.Sort()
	type alias Manifest
	mm := alias(*m)
	return json.Marshal(mm)
}

// Sort m's items
func (m *Manifest) Sort() {
	if m == nil {
		return
	}

	// make unique index
	idx := make(map[string]Descriptor, len(m.Items))
	for _, d := range m.Items {
		k := d.Digest.String()
		if dd, ok := idx[k]; ok {
			dd.Paths = append(dd.Paths, d.Paths...)
			idx[k] = dd
		} else {
			idx[k] = d
		}
	}

	// recreate unique digest list and sort paths
	m.Items = make([]Descriptor, len(idx))
	i := 0
	for _, d := range idx {
		d.sortUnique()
		m.Items[i] = d
		i++
	}

	// finally, sort items by digest
	sort.SliceStable(m.Items, func(k, j int) bool {
		return m.Items[k].Digest.String() < m.Items[j].Digest.String()
	})
}

// Digest digests the JSON encoded m and returns a digest.Digest.
func (m *Manifest) Digest() (digest.Digest, error) {
	b, err := json.Marshal(m) // sorting is implicitly called by Marshal
	if err != nil {
		return "", err
	}

	return digest.SHA256.FromBytes(b), nil
}

// NewManifest returns a new empty Manifest.
func NewManifest(items ...Descriptor) *Manifest {
	if items == nil {
		items = make([]Descriptor, 0)
	}
	return &Manifest{
		SchemaVersion: ManifestSchemaVersion,
		Items:         items,
	}
}

// WriteManifest writes manifest's data to a file named by filename.
func WriteManifest(manifest Manifest, filename string) error {
	data, err := json.Marshal(manifest)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// ReadManifest reads the file named by filename and returns the decoded manifest.
func ReadManifest(filename string) (*Manifest, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	m := Manifest{}
	return &m, json.Unmarshal(data, &m)
}
