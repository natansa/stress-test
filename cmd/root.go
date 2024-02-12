package cmd

import (
	"os"

	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

var url string
var requests int
var concurrency int

type RunEFunc func(cmd *cobra.Command, args []string) error

var rootCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "Stress test for web services",
	Long:  `Stress test for web services using GoLang and Cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ = cmd.Flags().GetString("url")
		requests, _ = cmd.Flags().GetInt("requests")
		concurrency, _ = cmd.Flags().GetInt("concurrency")
		Start(url, requests, concurrency)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "Preencha a URL", "URL of the service to be tested")
	rootCmd.Flags().IntP("requests", "r", 0, "Total number of requests")
	rootCmd.Flags().IntP("concurrency", "c", 0, "Number of simultaneous calls")
}
