/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/allancapistrano/gbc/config"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update gbc to the latest version available.",
	Long: `Update gbc to the latest version available.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !config.IsLatestVersion(GBC_VERSION) {
			//TODO
			fmt.Println("It is not the latest version")
		} else {
			fmt.Println("It is the latest version")
		}

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
