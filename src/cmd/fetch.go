/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/Fetcher"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch <Url> <Branch> <Whitelist_Filepath> [Url1] [Branch1] [Whitelist_Filepath1]",
	Short: "fetch one or more files from a Git repository",
	Long: `The fetcher is used to pull whitelisted files/folders from one or more git repositories. It takes a string containing 3 values:

	1. A URL to a git based repository
	2. A branch name
	3. A path to a whitelist file
	`,

	//Args: cobra.MinimumNArgs(3),

	Run: func(cmd *cobra.Command, args []string) {
		// Set flags
		filePath, err := cmd.Flags().GetString("repo-file")
		if err != nil {
			panic(err)
		}
		branch, err := cmd.Flags().GetString("branch")
		if err != nil {
			panic(err)
		}
		whitelist, err := cmd.Flags().GetString("whitelist")
		if err != nil {
			panic(err)
		}

		// Check flags and adjust required arguments
		requiredArgs := 3
		if filePath != "" {

			if branch != "" {
				requiredArgs--
			}
			if whitelist != "" {
				requiredArgs--
			}
			file, err := os.Open(filePath)
			if err != nil {
				panic(err)
			}

			defer file.Close()

			// Read the file
			// Check if the file is empty or incorrect amount of arguments
			scanner := bufio.NewScanner(file)
			scanner.Scan()
			temp_args := strings.Split(strings.Trim(scanner.Text(), " \n"), " ")

			if len(temp_args) == 0 || len(temp_args) != requiredArgs {
				panic(fmt.Sprintf("file is empty or arguments is not the required amount -> %d", requiredArgs))
			}

			// Add the rest of the arguments
			for scanner.Scan() {
				text := strings.Split(scanner.Text(), " ")
				temp_args = append(temp_args, text...)
			}

			args = temp_args
		}

		if len(args)%3 != 0 && branch == "" && whitelist == "" && filePath == "" {
			panic("arguments is not divisable by 3")
		}

		if branch != "" && whitelist != "" {
			// construct the args array with branch and whitelist
			temp_args := []string{}
			for i := 0; i < len(args); i++ {
				temp_args = append(temp_args, args[i])
				temp_args = append(temp_args, branch)
				temp_args = append(temp_args, whitelist)
			}
			args = temp_args
		} else {

			// construct the args array based on individual flags
			temp_args := []string{}
			if branch != "" {
				for i := 0; i < len(args); i++ {
					// inserts branch at every 2nd index and shifts the rest of the values
					if (i+1)%2 == 0 {
						temp_args = append(temp_args, branch)
						temp_args = append(temp_args, args[i])

					} else {
						temp_args = append(temp_args, args[i])
					}
				}
				args = temp_args
			}
			if whitelist != "" {
				for i := 0; i < len(args); i++ {
					if (i+1)%2 == 0 {
						temp_args = append(temp_args, args[i])
						temp_args = append(temp_args, whitelist)
					} else {
						temp_args = append(temp_args, args[i])
					}
				}
				args = temp_args
			}
		}
		err = Fetcher.ListingReposForFetch(args)
		if err != nil {
			log.Print(err)
		}
		fmt.Println("\nFetch complete.")
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	fetchCmd.Flags().String("branch", "", "Branch name for all repositories")
	fetchCmd.Flags().String("whitelist", "", "Path to a whitelist file for all repositories")
	fetchCmd.Flags().String("repo-file", "", "Path to a file containing the list of repositories to fetch")
}
