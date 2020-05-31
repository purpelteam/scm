package cmd

import (
	"fmt"

	"github.com/purpeltim/scm/standards/util"
	"github.com/purpeltim/scm/tools/sysinfo"
	"github.com/spf13/cobra"
)

// SCMVersion variable
var SCMVersion string = "0.0.1"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the version of SCM.",
	Long:  `Shows the version of SCM.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Host Info
		hostInfo, err := sysinfo.GetHostInfo()
		if err != nil {
			util.ExitWithError(err)
		}
		fmt.Println("SCM version ", SCMVersion, hostInfo.OS.Family, "/", hostInfo.Architecture, "(", hostInfo.OS.Name, hostInfo.OS.Version, ")")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
