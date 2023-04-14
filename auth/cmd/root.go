package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "auth",
	Short: "Auth service",
	Long: `Auth services that interfaces with the migration service, frontend, and keycloak
to provide authentication and authorization services.`,
}

func GetRootCommand() *cobra.Command {

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(getServeCommand())
	return rootCmd
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
