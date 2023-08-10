package cmd

import (
	"fmt"
	"goProcessReporter/drivers"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var password string
var apiUrl string
var gwUrl string
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		if apiUrl == "" && gwUrl == "" {
			fmt.Println("need url")
			os.Exit(1)
		}
		if apiUrl == "" {
			apiUrl = fmt.Sprintf("%s/api/v2/fn/ps/update", gwUrl)
			fmt.Println("use gateway")
			fmt.Println(apiUrl)
		}
		fmt.Println("Start Processing...")
		for {
			title := drivers.GetApplicationName(drivers.GetActiveWindowProcessAndTitle())
			fmt.Println(title)
			mediaTitle, mediaArtist := drivers.GetNowPlaying()
			drivers.Report(title, password, apiUrl, mediaTitle, mediaArtist)
			time.Sleep(1 * time.Second)
		}

	},
}

func init() {
	StartCmd.Flags().StringVarP(&password, "password", "p", "", "Password you set on your mx-js admin")
	StartCmd.Flags().StringVarP(&apiUrl, "apiurl", "u", "", "URL of your api. like https://GATEWAY/api/v2/fn/ps/update")
	StartCmd.Flags().StringVarP(&gwUrl, "apigw", "g", "", "GATEWAY of your api. like https://GATEWAY")
	StartCmd.MarkFlagRequired("password")
	RootCmd.AddCommand(StartCmd)
}
