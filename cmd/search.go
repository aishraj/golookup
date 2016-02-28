package cmd

import (
	"fmt"
	"github.com/aishraj/golookup/search"
	"github.com/spf13/cobra"
	"os"
)

var maxResults int

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
		results, err := search.Search(args[0], maxResults)
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
	searchCmd.Flags().IntVarP(&maxResults, "maxResults", "m", 10, "Maximum number of search results returned")

}
