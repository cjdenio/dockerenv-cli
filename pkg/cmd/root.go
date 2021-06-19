package cmd

import (
	"fmt"

	"github.com/cjdenio/dockerenv-cli/pkg/api"
	"github.com/spf13/cobra"
)

// Root command for the cli
var RootCmd = &cobra.Command{
	Use:                   "denv <IMAGE NAME>",
	Example:               "denv postgres\ndenv node",
	Long:                  "CLI view for dockerenv",
	Args:                  cobra.ExactValidArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := api.Image(args[0])
		if err != nil {
			return err
		}

		// Outputting URL
		dontIncludeURL, err := cmd.Flags().GetBool("url")
		if err != nil {
			return err
		}
		if !dontIncludeURL {
			fmt.Println("URL:", data.URL)
		}

		// Outputting variables
		dontIncludeVariables, err := cmd.Flags().GetBool("variables")
		if err != nil {
			return err
		}
		if !dontIncludeVariables {
			fmt.Println("Variables:")
			fmt.Println()
			for _, variable := range data.Variables {
				fmt.Println("\tName:", variable.Name)
				fmt.Println("\tDescription:", variable.Description)
				fmt.Println("\tDefault:", variable.Default)
				fmt.Println("\tRequired:", variable.Required)
				fmt.Println("\tUncommon:", variable.Uncommon)
				fmt.Println()
			}
		}

		return nil
	},
}

func init() {
	RootCmd.Flags().BoolP("url", "u", false, "Only output the url for the image")
	RootCmd.Flags().BoolP("variables", "v", false, "Only output the variables of the image")
}
