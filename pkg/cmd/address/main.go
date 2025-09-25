package address

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "address",
	Short: "Manage & query public IP addresses",
}

func init() {
	Cmd.AddCommand(getCmd)
	Cmd.AddCommand(addCmd)
	Cmd.AddCommand(deleteCmd)
	Cmd.AddCommand(assignCmd)
	Cmd.AddCommand(unassignCmd)
}
