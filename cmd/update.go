/*
Copyright © 2023 Allan Capistrano <allan.capistrano3@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/allancapistrano/gbc/config"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update gbc to the latest version available.",
	Long:  `Update gbc to the latest version available.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !config.IsLatestVersion(GBC_VERSION) {
			fmt.Printf(
				"It is not the latest version. Updating to v%s\n",
				GBC_VERSION,
			)
			var commandString string

			if runtime.GOOS == "android" {
				commandString = `bash -c "$(curl --fail --show-error --silent --location https://raw.githubusercontent.com/AllanCapistrano/gbc/main/scripts/update-termux.sh)"`
			} else {
				commandString = `bash -c "$(curl --fail --show-error --silent --location https://raw.githubusercontent.com/AllanCapistrano/gbc/main/scripts/update.sh)"`
			}

			command := exec.Command("/bin/bash", "-c", commandString)

			err := command.Run()
			if err != nil {
				log.Fatal("\nUnable to update gbc to the latest version.")
			}
		} else {
			fmt.Println("gbc is already at the latest version.")
		}

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
