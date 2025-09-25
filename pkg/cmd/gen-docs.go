// cmd/init.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
)

func genMarkdownDocs(path string) {
	cmd := RootCmd
	err := doc.GenMarkdownTree(cmd, path)
	if err != nil {
		fmt.Println(err)
		os.Exit(127)
	}
}

func genReSTDocs(path string) {
	cmd := RootCmd
	err := doc.GenReSTTree(cmd, path)
	if err != nil {
		fmt.Println(err)
		os.Exit(127)
	}
}

func genYAMLDocs(path string) {
	cmd := RootCmd
	err := doc.GenYamlTree(cmd, path)
	if err != nil {
		fmt.Println(err)
		os.Exit(127)
	}
}

func genManDocs(path string) {
	cmd := RootCmd
	err := doc.GenManTree(cmd, &doc.GenManHeader{
		Title:   "anetctl",
		Section: "1",
	}, path)
	if err != nil {
		fmt.Println(err)
		os.Exit(127)
	}
}

var genDocsCmd = &cobra.Command{
	Use:   "gen-docs PATH",
	Short: "Generate documentation for this tool",
	Long: `Generates documentation in the specified format for this tool.
Supported formats:
- markdown
- ReST
- YAML
- Man pages
	`,
	Hidden: true,
	Args:   cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		quietMode, _ := cmd.Flags().GetBool("quiet")
		format, _ := cmd.Flags().GetString("format")

		// Check if the path exists
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if !quietMode {
				fmt.Printf("Path %s does not exist\n", path)
			}
			os.Exit(127)
		}

		switch format {
		case "markdown":
			genMarkdownDocs(path)
		case "rest":
			genReSTDocs(path)
		case "yaml":
			genYAMLDocs(path)
		case "man":
			genManDocs(path)
		default:
			if !quietMode {
				fmt.Printf("Unsupported format: %s\n", format)
			}
		}
	},
}

func init() {
	// Add subcommands to the `user` command
	genDocsCmd.Flags().StringP("format", "f", "markdown", "The format to generate documentation in")

	// Bind the flags to the viper configuration
	viper.BindPFlag("format", genDocsCmd.Flags().Lookup("format"))
}
