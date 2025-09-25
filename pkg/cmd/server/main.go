package server

import (
	"github.com/bnassif/anetgo/pkg/cmd/server/power"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "Manage & query VMs in Altantic.net",
}

func init() {
	Cmd.AddCommand(addCmd)
	Cmd.AddCommand(deleteCmd)
	Cmd.AddCommand(getCmd)
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(reprovisionCmd)
	Cmd.AddCommand(resetPasswordCmd)
	Cmd.AddCommand(resizeCmd)
	Cmd.AddCommand(power.Cmd)
}
