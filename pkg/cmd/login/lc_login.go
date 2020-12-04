package login

import (
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/cmd/internal/cli"
	"github.com/vchain-us/vcn/pkg/store"
)

// Execute the login action
func ExecuteLC(host, port string) error {
	apiKey, err := cli.ProvideLcApiKey()
	if err != nil {
		return err
	}
	if apiKey != "" {
		_ = api.NewLcUser(apiKey, host, port)
		// Store the new config
		if err := store.SaveConfig(); err != nil {
			return err
		}
	}
	return nil
}
