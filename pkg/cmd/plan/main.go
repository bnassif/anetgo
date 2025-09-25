package plan

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "plan",
	Short: "Manage & query for VM Plans in Altantic.net",
}

func init() {
	Cmd.AddCommand(getCmd)
}
