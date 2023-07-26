/*
Copyright © 2023 Allan Capistrano <allan.capistrano3@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/allancapistrano/gbc/config"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update gbc to the latest version available.",
	Long:  `Update gbc to the latest version available.`,
	Run: func(cmd *cobra.Command, args []string) {
		// if !config.IsLatestVersion(GBC_VERSION) {
		if !config.IsLatestVersion("v1.0.0") {
			// TODO: Just for testing, remove it later
			option := config.Option{Type: "rebuild", Value: ":eyes:"}
			config.UpdateEmojiSettings("gbc.conf", option, true)

			fmt.Println("It is not the latest version")
		} else {
			fmt.Println("It is the latest version")
		}

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
