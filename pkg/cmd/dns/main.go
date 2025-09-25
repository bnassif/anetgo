package dns

import (
	"github.com/bnassif/anetgo/pkg/cmd/dns/record"
	"github.com/bnassif/anetgo/pkg/cmd/dns/zone"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "dns",
	Short: "Manage DNS Zones & Records",
}

func init() {
	Cmd.AddCommand(record.Cmd)
	Cmd.AddCommand(zone.Cmd)
}
