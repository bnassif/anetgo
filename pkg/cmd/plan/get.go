package plan

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getOpts *viper.Viper

var getCmd = &cobra.Command{
	Use:   "get [PLAN_NAME]",
	Short: "List available plans, limit by platform, or get a specific one",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("describe-plan")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "plans"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		// Build params if an arg was given
		params := make(map[string]string)
		if len(args) > 0 {
			params["plan_name"] = args[0]
		}

		if platform := getOpts.GetString("platform"); platform != "" {
			params["platform"] = platform
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
	getOpts = viper.New()

	getCmd.PersistentFlags().StringP("platform", "p", "linux", "The platform to limit by; defaults to linux")

	_ = getOpts.BindPFlag("platform", getCmd.PersistentFlags().Lookup("platform"))
}
