package key

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "List available SSH Keys",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		action := string("list-sshkeys")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "KeysSet"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		cmdutil.HandleRequest(
			client,
			action,
			nil,
			*rawFlag,
			rootKey,
			nestedKey,
		)
	},
}
