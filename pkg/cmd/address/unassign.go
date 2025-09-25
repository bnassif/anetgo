package address

import (
	"fmt"
	"strings"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var unassignCmd = &cobra.Command{
	Use:   "unassign IP_ADDRESS",
	Short: "Unassign an IP address from a VM",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("unassign-public-ip")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "unassign-ip"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		// Build params if an arg was given
		params := map[string]string{
			"ip_address": strings.Join(args[:], ","),
		}

		// Send the request
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
