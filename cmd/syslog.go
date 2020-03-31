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
	"net"
	"time"
)

// syslogCmd represents the syslog command
var syslogCmd = &cobra.Command{
	Use:   "syslog",
	Short: "Write fakelogs to a syslog server",
	Run:   syslogRun,
}

func init() {
	rootCmd.AddCommand(syslogCmd)
	syslogCmd.Flags().String("server", "127.0.0.1:514", "syslog server and port")
}

func syslogRun(cmd *cobra.Command, _ []string) {

	source, _ := rootCmd.Flags().GetString("source")
	count, _ := rootCmd.Flags().GetInt("count")
	server, _ := cmd.Flags().GetString("server")

	s, err := net.ResolveUDPAddr("udp", server)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("------------------")
	fmt.Println("Log count      :", count)
	fmt.Println("Log source     :", source)
	fmt.Println("Log destination:", c.RemoteAddr().String())
	fmt.Println("------------------")

	defer closeConnection(c)

	for i := 0; i < count; i++ {
		fakelogit := lib.NewLog(source)
		_, err := c.Write([]byte(fakelogit))
		if err != nil {
			fmt.Println("Error: ", err)
		}
		time.Sleep(time.Millisecond * 10)
	}
}

//Close closes connection on error
func closeConnection(c io.Closer) {
	err := c.Close()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
