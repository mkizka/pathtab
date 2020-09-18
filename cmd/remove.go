/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"strings"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove PATH",
	Short: "Remove value from ~/.pathtab",
	Run: func(cmd *cobra.Command, args []string) {
		conf := readConf()
		newConfSlice := []string{}
		newPath := args[0]
		for _, path := range strings.Split(conf, "\n") {
			if path != newPath {
				newConfSlice = append(newConfSlice, path)
			}
		}
		newConf := strings.Join(newConfSlice, "\n")
		writeConf(newConf)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
