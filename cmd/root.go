/*
Copyright © 2023 Allan Capistrano <allan.capistrano3@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/kyokomi/emoji/v2"
	"github.com/nexidian/gocliselect"
	"github.com/spf13/cobra"

	"github.com/allancapistrano/gbc/config"
	"github.com/allancapistrano/user-input"
)

const GBC_VERSION = "1.1.0"
const SETTINGS_FILE_NAME = "gbc.conf"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gbc",
	Short: "Simple CLI in Go for writing better commits.",
	Long: `Git Better Commit (gbc) is a CLI written in Golang that provides a 
simple way to write commits following the Conventional Commits 
(https://www.conventionalcommits.org/).`,
	Run: func(cmd *cobra.Command, args []string) {

		fstatus, _ := cmd.Flags().GetBool("version")
		if fstatus {
			fmt.Printf("gbc version %s\n", GBC_VERSION)
		} else {
			commitTypeMenu := gocliselect.NewMenu("Commit type?")
			gbcEmojis := config.GetEmojis(SETTINGS_FILE_NAME, false)
			enableEmojis := config.EnableEmojis(SETTINGS_FILE_NAME, false)

			commitTypeMenu.AddItem(
				emoji.Sprintf("%sFeature", gbcEmojis.Feat),
				"feat",
			)
			commitTypeMenu.AddItem(
				emoji.Sprintf("%sBug Fix", gbcEmojis.Fix),
				"fix",
			)
			commitTypeMenu.AddItem(
				emoji.Sprintf("%sChore", gbcEmojis.Chore),
				"chore",
			)
			commitTypeMenu.AddItem(
				emoji.Sprintf("%sRefactor", gbcEmojis.Refactor),
				"refactor",
			)
			commitTypeMenu.AddItem(
				emoji.Sprintf("%sTests", gbcEmojis.Test),
				"test",
			)
			commitTypeMenu.AddItem(
				emoji.Sprintf("%sDocumentation", gbcEmojis.Docs),
				"docs",
			)
			commitTypeMenu.AddItem(
				emoji.Sprintf("%sStyle/Clean Up", gbcEmojis.Style),
				"style",
			)
			commitTypeMenu.AddItem(
				emoji.Sprintf("%sBuild", gbcEmojis.Build),
				"build",
			)
			commitTypeMenu.AddItem(
				emoji.Sprintf("%sCI", gbcEmojis.Ci),
				"ci",
			)
			commitTypeMenu.AddItem(
				emoji.Sprintf("%sPerformance Improvement", gbcEmojis.Perf),
				"perf",
			)
			commitTypeMenu.AddItem(
				emoji.Sprintf("%sFirst Commit", gbcEmojis.FirstCommit),
				"first_commit",
			)

			commitType := commitTypeMenu.Display()

			if commitType == "" {
				os.Exit(0)
			}

			var commitMessage string
			var commandString string

			if enableEmojis {
				var commitTypeEmoji string

				switch commitType {
				case "feat":
					commitTypeEmoji = gbcEmojis.Feat
				case "fix":
					commitTypeEmoji = gbcEmojis.Fix
				case "chore":
					commitTypeEmoji = gbcEmojis.Chore
				case "refactor":
					commitTypeEmoji = gbcEmojis.Refactor
				case "test":
					commitTypeEmoji = gbcEmojis.Test
				case "docs":
					commitTypeEmoji = gbcEmojis.Docs
				case "style":
					commitTypeEmoji = gbcEmojis.Style
				case "build":
					commitTypeEmoji = gbcEmojis.Build
				case "ci":
					commitTypeEmoji = gbcEmojis.Ci
				case "perf":
					commitTypeEmoji = gbcEmojis.Perf
				case "first_commit":
					commitMessage = emoji.Sprintf("%s", gbcEmojis.FirstCommit)
				}

				commitType = emoji.Sprintf(
					"%s%s",
					commitTypeEmoji,
					commitType,
				)
			}

			if commitType == "first_commit" {
				commitMessage += "First commit"
				commandString = fmt.Sprintf(
					`git commit -m "%s"`, commitMessage,
				)
			} else {
				commitMessage = userinput.GetUserInput("What commit message do you want? ", true)
				commandString = fmt.Sprintf(
					`git commit -m "%s: %s"`, commitType, commitMessage,
				)
			}

			command := exec.Command("/bin/bash", "-c", commandString)

			err := command.Run()
			if err != nil {
				log.Fatal("\nCould not create the commit! Make sure you have file(s) to commit.")
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("Unable to execute root command")
	}
}

func init() {
	// Version flag
	rootCmd.Flags().BoolP("version", "V", false, "Shows the gbc version.")
	// Removing 'help' subcommand
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	// Removing 'completion' subcommand
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
