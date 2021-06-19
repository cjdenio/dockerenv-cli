package cmd

import "github.com/spf13/cobra"

// Root command for the cli
var RootCmd = &cobra.Command{
	Use:  "denv",
	Long: "CLI view for dockerenv",
}
