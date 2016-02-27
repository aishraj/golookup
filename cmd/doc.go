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
	"github.com/aishraj/golookup/doc"
	"github.com/spf13/cobra"
	"os"
)

// docCmd represents the doc command
var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Doc allows viewing the godoc of a Go package without having it locally set up.",
	Long: `Doc allows viewing the godoc of a Go package without having it locally set up.

It is an HTTP client to godoc.org and uses the response available from godoc.org to display the godoc.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("usage: ", RootCmd.Usage())
			return
		}
		d, err := doc.FetchDoc(args[0])
		if err != nil {
			fmt.Println("Unable to perform search. Error is: ", err)
			os.Exit(1)
		}
		fmt.Println(d)
	},
}

func init() {
	RootCmd.AddCommand(docCmd)
}
