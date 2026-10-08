// Package cmd is responsible for all the cli commands
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jnawaz/kmp-gen/internal/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise kmp-gen in your project",
	RunE:  runInit,
}

func runInit(cmd *cobra.Command, args []string) error {
	reader := bufio.NewReader(os.Stdin)

	pkg, err := prompt(reader, "Base package name (e.g. com.yourapp): ")
	if err != nil {
		return err
	}

	sourceRoot, err := config.DetectSourceRoot()
	if err != nil {
		return err
	}

	if sourceRoot != "" {
		fmt.Printf("Detected source root: %s\n", sourceRoot)
		confirm, err := prompt(reader, "Is this correct? (y/n): ")
		if err != nil {
			return err
		}
		if strings.ToLower(confirm) != "y" {
			sourceRoot, err = prompt(reader, "Enter source root path: ")
			if err != nil {
				return err
			}
		}
	} else {
		sourceRoot, err = prompt(reader, "Source root path (e.g. composeApp/src/commonMain/kotlin): ")
		if err != nil {
			return err
		}
	}

	arch, err := prompt(reader, "Default architecture [ddd]: ")
	if err != nil {
		return err
	}
	if arch == "" {
		arch = "ddd"
	}

	cfg := config.Config{
		Package:      pkg,
		SourceRoot:   sourceRoot,
		Architecture: arch,
	}

	if err := config.Write(cfg); err != nil {
		return err
	}

	fmt.Println("✓ Created kmp-gen.yaml")
	return nil
}

func prompt(reader *bufio.Reader, question string) (string, error) {
	fmt.Print(question)
	answer, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(answer), nil
}

func init() {
	rootCmd.AddCommand(initCmd)
}
