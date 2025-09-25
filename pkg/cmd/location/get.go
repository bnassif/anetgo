package location

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "List available locations, limit by platform, or get a specific one",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		action := string("list-locations")
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
