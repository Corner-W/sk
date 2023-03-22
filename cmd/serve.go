/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Corner-W/sk/log"
	"github.com/Corner-W/sk/module/register"
	"github.com/Corner-W/sk/telnet"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a sk server",
	Long: `Start your sk server command, include one telnet serve (default :5001).

you can login your shell localhost:5001 with your telnet client`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("server start")

		log.Debug("server start...")
		go telnet.Run()

		register.ModulesRun()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
