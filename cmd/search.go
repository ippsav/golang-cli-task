/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ippsav/golang-cli-task/handler"
	"github.com/spf13/cobra"
)

var Sort string
var Ignore string
var Handler *handler.Handler

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [VALUE] [FLAGS]",
	Short: "Search command accepts a search string to query Github API",
	Run: func(cmd *cobra.Command, args []string) {
    if len(args) == 0{
      fmt.Println("search command needs a search string to query Github API")
      os.Exit(1)
    }
    Handler.GetRepositories(args[0],"","")
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
  searchCmd.Flags().StringVarP(&Sort,"sort", "s", "asc", "sort parameter will sort results by repository name in Ascending (asc) or Descending (desc) order.")
  searchCmd.Flags().StringVarP(&Ignore,"ignore", "i", "", "ignore parameter will ignore repositories, where the name of the repository includes the provided string")

  Handler = handler.NewHandler()
}

