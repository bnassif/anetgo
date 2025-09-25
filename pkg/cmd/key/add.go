package key

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add KEY_NAME KEY_VALUE",
	Short: "Add an SSH key for use in the Atlantic.net Cloud",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("add-sshkey")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "result"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		// Build params if an arg was given
		params := map[string]string{
			"key_name":   args[0],
			"public_key": args[1],
		}

		cmdutil.HandleRequest(
			client,
			action,
			params,
			*rawFlag,
			rootKey,
			nestedKey,
		)
	},
}
