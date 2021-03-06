package verify

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/cmd/internal/cli"
	"github.com/vchain-us/vcn/pkg/cmd/internal/types"
	"github.com/vchain-us/vcn/pkg/meta"
	"google.golang.org/grpc/status"
	"strconv"
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
			viper.Set("exit-code", strconv.Itoa(meta.StatusUnknown.Int()))
		}
		return cli.PrintWarning(output, err.Error())
	}
	if !verified {
		color.Set(meta.StyleError())
		fmt.Println("the ledger is compromised. Please contact the CodeNotary Ledger Compliance administrators")
		color.Unset()
		fmt.Println()
		viper.Set("exit-code", strconv.Itoa(meta.StatusUnknown.Int()))
		ar.Status = meta.StatusUnknown
	}

	exitCode, err := cmd.Flags().GetInt("exit-code")
	if err != nil {
		return err
	}
	// if exitCode == VcnDefaultExitCode user didn't specify to use a custom exit code in case of success.
	// In that case we return the ar.Status as exit code.
	if exitCode == meta.VcnDefaultExitCode {
		viper.Set("exit-code", strconv.Itoa(ar.Status.Int()))
	}

	cli.PrintLc(output, types.NewLcResult(ar, verified))

	return
}
