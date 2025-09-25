package address

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getOpts *viper.Viper

var getCmd = &cobra.Command{
	Use:   "get [IP_ADDRESS...]",
	Short: "List available public IPs, optionally filtering items",
	Run: func(cmd *cobra.Command, args []string) {
		action := string("list-public-ips")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "KeysSet"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		params := make(map[string]string)
		if len(args) > 0 {
			params["ip_address"] = args[0]
		} else {
			// If a specific IP is not passed, then consider filter flags
			if location := getOpts.GetString("location"); len(location) > 0 {
				params["location"] = location
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

func init() {
	getOpts = viper.New()

	getCmd.PersistentFlags().StringP("location", "l", "", "Filter public IPs by their location")

	_ = getOpts.BindPFlag("location", getCmd.PersistentFlags().Lookup("location"))
}
