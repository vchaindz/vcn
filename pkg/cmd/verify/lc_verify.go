package verify

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/cmd/internal/cli"
	"github.com/vchain-us/vcn/pkg/cmd/internal/types"
	"github.com/vchain-us/vcn/pkg/meta"
	"google.golang.org/grpc/status"
)

func lcVerify(cmd *cobra.Command, a *api.Artifact, user *api.LcUser, signerID string, output string) (err error) {
	hook := newHook(cmd, a)
	err = hook.lcFinalizeWithoutAlert(user, output, 0)
	if err != nil {
		return err
	}
	ar, verified, err := user.LoadArtifact(a.Hash, signerID, 0)
	if err != nil {
		if status.Convert(err).Message() == "key not found" {
			err = fmt.Errorf("%s was not notarized", a.Hash)
		}
		return cli.PrintWarning(output, err.Error())
	}
	if !verified {
		color.Set(meta.StyleError())
		fmt.Println("the ledger is compromised. Please contact the CodeNotary Ledger Compliance administrators")
		color.Unset()
		fmt.Println()
		ar.Status = meta.StatusUnknown
	}

	cli.PrintLc(output, types.NewLcResult(ar, verified))

	return
}
