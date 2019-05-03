/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package file

import (
	"os"
	"strings"

	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/extractor/file/internal/sniff"
)

func xInfo(file *os.File, contentType *string) (bool, api.Metadata, error) {
	if strings.HasPrefix(*contentType, "application/") {
		d, err := sniff.File(file)
		if err != nil {
			return false, nil, err
		}
		*contentType = d.ContentType()
		return true, api.Metadata{
			"architecture": strings.ToLower(d.Arch),
			"platform":     d.Platform,
			"file":         d,
		}, nil
	}
	return false, nil, nil
}
