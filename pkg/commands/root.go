package commands

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Args:  cobra.MaximumNArgs(1),
	Use:   "dockerenv [image name]",
	Short: "Look up a Docker image",
	Long:  `Quickly look up supported environment variables for popular Docker images`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		fmt.Println(fmt.Sprintf("Environment variables for %s:\n", text.Bold.Sprint(args[0])))

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)

		t.AppendRow(table.Row{
			"POSTGRES_USER",
			"required",
			"does stuff",
		})

		t.Render()

		fmt.Println(text.Faint.Sprint("\nView on Docker Hub: https://hub.docker.com/_/mongo"))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
