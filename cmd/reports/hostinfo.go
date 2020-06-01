package reports

import (
	"encoding/json"
	"fmt"

	"github.com/purpeltim/scm/standards/util"
	"github.com/purpeltim/scm/tools/sysinfo"
	"github.com/spf13/cobra"
)

// hostInfoCmd represents the report of hostinformation command
var hostInfoCmd = &cobra.Command{
	Use:   "hostinfo",
	Short: "Shows the host information.",
	Long:  `Shows the host information`,
	Run: func(cmd *cobra.Command, args []string) {
		// Host Info
		hostInfo, err := sysinfo.GetHostInfo()
		if err != nil {
			util.ExitWithError(err)
		}
		b, err := json.Marshal(hostInfo)
		if err != nil {
			util.ExitWithError(err)
		}
		fmt.Println(string(b))
	},
}

func init(){
	ReportCmd.AddCommand(hostInfoCmd)
}