package zone

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create"},
	Short:   "Create a new DNS zone",
	Run: func(cmd *cobra.Command, args []string) {
		action := string("dns-list-zones")
		rootKey := fmt.Sprintf("%sresponse", action)
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		params := map[string]string{
			"domain_name": args[0],
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
