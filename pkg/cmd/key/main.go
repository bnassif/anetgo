package key

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "key",
	Short: "Manage & query SSH Keys in Altantic.net",
}

func init() {
	Cmd.AddCommand(getCmd)
	Cmd.AddCommand(addCmd)
	Cmd.AddCommand(deleteCmd)
}
