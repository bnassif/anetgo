package snapshot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getOpts *viper.Viper

func detectSnapshotKey(arg string) (key string, value string) {
	// Try integer ID
	if _, err := strconv.Atoi(arg); err == nil {
		return "SnapshotId", arg
	}

	// Try UUID
	if _, err := uuid.Parse(arg); err == nil {
		return "uuid", arg
	}

	// Fall back to checksum (lowercased, just in case)
	return "checksum", strings.ToLower(arg)
}

var getCmd = &cobra.Command{
	Use:   "get [SNAPSHOT_ID/UUID/CHECKSUM]",
	Short: "List available snapshots, optionally filtering items",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := string("list-snapshots")
		rootKey := fmt.Sprintf("%sresponse", action)
		nestedKey := "snapshotsSet"
		// Get a client object based on cmd Context
		client := cmdutil.GetClient(cmd.Context())
		rawFlag := cmdutil.GetRawFlagValue(cmd.Context())

		params := make(map[string]string)
		if len(args) > 0 {
			key, value := detectSnapshotKey(args[0])
			params[key] = value
		} else {
			// If a specific ID is not passed, then consider filter flags
			if vmid := getOpts.GetInt("server"); vmid != 0 {
				params["InstanceId"] = fmt.Sprintf("%d", vmid)
			}

			if image := getOpts.GetString("image"); image != "" {
				params["image_key"] = image
			}

			if platform := getOpts.GetString("platform"); platform != "" {
				params["platform"] = platform
			}

			if ostype := getOpts.GetString("ostype"); ostype != "" {
				params["ostype"] = ostype
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

	getCmd.PersistentFlags().IntP("server", "s", 0, "Filter snapshots by the target VM")
	getCmd.PersistentFlags().StringP("image", "i", "", "Filter snapshots by the VMs' image")
	getCmd.PersistentFlags().StringP("platform", "p", "", "Filter snapshots by the VMs' platform")
	getCmd.PersistentFlags().StringP("os-type", "o", "", "Filter snapshots by the VMs' OS type")

	_ = getOpts.BindPFlag("server", getCmd.PersistentFlags().Lookup("server"))
	_ = getOpts.BindPFlag("image", getCmd.PersistentFlags().Lookup("image"))
	_ = getOpts.BindPFlag("platform", getCmd.PersistentFlags().Lookup("platform"))
	_ = getOpts.BindPFlag("os-type", getCmd.PersistentFlags().Lookup("os-type"))
}
