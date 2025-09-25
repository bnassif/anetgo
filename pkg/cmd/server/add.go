package server

import (
	"fmt"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addOpts *viper.Viper

var addCmd = &cobra.Command{
	Use:     "add NAME PLAN LOCATION IMAGE/SNAPHOT",
	Aliases: []string{"create", "make"},
	Short:   "Create a new VM or many with the same specs",
	Args:    cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("run-instance")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "instancesSet"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		// Build params if an arg was given
		params := map[string]string{
			"servername":  args[0],
			"planname":    args[1],
			"vm_location": args[2],
		}

		// Determine which Image ID param to add
		if fromSnapshot := addOpts.GetBool("from-snapshot"); fromSnapshot {
			params["Snapshotid"] = args[3]
		} else {
			params["imageid"] = args[3]
		}

		// Start adding optional params
		if backups := addOpts.GetBool("backups"); backups {
			params["enablebackup"] = "Y"
		}

		if newPass := addOpts.GetBool("new-password"); newPass {
			params["newclonepassword"] = "Y"
		}

		if qty := addOpts.GetInt("qty"); qty > 1 {
			params["qty"] = fmt.Sprintf("%d", qty)
		}

		if term := addOpts.GetString("term"); term != "" {
			params["term"] = term
		}

		if sshKey := addOpts.GetString("key"); sshKey != "" {
			params["key_id"] = sshKey
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

	addCmd.PersistentFlags().BoolP("backups", "b", false, "Whether to enable backups for the new VM(s)")
	addCmd.PersistentFlags().BoolP("from-snapshot", "s", false, "Whether to create VMs from a snapshot rather than bare image")
	addCmd.PersistentFlags().BoolP("new-password", "n", false, "When made from a snapshot, determines whether a new password should be set")
	addCmd.PersistentFlags().IntP("qty", "q", 0, "The amount of VMs to create in this call")
	addCmd.PersistentFlags().StringP("term", "t", "", "The term to create the new VM(s) with")
	addCmd.PersistentFlags().StringP("key", "k", "", "The SSH key to add to the admin user for the new VM(s)")

	_ = addOpts.BindPFlag("backups", addCmd.PersistentFlags().Lookup("backups"))
	_ = addOpts.BindPFlag("from-snapshot", addCmd.PersistentFlags().Lookup("from-snapshot"))
	_ = addOpts.BindPFlag("new-password", addCmd.PersistentFlags().Lookup("new-password"))
	_ = addOpts.BindPFlag("qty", addCmd.PersistentFlags().Lookup("qty"))
	_ = addOpts.BindPFlag("term", addCmd.PersistentFlags().Lookup("term"))
	_ = addOpts.BindPFlag("key", addCmd.PersistentFlags().Lookup("key"))
}
