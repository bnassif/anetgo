package address

import (
	"fmt"
	"strings"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete IP_ADDRESS [IP_ADDRESS...]",
	Aliases: []string{"release"},
	Short:   "Release one or more public IP addresses from your account",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("release-public-ip")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "release-ip"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		params := map[string]string{
			"ip_address": strings.Join(args[:], ","),
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
