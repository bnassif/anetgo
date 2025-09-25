package address

import (
	"fmt"
	"strings"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var assignCmd = &cobra.Command{
	Use:   "assign SERVER_ID IP_ADDRESS [IP_ADDRESS...]",
	Short: "Assign one or more public IP addresses to a VM",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("assign-public-ip")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "assign-ip"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		// Build params if an arg was given
		params := map[string]string{
			"instanceid": args[0],
			"ip_address": strings.Join(args[1:], ","),
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
