package snapshot

import (
	"fmt"
	"strings"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add INSTANCE_ID DESCRIPTION",
	Aliases: []string{"create", "make"},
	Short:   "Create a snapshot of the target VM",
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("create-snapshot")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "return"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		// Build params if an arg was given
		params := map[string]string{
			"InstanceId":  args[0],
			"Description": strings.Join(args[1:], " "),
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
