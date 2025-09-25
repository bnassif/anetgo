package record

import (
	"github.com/spf13/cobra"
)

var nestedKey = "DNSSet"
var Cmd = &cobra.Command{
	Use:   "record",
	Short: "Manage & query DNS Records",
}

func init() {
	Cmd.AddCommand(addCmd)
	Cmd.AddCommand(deleteCmd)
	Cmd.AddCommand(getCmd)
}
