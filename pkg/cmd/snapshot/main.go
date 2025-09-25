package snapshot

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Manage & query VM Snapshots in Altantic.net",
}

func init() {
	Cmd.AddCommand(getCmd)
	Cmd.AddCommand(addCmd)
	Cmd.AddCommand(deleteCmd)
	Cmd.AddCommand(restoreCmd)
}
