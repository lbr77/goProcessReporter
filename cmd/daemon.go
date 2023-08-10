package cmd

import (
	"fmt"
	"goProcessReporter/drivers/logger"
	"goProcessReporter/drivers/winapi"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var configPath string
var DaemonStartCmd = &cobra.Command{
	Use:   "start-daemon",
	Short: "Start in background",
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command(os.Args[0], "start", "--config", configPath)
		fmt.Println(command)
		command.Start()
		os.Exit(0)
	},
}
var DaemonStopCmd = &cobra.Command{
	Use:   "stop-daemon",
	Short: "Stop in background",
	Run: func(cmd *cobra.Command, args []string) {
		programName := filepath.Base(os.Args[0])
		pids := winapi.GetRunningPids(programName)
		for _, pid := range pids {
			if pid != uint32(os.Getpid()) {
				winapi.StopPid(pid)
				logger.Log.Info("Daemon Closed.")
			}
		}
	},
}

func init() {
	DaemonStartCmd.Flags().StringVarP(&configPath, "config", "c", "config.yml", "Config Path")

	RootCmd.AddCommand(DaemonStartCmd)
	RootCmd.AddCommand(DaemonStopCmd)
}
