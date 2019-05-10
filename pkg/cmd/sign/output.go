/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package sign

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/dustin/go-humanize"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/meta"
)

func print(a *api.Artifact, v *api.BlockchainVerification) error {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	printf := func(key string, value interface{}) {
		if value != "" {
			fmt.Fprintf(w, "%s:\t%s\n", key, value)
		}
	}

	printf("Kind", a.Kind)
	printf("Name", a.Name)
	printf("Hash", a.Hash)
	printf("Size", humanize.Bytes(a.Size))
	printf("Date", v.Date())
	printf("Status", meta.StatusNameStyled(v.Status))

	return w.Flush()
}
