/*
Copyright © 2024 4rkal <4rkal@horsefucker.org>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Loads an existing routine",
	Long: `Load an existing pushup routine.
	Example usage:
	go-pushups load glorious-mindworm`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			show_files()
		} else {
			routine, error := loadRoutine(args[0])
			if error != nil {
				fmt.Println(error)
				os.Exit(1)
			}
			run2(routine)
		}
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
