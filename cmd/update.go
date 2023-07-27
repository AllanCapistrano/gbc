/*
Copyright © 2023 Allan Capistrano <allan.capistrano3@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

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
		if !config.IsLatestVersion("v1.0.0") { // TODO: Use GBC_VERSION constant
			fmt.Printf(
				"It is not the latest version. Updating to v%s\n",
				GBC_VERSION,
			)

			commandString := `bash -c "$(curl --fail --show-error --silent --location https://raw.githubusercontent.com/AllanCapistrano/gbc/feat/updater/scripts/update.sh)"`
			command := exec.Command("/bin/bash", "-c", commandString)

			err := command.Run()
			if err != nil {
				log.Fatal("\nUnable to update gbc to the latest version.")
			}
		} else {
			// TODO: Change the message.
			fmt.Println("It is the latest version")
		}

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
