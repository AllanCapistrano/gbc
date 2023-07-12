/*
Copyright © 2023 Allan Capistrano <allan.capistrano3@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show the description of types of commits.",
	Long: `Show the description and some examples of each type of commit that 
are available.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 && args[0] != "" {
			commitType := strings.ToLower(args[0])
			var message string
			var example string
			var commandHelp string

			switch commitType {
			case "feature", "feat":
				message = "A commit that adds new functionality"
				example = "feat: new login page"
				commandHelp = fmt.Sprintf(
					"%s.\n\nFor example: \n\t%s.",
					message,
					example,
				)
			case "bug", "fix", "bugfix":
				message = "A commit that fixes a bug"
				example = "fix: getUser function"
				commandHelp = fmt.Sprintf(
					"%s.\n\nFor example: \n\t%s.",
					message,
					example,
				)
			case "chore":
				message = "A catch-all type for any other commits"
				example = "chore: add validations to updateUser function"
				commandHelp = fmt.Sprintf(
					"%s.\n\nFor example: \n\t%s.",
					message,
					example,
				)
			case "refactor":
				message = "A commit that refactors some functionality"
				example = "refactor: addUser function"
				commandHelp = fmt.Sprintf(
					"%s.\n\nFor example: \n\t%s.",
					message,
					example,
				)
			case "test", "tests":
				message = "A commit that adds or improves tests"
				example = "test: add unit test to updateUser"
				commandHelp = fmt.Sprintf(
					"%s.\n\nFor example: \n\t%s.",
					message,
					example,
				)
			case "documentation", "docs":
				message = "A commit that adds or improves documentation"
				example = "docs: add docs to getUser function"
				commandHelp = fmt.Sprintf(
					"%s.\n\nFor example: \n\t%s.",
					message,
					example,
				)
			case "style", "clean", "cleanup", "styleclean", "stylecleanup",
				"style/clean", "style/cleanup":
				message = "A commit that styles source code or removes unnecessary code, including comments"
				example = "style: formatting code"
				commandHelp = fmt.Sprintf(
					"%s.\n\nFor example: \n\t%s.",
					message,
					example,
				)
			case "build":
				message = "A commit that builds a new version of the project"
				example = "build: gbc v1.0.0"
				commandHelp = fmt.Sprintf(
					"%s.\n\nFor example: \n\t%s.",
					message,
					example,
				)
			case "ci", "continuos", "integration", "continuosintegration":
				message = "A commit that adds or improves continuous integration to the project"
				example = "ci: using GitHub Actions"
				commandHelp = fmt.Sprintf(
					"%s.\n\nFor example: \n\t%s.",
					message,
					example,
				)
			case "performance", "improvement", "performanceimprovement", "perf":
				message = "A commit that improves performance, without functional changes"
				example = "perf: getUser function 2x faster"
				commandHelp = fmt.Sprintf(
					"%s.\n\nFor example: \n\t%s.",
					message,
					example,
				)
			default:
				fmt.Println("Could not find this type of confirmation! Try again.")
				os.Exit(0)
			}

			fmt.Println(commandHelp)
		} else {
			fmt.Println("You need to provide the commit type. For example: gbc help build")
			os.Exit(0)
		}

	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}
