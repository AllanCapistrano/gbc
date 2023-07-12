/*
Copyright © 2023 Allan Capistrano <allan.capistrano3@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/kyokomi/emoji/v2"
	"github.com/nexidian/gocliselect"
	"github.com/spf13/cobra"

	"github.com/allancapistrano/gbc/config"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gbc",
	Short: "Simple CLI in Go for writing better commits.",
	Long: `Git Better Commit (gbc) is a CLI written in Golang that provides a 
simple way to write commits following the Conventional Commits 
(https://www.conventionalcommits.org/).`,
	Run: func(cmd *cobra.Command, args []string) {
		commitTypeMenu := gocliselect.NewMenu("Commit type?")
		gbcEmoji := config.GetEmojis("config/gbc.conf")

		commitTypeMenu.AddItem(
			emoji.Sprintf("%sFeature", gbcEmoji.Feat),
			"feat",
		)
		commitTypeMenu.AddItem(
			emoji.Sprintf("%sBug Fix", gbcEmoji.Fix),
			"fix",
		)
		commitTypeMenu.AddItem(
			emoji.Sprintf("%sChore", gbcEmoji.Chore),
			"chore",
		)
		commitTypeMenu.AddItem(
			emoji.Sprintf("%sRefactor", gbcEmoji.Refactor),
			"refactor",
		)
		commitTypeMenu.AddItem(
			emoji.Sprintf("%sTests", gbcEmoji.Test),
			"test",
		)
		commitTypeMenu.AddItem(
			emoji.Sprintf("%sDocumentation", gbcEmoji.Docs),
			"docs",
		)
		commitTypeMenu.AddItem(
			emoji.Sprintf("%sStyle/Clean Up", gbcEmoji.Style),
			"style",
		)
		commitTypeMenu.AddItem(
			emoji.Sprintf("%sBuild", gbcEmoji.Build),
			"build",
		)
		commitTypeMenu.AddItem(
			emoji.Sprintf("%sCI", gbcEmoji.Ci),
			"ci",
		)
		commitTypeMenu.AddItem(
			emoji.Sprintf("%sPerformance Improvement", gbcEmoji.Perf),
			"perf",
		)

		commitType := commitTypeMenu.Display()
		buffer := bufio.NewReader(os.Stdin)

		fmt.Print("What commit message do you want? ")

		commitMessage, err := buffer.ReadString('\n')
		if err != nil {
			fmt.Println("Unable to read commit message! Try again.")
			os.Exit(0)
		}

		commandString := fmt.Sprintf(
			`git commit -m "%s: %s"`, commitType, commitMessage,
		)

		command := exec.Command("/bin/bash", "-c", commandString)

		err = command.Run()
		if err != nil {
			fmt.Println(
				"\nCould not create the commit! Make sure you have files to commit.",
			)
			os.Exit(0)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Removing 'help' subcommand
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	// Removing 'completion' subcommand
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
