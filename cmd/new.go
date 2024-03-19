/*
Copyright © 2024 4rkal <4rkal@horsefucker.org>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new routine",
	Long: `Will create a new routine
	Example usage:
	go-pushups new`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := newRoutine("")
			if err != nil {
				fmt.Println("oh oh", err)
				os.Exit(1)
			}
		} else {
			err := newRoutine(args[0])
			if err != nil {
				fmt.Println("oh oh", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
