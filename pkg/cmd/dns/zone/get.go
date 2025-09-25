package zone

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmd/dns/util"
	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get [ZONE_ID_OR_NAME]",
	Aliases: []string{"list"},
	Short:   "Get all DNS zones, or optionally a single one",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("dns-list-zones")
		rootKey := fmt.Sprintf("%sresponse", action)
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		params := make(map[string]string)
		if len(args) > 0 {
			key, value := util.ZoneIdOrName(args[0])
			params[key] = value
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
