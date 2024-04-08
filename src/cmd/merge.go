/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/Merger"
	"github.com/spf13/cobra"
)

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge \"<FilePath1> <FilePath2> [FilePath3, ...]\"",
	Short: "This command merges one or more CSV-files into one.",
	Long: `This command reads data from each provided CSV-file and writes the data into a singular file in the order given as arguments.
	
<> are mandatory arguments, whereas [] are optional arguments.
Example of a <FilePath>: 'C://Program/MyCSVFile.csv'.
Example of command usage: 'merge C://Program/MyCSVFile.csv C://Program/MyCSVFile1.csv'`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			panic("Not enough arguments have been provided.")
		}
		useCache, _ := cmd.Flags().GetBool("cache")
		if useCache {
			fmt.Println("cache is set.")
		} else {
			Merger.MergeCSV(args)
			fmt.Println("merge called.")
		}
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
	mergeCmd.Flags().BoolP("cache", "c", false, "Help message for cache")
}
