package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("no version yet.")
		os.Exit(0)
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}
