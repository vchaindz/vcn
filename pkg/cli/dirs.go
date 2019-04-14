/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package cli

import (
	"os"

	"github.com/vchain-us/vcn/pkg/logs"
	"github.com/vchain-us/vcn/pkg/meta"
)

func CreateVcnDirectories() {
	if err := os.MkdirAll(meta.WalletDirectory(),
		os.FileMode(meta.VcnDirectoryPermissions)); err != nil {
		logs.LOG.Fatal(err)
	}
}
