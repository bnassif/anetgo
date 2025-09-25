package snapshot

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete SNAPSHOT_ID",
	Short: "Delete a snapshot",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("delete-snapshot")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "return"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		// Build params if an arg was given
		params := map[string]string{
			"SnapshotID": args[0],
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
