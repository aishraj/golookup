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
