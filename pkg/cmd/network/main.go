package network

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "network",
	Short: "Manage & query for LANs in Altantic.net",
}

func init() {
	Cmd.AddCommand(getCmd)
}
