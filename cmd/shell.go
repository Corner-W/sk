package cmd

import (
	"github.com/Corner-W/sk/log"
	"github.com/Corner-W/sk/pkg/shellcommand"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start a shell server",
	Long: `Start your shell server command, login server (default :2222).

you can login your shell localhost:2222 with your ssh client`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("server start")

		log.Debug("ssh shell server start...")
		shellcommand.Run()

	},
}

func init() {
	rootCmd.AddCommand(shellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
