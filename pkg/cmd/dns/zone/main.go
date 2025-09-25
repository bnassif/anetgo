package zone

import (
	"github.com/spf13/cobra"
)

var nestedKey = "DNSSet"
var Cmd = &cobra.Command{
	Use:   "zone",
	Short: "Manage & query DNS Zones",
}

func init() {
	Cmd.AddCommand(addCmd)
	Cmd.AddCommand(deleteCmd)
	Cmd.AddCommand(getCmd)
}
