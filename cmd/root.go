/*
Copyright © 2023 Allan Capistrano <allan.capistrano3@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/nexidian/gocliselect"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gbc",
	Short: "Simple CLI in Go for writing better commits.",
	Long: `A CLI written in Golang that provides a simple way to write commits 
following the Conventional Commits  (https://www.conventionalcommits.org/).`,
	Run: func(cmd *cobra.Command, args []string) {
		commitTypeMenu := gocliselect.NewMenu("Commit type?")

		commitTypeMenu.AddItem("Feat", "feat")
		commitTypeMenu.AddItem("Fix", "fix")
		commitTypeMenu.AddItem("Chore", "chore")
		commitTypeMenu.AddItem("Refactor", "refactor")
		commitTypeMenu.AddItem("Test", "test")
		commitTypeMenu.AddItem("Docs", "docs")
		commitTypeMenu.AddItem("Style", "style")
		commitTypeMenu.AddItem("Build", "build")
		commitTypeMenu.AddItem("CI", "ci")
		commitTypeMenu.AddItem("Perf", "perf")

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
			panic(err)
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gbc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
