package cmd

import (
	"github.com/purpeltim/scm/cmd/reports"
)

func init() {
	RootCmd.AddCommand(reports.ReportCmd)
}
