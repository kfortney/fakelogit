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
	"github.com/kfortney/fakelogit/lib"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print fakelogs to screen",
	Run:   printRun,
}

func init() {
	rootCmd.AddCommand(printCmd)
}

func printRun(_ *cobra.Command, _ []string) {

	source, _ := rootCmd.Flags().GetString("source")
	count, _ := rootCmd.Flags().GetInt("count")

	fmt.Println("------------------")
	fmt.Println("Log count      :", count)
	fmt.Println("Log source     :", source)
	fmt.Println("Log destination:", "Screen")
	fmt.Println("------------------")

	for i := 0; i < count; i++ {
		fakelogit := lib.NewLog(source)
		fmt.Println(fakelogit + "\n")
	}
}
