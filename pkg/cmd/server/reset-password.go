package server

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var resetPasswordCmd = &cobra.Command{
	Use:   "reset-password INSTANCE_ID",
	Short: "Reset a VMs default user password with the new value returned",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("reset-password-instance")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "instancesSet"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		// Build params if an arg was given
		params := map[string]string{
			"instanceid": args[0],
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
