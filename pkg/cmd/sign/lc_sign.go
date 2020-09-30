package sign

import (
	"fmt"
	"github.com/caarlos0/spin"
	"github.com/fatih/color"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/cmd/internal/cli"
	"github.com/vchain-us/vcn/pkg/cmd/internal/types"
	"github.com/vchain-us/vcn/pkg/meta"
)

func LcSign(u *api.LcUser, a api.Artifact, state meta.Status, output string) error {

	if output == "" {
		color.Set(meta.StyleAffordance())
		fmt.Println("Your assets will not be uploaded. They will be processed locally.")
		color.Unset()
		fmt.Println()
	}

	s := spin.New("%s Notarization in progress...")
	s.Set(spin.Spin1)

	_, err := u.Sign(
		a,
		api.LcSignWithStatus(state),
	)

	if output == "" {
		fmt.Println()
	}

	artifact, verified, err := u.LoadArtifact(a.Hash)
	if err != nil {
		return cli.PrintWarning(output, err.Error())
	}

	if !verified {
		color.Set(meta.StyleError())
		fmt.Println("the ledger is compromised. Please contact the CodeNotary Ledger Compliance administrators")
		color.Unset()
		fmt.Println()
	}

	cli.PrintLc(output, types.NewLcResult(artifact))

	return nil
}
