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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

//VERSION for project
var VERSION string
var cfgFile string
var source string
var count int

var rootCmd = &cobra.Command{
	Use:   "fakelogit",
	Short: "Generate fake logs",
	Long:  `Generate fake logs to test logging applications`,
}

//Execute version
func Execute(version string) {
	VERSION = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fakelogit.yaml)")
	rootCmd.PersistentFlags().StringVarP(&source, "source", "s", "rfc3164", `Choose log source below: ("rfc3164"|"rfc5424"|"fidelis"|"landcope"|"trendmicro"|"palottraffic"|"palothreat")`)
	rootCmd.PersistentFlags().IntVarP(&count, "count", "c", 10, "count of logs to generate")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.AddConfigPath("$HOME")
	viper.SetConfigName(".fakelogit")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Fakelog config file:", viper.ConfigFileUsed())
	}
}
