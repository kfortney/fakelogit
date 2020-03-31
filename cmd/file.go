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
	"io"
	"os"
)

//fileCmd execute file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Write fakelogs to a file",
	Run:   fileRun,
}

func init() {
	rootCmd.AddCommand(fileCmd)
	fileCmd.Flags().String("out", "fakelogit.log", "output filename")
}

//fileRun execute file command
func fileRun(cmd *cobra.Command, _ []string) {

	source, _ := rootCmd.Flags().GetString("source")
	count, _ := rootCmd.Flags().GetInt("count")
	filename, _ := cmd.Flags().GetString("out")

	fmt.Println("------------------")
	fmt.Println("Log count      :", count)
	fmt.Println("Log source     :", source)
	fmt.Println("Log destination:", filename)
	fmt.Println("------------------")

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	for i := 0; i < count; i++ {
		fakelogit := lib.NewLog(source, lib.LogTimeFormat())
		_, _ = fmt.Fprintln(f, fakelogit)
		if err != nil {
			fmt.Println(err)
			closeFile(f)
		}
	}
	fmt.Println("log written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

//closeFile closes file on error
//
func closeFile(f io.Closer) {
	err := f.Close()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
