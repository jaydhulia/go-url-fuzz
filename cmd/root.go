package cmd

import (
	"os"

	"github.com/go-url-fuzz/config"
	"github.com/go-url-fuzz/requests"
	"github.com/spf13/cobra"
)

var (
	// root cobra command
	rootCmd = &cobra.Command{
		Use:   "go-url-fuzz",
		Short: "Fuzz URLs in Go!",
		Long:  "Just a fun tool I am developing as a way learn Go and to fuzz URLs in Go. Find those hidden directories on websites that developers forgot to remove! (inspired by wfuzz)",
		RunE:  runMain,
	}
)

func init() {
	rootCmd.Flags().StringVarP(&config.AppConfiguration.RequestType, "type", "X", config.DefaultRequestType, "request type (currently only GET supported)")
	rootCmd.Flags().StringVarP(&config.AppConfiguration.URL, "url", "u", "", "URL to fuzz, for example: \"https://example.com/FUZZ\" ")
	rootCmd.MarkFlagRequired("url")
	rootCmd.Flags().StringVarP(&config.AppConfiguration.Wordlist, "wordlist", "w", "", "path to a file containing wordlist to fuzz through")
	rootCmd.MarkFlagRequired("wordlist")
	rootCmd.Flags().IntSliceVarP(&config.AppConfiguration.ShowCodes, "show", "s", config.DefaultShowCodes, "example: show 200,204,300")
	rootCmd.Flags().IntVarP(&config.AppConfiguration.Concurrency, "concurrency", "c", config.DefaultConcurrentThreads, "path to a file containing wordlist to fuzz through")
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runMain(cmd *cobra.Command, args []string) error {
	if len(config.AppConfiguration.ShowCodes) > 0 {
		config.AppConfiguration.ShowCodesMap = make(map[int]int)
		for _, val := range config.AppConfiguration.ShowCodes {
			config.AppConfiguration.ShowCodesMap[val] = 1
		}
	}

	return requests.Execute()
}
