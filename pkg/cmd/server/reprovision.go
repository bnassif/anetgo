package server

import (
	"fmt"
	"strconv"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func detectImageKey(arg string) (key string, value string) {
	// Try integer ID
	if _, err := strconv.Atoi(arg); err == nil {
		return "snapshotid", arg
	}

	// Fall back to checksum (lowercased, just in case)
	return "imageid", arg
}

var reprovisionCmd = &cobra.Command{
	Use:   "reprovision NAME PLAN IMAGE/SNAPHOT",
	Short: "Reprovision a VM with a bare image or using a snapshot",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("reprovision-instance")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "return"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		// Build params if an arg was given
		params := map[string]string{
			"instanceid": args[0],
			"planname":   args[1],
		}

		key, value := detectImageKey(args[2])
		params[key] = value

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
