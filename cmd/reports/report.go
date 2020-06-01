package reports

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ReportCmd represents the report command
var ReportCmd = &cobra.Command{
	Use:   "report",
	Short: "Shows the reports of SCM.",
	Long:  `Shows the reports of SCM.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		fmt.Print("report")
	},
}
