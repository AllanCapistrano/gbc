/*
Copyright © 2023 Allan Capistrano <allan.capistrano3@gmail.com>
*/
package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Emoji struct {
	Feat     string `json:"feat"`
	Fix      string `json:"fix"`
	Chore    string `json:"chore"`
	Refactor string `json:"refactor"`
	Test     string `json:"test"`
	Docs     string `json:"docs"`
	Style    string `json:"style"`
	Build    string `json:"build"`
	Ci       string `json:"ci"`
	Perf     string `json:"perf"`
}

func GetEmojis(fileName string, debug bool) Emoji {
	foundSettingsFile := true
	file, err := os.Open(fileName)
	if err != nil {
		if debug {
			fmt.Printf(
				"Couldn't open the '%s' file! Check the path or file name.\n",
				fileName,
			)
		}
		foundSettingsFile = false
	}

	defer file.Close()

	if foundSettingsFile {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		flagEmoji := false

		emojiSettingInString := ""
		var emojiSetting Emoji

		for scanner.Scan() {
			// Ignore the comments.
			if strings.Contains(scanner.Text(), "#") {
				continue
			}

			// Found emojis setting.
			if strings.Contains(scanner.Text(), "emojis") {
				flagEmoji = true
				emojiSettingInString = "{\n"

				continue
			}

			// Getting each commit emoji.
			if flagEmoji {
				emojiSettingInString += scanner.Text() + "\n"
			}

			// Stop the loop when finds the last commit emoji.
			if strings.Contains(scanner.Text(), "}") && flagEmoji {
				break
			}
		}

		if emojiSettingInString != "" {
			err := json.Unmarshal([]byte(emojiSettingInString), &emojiSetting)
			if err != nil {
				fmt.Println("Unable to convert emoji settings to an expected type. Check the emoji settings in the config file.")
				os.Exit(0)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Couldn't read a line from the file.")
			os.Exit(0)
		}

		return emojiSetting
	}

	return Emoji{
		Feat:     ":rocket:",
		Fix:      ":bug:",
		Chore:    ":white_check_mark:",
		Refactor: ":hammer:",
		Test:     ":memo:",
		Docs:     ":books:",
		Style:    ":sparkles:",
		Build:    ":construction:",
		Ci:       ":factory:",
		Perf:     ":chart_with_upwards_trend:",
	}
}

func EnableEmojis(fileName string, debug bool) bool {
	foundSettingsFile := true
	file, err := os.Open(fileName)
	if err != nil {
		if debug {
			fmt.Printf(
				"Couldn't open the '%s' file! Check the path or file name.\n",
				fileName,
			)
		}
		foundSettingsFile = false
	}

	defer file.Close()

	if foundSettingsFile {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		flagError := false
		var conversionError string
		enableEmojis := false

		for scanner.Scan() {
			// Ignore the comments.
			if strings.Contains(scanner.Text(), "#") {
				continue
			}

			// Found emojis setting.
			if strings.Contains(scanner.Text(), "enableEmojis") {
				temp := strings.Split(scanner.Text(), " = ")

				enableEmojis, err = strconv.ParseBool(temp[1])
				if err != nil {
					flagError = true
					conversionError = temp[1]
					break
				}
			}
		}

		if flagError {
			fmt.Printf("Can't convert '%s' to bool type.", conversionError)
			os.Exit(0)
		}

		return enableEmojis
	}

	return false
}
