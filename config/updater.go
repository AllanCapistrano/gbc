/*
Copyright © 2023 Allan Capistrano <allan.capistrano3@gmail.com>
*/
package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tcnksm/go-latest"
)

type Option struct {
	Comment string
	Type    string
	Value   string
}

// Checks if gbc is at the latest version by passing the current version.
func IsLatestVersion(version string) bool {
	githubTag := &latest.GithubTag{
		Owner:      "AllanCapistrano",
		Repository: "gbc",
	}

	res, _ := latest.Check(githubTag, version)

	return !res.Outdated
}

// Adds a newOption in the emoji struct present in fileName. You can enable or
// disable debug messages.
func UpdateEmojiSettings(fileName string, newOption Option, debug bool) {
	foundSettingsFile := true
	homeDir, err := os.UserHomeDir()
	if err != nil {
		if debug {
			fmt.Println("Couldn't open the user home directory.")
		}
		foundSettingsFile = false
	}

	filePath := filepath.Join(homeDir, ".config", "gbc", fileName)
	file, err := os.ReadFile(filePath)
	if err != nil {
		if debug {
			fmt.Printf(
				"Couldn't open the '%s' file! Check the path or file name.\n",
				fileName,
			)
		}
		foundSettingsFile = false
	}

	if foundSettingsFile {
		flagEmoji := false
		lines := strings.Split(string(file), "\n")
		var index uint8

		for i, line := range lines {
			if strings.Contains(line, "#") {
				continue
			}

			if strings.Contains(line, "emojis") {
				flagEmoji = true
				continue
			}

			if flagEmoji {
				if strings.Contains(line, "}") && flagEmoji {
					// Getting the position before the last option (First Commit).
					index = uint8(i) - 1
					break
				}
			}
		}

		// Adding the new option.
		lines = append(lines[:index+1], lines[index:]...)
		lines[index] = fmt.Sprintf("    \"%s\": \"%s\",", newOption.Type, newOption.Value)

		if flagEmoji {
			output := strings.Join(lines, "\n")
			err = os.WriteFile(filePath, []byte(output), 0644)
			if err != nil {
				log.Fatalf(
					"Could not write to file '%s'! Check the path or file name.\n",
					fileName,
				)
			}
		}

	}
}

// Adds a new option with a comment to fileName. You can enable or disable debug 
// messages.
func AddNewSetting(fileName string, option Option, debug bool) {
	foundSettingsFile := true
	homeDir, err := os.UserHomeDir()
	if err != nil {
		if debug {
			fmt.Println("Couldn't open the user home directory.")
		}
		foundSettingsFile = false
	}

	filePath := filepath.Join(homeDir, ".config", "gbc", fileName)
	file, err := os.ReadFile(filePath)
	if err != nil {
		if debug {
			fmt.Printf(
				"Couldn't open the '%s' file! Check the path or file name.\n",
				fileName,
			)
		}
		foundSettingsFile = false
	}

	if foundSettingsFile {
		lines := strings.Split(string(file), "\n")
		comment := fmt.Sprintf("# %s\n", option.Comment)
		setting := fmt.Sprintf("%s = %s", option.Type, option.Value)

		// Adding the new option at the end of the file.
		lines = append(lines, comment, setting)

		output := strings.Join(lines, "\n")
		err = os.WriteFile(filePath, []byte(output), 0644)
		if err != nil {
			log.Fatalf(
				"Could not write to file '%s'! Check the path or file name.\n",
				fileName,
			)
		}
	}
}