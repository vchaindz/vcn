/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package docker

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/uri"
)

// Scheme for docker
const Scheme = "docker"

// Artifact returns a file *api.Artifact from a given u
func Artifact(u *uri.URI) (*api.Artifact, error) {

	if u.Scheme != Scheme {
		return nil, nil
	}

	id := strings.TrimPrefix(u.Opaque, "//")
	images, err := inspect(id)
	if err != nil {
		return nil, fmt.Errorf("failed to inspect docker image: %s", err)
	}
	if len(images) < 1 {
		return nil, fmt.Errorf("no docker image found for: %s", id)
	}

	i := images[0]
	return &api.Artifact{
		Name: Scheme + "://" + i.name(),
		Hash: i.hash(),
		Size: i.Size,
	}, nil
}

type image struct {
	ID       string   `json:"Id"`
	RepoTags []string `json:"RepoTags"`
	Size     uint64   `json:"Size"`
}

func (i image) hash() string {
	return strings.TrimSpace(strings.Replace(fmt.Sprint(i.ID), "sha256:", "", 1))
}

func (i image) name() string {
	if len(i.RepoTags) > 0 {
		return i.RepoTags[0]
	}
	return i.hash()
}

func inspect(arg string) ([]image, error) {
	cmd := exec.Command("docker", "inspect", arg)
	cmdOutput, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	data := []image{}
	if err = json.Unmarshal(cmdOutput, &data); err != nil {
		return nil, err
	}
	return data, nil
}
