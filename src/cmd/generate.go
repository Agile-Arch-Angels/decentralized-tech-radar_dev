/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	html "github.com/NovoNordisk-OpenSource/decentralized-tech-radar/HTML"
	Reader "github.com/NovoNordisk-OpenSource/decentralized-tech-radar/SpecReader"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate <FilePath>",
	Short: "Generates an HTML file that is populated with CSV data",
	Long: `This command reads from a given CSV file, and then generates an HTML file populated with the CSV data.
The FilePath refers to the designated path. An example would be: 'C://Program/MyCSVFile.csv'`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		specs := Reader.CsvToString(args[0])
		html.GenerateHtml(specs)
		open, _ := cmd.Flags().GetBool("open")
		if open {
			if runtime.GOOS == "windows" {
				fmt.Println("GOOS: " + runtime.GOOS)
				cmd := exec.Command("explorer", "index.html")
				_, err := cmd.CombinedOutput()
				if err != nil {
					panic(err)
				} 
			} else if runtime.GOOS == "linux" {
				fmt.Print("GOOS: " + runtime.GOOS)
				cmd := exec.Command("open", "index.html")
				err := cmd.Run()
				if err != nil {
					panic(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().Bool("open", false, "Add this flag to open the generated html file when it has been generated")
}
