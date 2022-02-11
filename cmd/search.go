/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ippsav/golang-cli-task/handler"
	"github.com/spf13/cobra"
)

// flags
var Sort string
var Ignore string
var Page int

// Github API
const githubAPI = "https://api.github.com"

// handler
var Handler *handler.Handler

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [value]",
	Short: "Search command accepts a search string to query Github API",
	Run: func(cmd *cobra.Command, args []string) {
		// setup logger
		f, err := os.OpenFile("/logs/request.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("error opening file: %v", err)
		}
    defer f.Close()
		l := log.New(f, "", log.LUTC)
		// setup handler
		Handler = handler.NewHandler(githubAPI, l)
		// checking args first
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
		if Sort != "asc" && Sort != "desc" {
			fmt.Println(cmd.Flag("sort").Usage)
			os.Exit(1)
		}
		if Page < 0 {
			fmt.Println(cmd.Flag("page").Usage)
			os.Exit(1)
		}

		// querying repositories
		repos, err := Handler.GetRepositories(args[0], Sort, Ignore, Page)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		prettifiedData, err := json.MarshalIndent(repos, "", "  ")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(string(prettifiedData))
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
	searchCmd.Flags().StringVarP(&Sort, "sort", "s", "asc", "sort parameter will sort results by repository name in Ascending (asc) or Descending (desc) order.")
	searchCmd.Flags().StringVarP(&Ignore, "ignore", "i", "", "ignore parameter will ignore repositories, where the name of the repository includes the provided string")
	searchCmd.Flags().IntVarP(&Page, "page", "p", 1, "page parameter will set the page number of the results to fetch.(page>0)")
}
