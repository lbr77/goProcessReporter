package cmd

import (
	"goProcessReporter/drivers/logger"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "goProcessReporter",
	Short: "A Process reporter used by Shiro",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logger.Log.Error(err)
		os.Exit(1)
	}
}
