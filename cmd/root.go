/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/nexidian/gocliselect"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gbc",
	Short: "Simple CLI in Go for writing better commits.",
	Long: `A CLI written in Golang that provides a simple way to write commits 
following the Conventional Commits  (https://www.conventionalcommits.org/).`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
		menu := gocliselect.NewMenu("Commit type?")

		menu.AddItem("feat", "a commit that adds new functionality to Uno.")
		menu.AddItem("fix", "a commit that fixes a bug in Uno.")
		menu.AddItem(
			"chore", 
			`a catch-all type for any other commits. For 
			instance, if you're implementing a single feature and it makes sense 
			to divide the work into multiple commits, you should mark one commit 
			as feat and the rest as chore.`,
		)
		menu.AddItem("refactor", "a commit that refactor a functionality.")
		menu.AddItem("test", "a commit that adds tests.")
		menu.AddItem(
			"docs", 
			"a commit that adds or improves Uno's documentation.",
		)
		menu.AddItem("style", "cyan")
		menu.AddItem("build", "cyan")
		menu.AddItem("ci", "cyan")
		menu.AddItem(
			"perf", 
			"a commit that improves performance, without functional changes.",
		)

		choice := menu.Display()

		fmt.Printf("Choice: %s\n", choice)
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


