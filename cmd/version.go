package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	ver = "v1.0"

	date = "2023-3-18"
)

// serverCmd represents the server command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "sk version",
	Long:  `version of this sk server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SK Site Generator " + ver + " " + date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
