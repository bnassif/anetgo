package image

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [IMAGE_ID]",
	Short: "List available images, or optionally a specific one",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("describe-image")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "imagesset"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		// Build params if an arg was given
		var params map[string]string
		if len(args) > 0 {
			params = map[string]string{
				"imageid": args[0],
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
