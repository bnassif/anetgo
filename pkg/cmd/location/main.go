package location

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "location",
	Short: "Manage & query for VM Plans in Altantic.net",
}

func init() {
	Cmd.AddCommand(getCmd)
}
