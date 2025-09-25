package power

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var resetOpts *viper.Viper

var resetCmd = &cobra.Command{
	Use:     "reset INSTANCE_ID [INSTANCE_ID...]",
	Aliases: []string{"cycle", "reboot"},
	Short:   "Power cycle a VM, or multiple VMs",
	Args:    cobra.MinimumNArgs(1),
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

		hardShutdown := offOpts.GetBool("immediate")
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

	resetCmd.PersistentFlags().BoolP("immediate", "i", false, "Whether to hard shutdown the VM(s)")

	_ = resetOpts.BindPFlag("immediate", resetCmd.PersistentFlags().Lookup("immediate"))
}
