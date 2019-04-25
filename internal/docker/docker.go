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
	"strconv"
	"strings"

	"github.com/yalp/jsonpath"
)

func GetHash(descriptor string) (hash string, err error) {
	id := getDockerId(descriptor)
	data, err := inspectDocker(id)
	if err != nil {
		return "", err
	}
	filter, err := jsonpath.Prepare("$..Id")
	if err != nil {
		return "", err
	}
	filteredData, err := filter(data)
	if err != nil {
		return "", err
	}
	dockerHash := filteredData.([]interface{})[0]
	hash = strings.TrimSpace(strings.Replace(fmt.Sprint(dockerHash), "sha256:", "", 1))
	return hash, nil
}

func GetSize(descriptor string) (size int64, err error) {
	id := getDockerId(descriptor)
	data, err := inspectDocker(id)
	if err != nil {
		return 0, err
	}
	filter, err := jsonpath.Prepare("$..Size")
	if err != nil {
		return 0, err
	}
	filteredData, err := filter(data)
	if err != nil {
		return 0, err
	}
	dockerSize := filteredData.([]interface{})[0]
	floatSize, err := strconv.ParseFloat(fmt.Sprint(dockerSize), 64)
	if err != nil {
		return 0, err
	}
	return int64(floatSize), nil
}

func getDockerId(descriptor string) string {
	dockerId := strings.Replace(descriptor, "docker:", "", 1)
	dockerId = strings.Replace(dockerId, " ", "", -1)
	return dockerId
}

func inspectDocker(id string) (payload interface{}, err error) {
	cmd := exec.Command("docker", "inspect", id)
	cmdOutput, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var data interface{}
	if err = json.Unmarshal(cmdOutput, &data); err != nil {
		return nil, err
	}
	return data, nil
}
