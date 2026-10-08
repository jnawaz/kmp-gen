package cmd

import (
	"fmt"

	"github.com/jnawaz/kmp-gen/internal/config"
	"github.com/jnawaz/kmp-gen/internal/scaffold"
	"github.com/spf13/cobra"
)

var featureCmd = &cobra.Command{
	Use:   "feature [name]",
	Short: "Scaffold a new feature",
	Args:  cobra.ExactArgs(1),
	RunE:  runFeature,
}

func runFeature(cmd *cobra.Command, args []string) error {
	featureName := args[0]

	cfg, err := config.Read()
	if err != nil {
		return err
	}

	if err := scaffold.Feature(featureName, cfg); err != nil {
		return err
	}

	fmt.Printf("✓ Created feature: %s\n", featureName)
	return nil
}

func init() {
	rootCmd.AddCommand(featureCmd)
}
