package power

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var resetOpts *viper.Viper

var resetCmd = &cobra.Command{
	Use:     "reset",
	Aliases: []string{"cycle"},
	Short:   "Power off a VM that is currently powered on",
	Run: func(cmd *cobra.Command, args []string) {
		action := string("shutdown-instance")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "return"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		// Build params if an arg was given
		params := map[string]string{
			"instanceid": args[0],
		}

		hardShutdown := offOpts.GetBool("hard")
		if hardShutdown {
			params["reboottype"] = "hard"
		} else {
			params["reboottype"] = "soft"
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
	resetOpts = viper.New()

	resetCmd.PersistentFlags().BoolP("hard", "H", false, "Whether to hard shutdown the VM(s)")

	_ = resetOpts.BindPFlag("hard", resetCmd.PersistentFlags().Lookup("hard"))
}
