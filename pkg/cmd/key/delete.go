package key

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete KEY_ID",
	Short: "Delete an SSH key by the ID in Atlantic.net's Cloud",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("delete-sshkey")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := action
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		// Build params if an arg was given
		params := map[string]string{
			"key_id": args[0],
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
