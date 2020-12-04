/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package serve

import (
	"fmt"
	"github.com/vchain-us/vcn/pkg/api"
	inspect2 "github.com/vchain-us/vcn/pkg/cmd/inspect"
	"net/http"
)

func lcInspect(user *api.LcUser, hash, signerID string, w http.ResponseWriter) {
	if user.Client.ApiKey == "" {
		writeError(w, http.StatusUnauthorized, fmt.Errorf("api key not provided"))
		return
	}
	err := user.Client.Connect()
	if err != nil {
		writeError(w, http.StatusBadGateway, err)
		return
	}

<<<<<<< HEAD
	results, err := inspect2.GetLcResults(hash, signerID, user)
=======
	results, err := inspect2.GetLcResults(hash, signerID, user, 0, 0, "", "")
>>>>>>> origin
	if err != nil {
		writeError(w, http.StatusBadGateway, err)
		return
	}

	writeLcResults(w, http.StatusOK, results)
}
