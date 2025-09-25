package power

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var onCmd = &cobra.Command{
	Use:   "on INSTANCE_ID [INSTANCE_ID...]",
	Short: "Power on a VM that is currently powered off, or multiple",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("power-on-instance")
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
