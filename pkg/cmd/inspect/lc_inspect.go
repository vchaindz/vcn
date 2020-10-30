/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package inspect

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/vchain-us/ledger-compliance-go/schema"
	"github.com/vchain-us/vcn/pkg/cmd/internal/cli"
	"github.com/vchain-us/vcn/pkg/meta"
	"google.golang.org/grpc/metadata"

	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/cmd/internal/types"
)

func lcInspect(hash string, signerID string, u *api.LcUser, output string) (err error) {
	hasher := sha256.New()
	hasher.Write([]byte(u.LcApiKey()))
	contextSignerID := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	if signerID == "" {
		if output == "" {
			fmt.Println("no signer ID provided. Full history of the item is returned")
		}
	} else {
		contextSignerID = signerID
	}

	results, err := GetLcResults(hash, signerID, u)

	l := len(results)
	if output == "" {
		fmt.Printf(
			`current signer ID "%s"
%d notarizations found for "%s"

`,
			contextSignerID, l, hash)
	}

	return cli.PrintLcSlice(output, results)
}

func GetLcResults(hash, signerID string, u *api.LcUser) ([]*types.LcResult, error) {
	var err error
	var items *schema.StructuredItemExtList

	md := metadata.Pairs(meta.VcnLCPluginTypeHeaderName, meta.VcnLCPluginTypeHeaderValue)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	if signerID == "" {
		items, err = u.Client.ZScanExt(ctx, []byte(hash))
		if err != nil {
			return nil, err
		}
	} else {
		key := api.AppendPrefix(meta.VcnLCPrefix, []byte(signerID))
		key = api.AppendSignerId(hash, key)
		items, err = u.Client.HistoryExt(ctx, key)
		if err != nil {
			return nil, err
		}
	}

	results := make([]*types.LcResult, len(items.Items))
	var i = 0
	for _, v := range items.Items {
		lca, err := api.ItemToLcArtifact(v)
		if err != nil {
			return nil, err
		}
		results[i] = types.NewLcResult(lca, true)
		if err != nil {
			results[i].AddError(err)
		}
		i++
	}
	return results, nil
}
