package record

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmd/dns/util"
	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addOpts *viper.Viper

var addCmd = &cobra.Command{
	Use:     "add TYPE NAME CONTENT",
	Aliases: []string{"create"},
	Short:   "Create a new DNS Record",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("dns-create-zone-record")
		rootKey := fmt.Sprintf("%sresponse", action)
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		params := make(map[string]string)
		key, value := util.ZoneIdOrName(args[0])
		params[key] = value

		params["type"] = args[1]
		params["name"] = args[2]
		params["content"] = args[3]

		if portArg := addOpts.GetInt("port"); portArg > 0 {
			params["port"] = fmt.Sprintf("%d", portArg)
		}

		if weightArg := addOpts.GetInt("weight"); weightArg > 0 {
			params["weight"] = fmt.Sprintf("%d", weightArg)
		}

		if ttlArg := addOpts.GetInt("ttl"); ttlArg > 0 {
			params["ttl"] = fmt.Sprintf("%d", ttlArg)
		}

		if prioArg := addOpts.GetInt("priority"); prioArg > 0 {
			params["priority"] = fmt.Sprintf("%d", prioArg)
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

func init() {
	addOpts = viper.New()

	addCmd.PersistentFlags().IntP("port", "p", 0, "The port for SRV records")
	addCmd.PersistentFlags().IntP("ttl", "t", 0, "The TTL of the record; defaults to 3600")
	addCmd.PersistentFlags().IntP("weight", "w", 0, "The weight for SRV records")
	addCmd.PersistentFlags().IntP("priority", "P", 0, "The priority of the record")

	_ = addOpts.BindPFlag("port", addCmd.PersistentFlags().Lookup("port"))
	_ = addOpts.BindPFlag("ttl", addCmd.PersistentFlags().Lookup("ttl"))
	_ = addOpts.BindPFlag("weight", addCmd.PersistentFlags().Lookup("weight"))
	_ = addOpts.BindPFlag("priority", addCmd.PersistentFlags().Lookup("priority"))
}
