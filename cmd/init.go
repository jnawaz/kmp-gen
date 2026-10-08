// Package cmd is responsible for all the cli commands
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise kmp-gen in your project",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("init running")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
