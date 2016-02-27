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
	"github.com/aishraj/golookup/search"
	"github.com/spf13/cobra"
	"os"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches for the given query across mutiple search providers",
	Long: `Search performs search for the given query throughout mutiple search providers.

It returns back a list of package names followed by pacakge description.
The maximum default number of packages returned is 20.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("usage: ", RootCmd.Usage())
			return
		}
		results, err := search.Search(args[0])
		if err != nil {
			fmt.Println("Unable to perform search. Error is: ", err)
			os.Exit(1)
		}
		for _, result := range results {
			fmt.Println(result)
		}
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.Flags().IntP("maxResults", "m", 20, "Maximum number of search results returned")

}
