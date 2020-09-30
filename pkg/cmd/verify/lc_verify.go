package verify

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/cmd/internal/cli"
	"github.com/vchain-us/vcn/pkg/cmd/internal/types"
	"github.com/vchain-us/vcn/pkg/meta"
)

func lcVerify(a *api.Artifact, user *api.LcUser, output string) (err error) {

	ar, verified, _ := user.LoadArtifact(a.Hash)

	if !verified {
		color.Set(meta.StyleError())
		fmt.Println("the ledger is compromised. Please contact the CodeNotary Ledger Compliance administrators")
		color.Unset()
		fmt.Println()
	}

	cli.PrintLc(output, types.NewLcResult(ar))

	return
}
