/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Use this command to search your images",
	Long: `You can search your images by using this command. For example:

Phootecles search "40" "speed" "mountain"

After the command search write the all the terms you want to
search for. In the above example Phootecles will return all
images that contain either of the terms "40" or "speed" or
"mountain" in their description or title or inner text.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("You must provide at least one search term")
			return
		}

		// loop through all images and store images that contain the search term
		result := []Image{}
		images := getImages()
		for _, image := range images {
			// loop through all args
			for _, arg := range args {
				// make strings lowercase
				if strings.Contains(strings.ToLower(image.Title), strings.ToLower(arg)) ||
					strings.Contains(strings.ToLower(image.Description), strings.ToLower(arg)) ||
					strings.Contains(strings.ToLower(image.InsideText), strings.ToLower(arg)) ||
					strings.Contains(strings.ToLower(image.Id), strings.ToLower(arg)) {

					result = append(result, image)
				}
			}
		}

		// print the results
		for _, image := range result {
			formatImage(image)
		}
		fmt.Printf("Found %v matching images", len(result))
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
