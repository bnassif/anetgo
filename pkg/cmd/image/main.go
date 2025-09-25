package image

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "image",
	Short: "Manage & query for VM Images in Altantic.net",
}

func init() {
	Cmd.AddCommand(getCmd)
}
