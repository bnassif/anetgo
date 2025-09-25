package power

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "power",
	Short: "Manage VM power states",
}

func init() {
	Cmd.AddCommand(offCmd)
	Cmd.AddCommand(onCmd)
	Cmd.AddCommand(resetCmd)
}
