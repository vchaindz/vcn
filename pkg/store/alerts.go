/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package store

import (
	"os"
	"path/filepath"
)

type Alert struct {
	Name   string
	Arg    string
	Config interface{}
}

type Alerts map[string]Alert

func loadAlerts(filepath string) (Alerts, error) {
	a := make(Alerts)
	err := ReadYAML(&a, filepath)
	if os.IsNotExist(err) {
		return a, nil
	}
	return a, err
}

func saveAlerts(alerts Alerts, filepath string) error {
	return WriteYAML(alerts, filepath)
}

func AlertFilepath(email string) (string, error) {
	path := filepath.Join(dir, defaultAlertsDir)
	if err := ensureDir(path); err != nil {
		return "", err
	}
	return filepath.Join(path, email+".yaml"), nil
}

func SaveAlert(userEmail string, alertID string, alert Alert) error {
	path, err := AlertFilepath(userEmail)
	if err != nil {
		return err
	}

	alerts, err := loadAlerts(path)
	if err != nil {
		return err
	}

	alerts[alertID] = alert

	return saveAlerts(alerts, path)
}

func DeleteAlert(userEmail string, id string) error {
	path, err := AlertFilepath(userEmail)
	if err != nil {
		return err
	}

	alerts, err := loadAlerts(path)
	if err != nil {
		return err
	}

	delete(alerts, id)

	return saveAlerts(alerts, path)
}

func ReadAlerts(userEmail string) (Alerts, error) {
	path, err := AlertFilepath(userEmail)
	if err != nil {
		return make(Alerts), err
	}
	return loadAlerts(path)
}
