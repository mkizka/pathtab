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
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var shouldAll bool
var shouldSort bool

func getPATH(isSystem bool) []string {
	if isSystem {
		splitter := ":"
		if runtime.GOOS == "windows" {
			splitter = ";"
		}
		return strings.Split(os.Getenv("PATH"), splitter)
	}
	conf := readConf()
	trimedConf := strings.TrimSpace(conf)
	return strings.Split(trimedConf, "\n")
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show list of PATH environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		path := getPATH(shouldAll)
		if shouldSort {
			sort.Strings(path)
		}
		fmt.Println(strings.Join(path, "\n"))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&shouldAll, "all", "a", false, "show all PATH")
	listCmd.Flags().BoolVarP(&shouldSort, "sort", "s", false, "sort result")
}
