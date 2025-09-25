package server

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available instances in your account",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		action := string("list-instances")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "instancesSet"
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
