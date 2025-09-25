package power

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var offOpts *viper.Viper

var offCmd = &cobra.Command{
	Use:   "off INSTANCE_ID [INSTANCE_ID...]",
	Short: "Power off a VM that is currently powered on, or multiple VMs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("shutdown-instance")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "instancesSet"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		// Build params if an arg was given
		params := make(map[string]string)

		if len(args) == 1 {
			// Single instance use just `instanceid`
			params["instanceid"] = args[0]
		} else if len(args) > 1 {
			// Multiple instances use `instanceid_1`, `instanceid_2`, ...
			for i, arg := range args {
				key := fmt.Sprintf("instanceid_%d", i+1)
				params[key] = arg
			}
		}

		hardShutdown := offOpts.GetBool("immediate")
		if hardShutdown {
			params["shutdowntype"] = "hard"
		} else {
			params["shutdowntype"] = "soft"
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

func init() {
	offOpts = viper.New()

	offCmd.PersistentFlags().BoolP("immediate", "i", false, "Whether to hard shutdown the VM(s)")

	_ = offOpts.BindPFlag("immediate", offCmd.PersistentFlags().Lookup("immediate"))
}
