package cmd

import (
	"goProcessReporter/drivers/api"
	"goProcessReporter/drivers/config"
	"goProcessReporter/drivers/music"
	"goProcessReporter/drivers/utils"
	"goProcessReporter/drivers/winapi"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(filepath.Dir(os.Args[0]), os.Args[0])
		config := config.ReadConfig(configPath)
		cycle(config)
		os.Exit(0)
	},
}

func init() {
	StartCmd.Flags().StringVarP(&configPath, "config", "c", "./config.yml", "Config Path")
	StartCmd.MarkFlagRequired(configPath)
	RootCmd.AddCommand(StartCmd)
}

func cycle(configs config.Config) {
	for {
		title := utils.GetApplicationName(winapi.GetActiveWindowProcessAndTitle())
		title = utils.ReplaceString(title, configs.Replace, configs.ReplaceTo)
		title = utils.HideString(title, configs.Keywords)
		mediaTitle, mediaArtist := music.GetNowPlaying()
		api.Report(title, configs.ApiKey, configs.ApiURL, mediaTitle, mediaArtist)
		time.Sleep(time.Duration(configs.ReportTime) * time.Second)
	}
}
