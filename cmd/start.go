/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/Gentleelephant/account-srv/config"
	"github.com/Gentleelephant/account-srv/internal"
	"github.com/Gentleelephant/account-srv/server"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: " start the server",
	Long:  `start the server`,
	Run: func(cmd *cobra.Command, args []string) {
		err := server.Start()
		if err != nil {
			internal.Logger.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	startCmd.PersistentFlags().StringVarP(&config.FilePath, "config", "c", "./config.yaml", "config file path")
}
