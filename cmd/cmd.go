package cmd

import (
	"github.com/lcok/vault-sync/internal/app"
	"github.com/spf13/cobra"
	"log"
)

// Execute cmd entrance
func Execute() error {
	rootCmd.AddCommand(serverCmd, taskCmd, encCmd)
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "vault-sync",
	Short: "vault-sync is a lightweight backup tool designed for Vaultwarden.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			log.Fatal(err)
		}
	},
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a continuously running backup handler that periodically performs backup tasks based on environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		server, err := app.InitServer()
		if err != nil {
			log.Fatal(err)
		}
		server.Start()
	},
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Performs a one-time task to perform a local or remote backup;",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := app.InitTask()
		if err != nil {
			log.Fatal(err)
		}
		client.Start()
	},
}

var encCmd = &cobra.Command{
	Use:   "enc",
	Short: "Performs an encryption or decryption of a backup",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO impl this
	},
}
