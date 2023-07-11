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

			switch commitType {
			case "feature", "feat":
				fmt.Println("feat")
			case "bug", "fix", "bugfix":
				fmt.Println("bugfix")
			case "chore":
				fmt.Println("chore")
			case "refactor":
				fmt.Println("refactor")
			case "test", "tests":
				fmt.Println("tests")
			case "documentation", "docs":
				fmt.Println("docs")
			case "style", "clean", "cleanup", "styleclean", "stylecleanup",
				"style/clean", "style/cleanup":
				fmt.Println("style clean up")
			case "build":
				fmt.Println("build")
			case "ci", "continuos", "integration", "continuosintegration":
				fmt.Println("ci")
			case "performance", "improvement", "performanceimprovement":
				fmt.Println("performance")
			default:
				fmt.Println("Could not find this type of confirmation! Try again.")
				os.Exit(0)
			}
		} else {
			fmt.Println("You need to provide the commit type. For example: gbc help build")
			os.Exit(0)
		}

	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}
