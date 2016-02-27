// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command (golookup) when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "golookup",
	Short: "golookup is a quick and dirty way to discover Go pacakges.",
	Long: `golookup is a command line tool that allows discovery and
documenation access for Go packages without having them locally available.

It searchs across package name or description across multiple providers and
allows viewing its documentation without having to fetch the package locally.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
