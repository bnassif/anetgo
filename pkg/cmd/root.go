package cmd

import (
	"fmt"
	"os"

	"github.com/bnassif/anetgo/pkg/api"
	"github.com/bnassif/anetgo/pkg/cmd/address"
	"github.com/bnassif/anetgo/pkg/cmd/dns"
	"github.com/bnassif/anetgo/pkg/cmd/image"
	"github.com/bnassif/anetgo/pkg/cmd/key"
	"github.com/bnassif/anetgo/pkg/cmd/location"
	"github.com/bnassif/anetgo/pkg/cmd/network"
	"github.com/bnassif/anetgo/pkg/cmd/plan"
	"github.com/bnassif/anetgo/pkg/cmd/server"
	"github.com/bnassif/anetgo/pkg/cmd/snapshot"
	"github.com/bnassif/anetgo/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Version is the CLI build version (override at build time with -ldflags)
	Version = "dev"

	// Global viper instance for config/flags/env
	rootOpts *viper.Viper
)

var RootCmd = &cobra.Command{
	Version: Version,
	Use:     "anetctl",
	Short:   "CLI for Atlantic.Net Cloud API",
	Long: `A CLI interface to Atlantic.Net's Cloud API, enabling querying and managing
any resources in Atlantic.Net's Cloud.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cfg := &api.Config{
			URL:     rootOpts.GetString("url"),
			Version: rootOpts.GetString("api-version"),
			Timeout: rootOpts.GetInt("timeout"),
			Key:     rootOpts.GetString("key"),
			Secret:  rootOpts.GetString("secret"),
		}

		client := api.NewClient(cfg)
		clientCtx := cmdutil.WithClient(cmd.Context(), client)
		cmd.SetContext(clientCtx)

		rawResp := rootOpts.GetBool("raw")
		flagCtx := cmdutil.WithBool(cmd.Context(), &rawResp)
		cmd.SetContext(flagCtx)

		return nil
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Use a single viper instance
	rootOpts = viper.New()

	// Flags (available to all subcommands)
	RootCmd.PersistentFlags().String("key", "", "The API key to use")
	RootCmd.PersistentFlags().String("secret", "", "The API secret to use")
	RootCmd.PersistentFlags().String("url", "https://cloudapi.atlantic.net", "The API base URL to use")
	RootCmd.PersistentFlags().String("api-version", "2010-12-30", "The API version to target")
	RootCmd.PersistentFlags().Int("timeout", 30, "HTTP client timeout in seconds")
	RootCmd.PersistentFlags().Bool("raw", false, "Whether to return the raw, unparsed response from the API")

	// Bind cobra flags to viper
	_ = rootOpts.BindPFlag("key", RootCmd.PersistentFlags().Lookup("key"))
	_ = rootOpts.BindPFlag("secret", RootCmd.PersistentFlags().Lookup("secret"))
	_ = rootOpts.BindPFlag("url", RootCmd.PersistentFlags().Lookup("url"))
	_ = rootOpts.BindPFlag("api-version", RootCmd.PersistentFlags().Lookup("api-version"))
	_ = rootOpts.BindPFlag("timeout", RootCmd.PersistentFlags().Lookup("timeout"))
	_ = rootOpts.BindPFlag("raw", RootCmd.PersistentFlags().Lookup("raw"))

	// Environment variable support (e.g. ANET_KEY, ANET_SECRET, etc.)
	rootOpts.SetEnvPrefix("ANET")
	rootOpts.AutomaticEnv()

	// Add subcommands
	RootCmd.AddCommand(address.Cmd)
	RootCmd.AddCommand(dns.Cmd)
	RootCmd.AddCommand(image.Cmd)
	RootCmd.AddCommand(key.Cmd)
	RootCmd.AddCommand(location.Cmd)
	RootCmd.AddCommand(network.Cmd)
	RootCmd.AddCommand(plan.Cmd)
	RootCmd.AddCommand(server.Cmd)
	RootCmd.AddCommand(snapshot.Cmd)
	RootCmd.AddCommand(genDocsCmd)
}
