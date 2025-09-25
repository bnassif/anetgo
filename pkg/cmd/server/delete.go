package server

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete INSTANCE_ID [INSTANCE_ID...]",
	Aliases: []string{"terminate", "destroy"},
	Short:   "Remove one or more VMs from your account",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("terminate-instance")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "instancesSet"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		// Build params if an arg was given
		params := make(map[string]string)

		if len(args) == 1 {
			// Single instance use just `instanceid`
			params["instanceid"] = args[0]
		} else if len(args) > 1 {
			// Multiple instances use `instanceid_1`, `instanceid_2`, ...
			for i, arg := range args {
				key := fmt.Sprintf("instanceid_%d", i+1)
				params[key] = arg
			}
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
