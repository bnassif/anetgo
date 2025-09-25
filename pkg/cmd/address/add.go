package address

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addOpts *viper.Viper

var addCmd = &cobra.Command{
	Use:     "add [LOCATION]",
	Aliases: []string{"reserve"},
	Short:   "Reserve a public IP, or many",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("reserve-public-ip")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "reserve-ip"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		params := map[string]string{
			"location": args[0],
		}

		if qty := addOpts.GetInt("qty"); qty > 0 {
			params["qty"] = fmt.Sprintf("%d", qty)
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
	addOpts = viper.New()

	addCmd.PersistentFlags().IntP("qty", "q", 0, "The number of IP addresses to reserve")

	_ = addOpts.BindPFlag("qty", addCmd.PersistentFlags().Lookup("qty"))
}
