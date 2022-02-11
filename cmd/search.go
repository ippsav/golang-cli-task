/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Sort string
var Ignore string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search command accepts a search string to query Github API",
	Run: func(cmd *cobra.Command, args []string) {
    fmt.Println(args)
    fmt.Println(Sort,Ignore)
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
}

