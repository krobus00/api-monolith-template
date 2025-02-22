/*
Copyright © 2024 Michael Putera Wardana <michaelputeraw@gmail.com>
*/
package cmd

import (
	"github.com/api-monolith-template/internal/bootstrap"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "perform database migration",
	Long:  `perform database migration`,
	Run: func(cmd *cobra.Command, args []string) {
		action, _ := cmd.Flags().GetString("action")
		migrationName, _ := cmd.Flags().GetString("name")
		version, _ := cmd.Flags().GetInt64("version")

		bootstrap.StartMigrate(action, migrationName, &version)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.PersistentFlags().String("action", "up", "action create|up|up-by-one|up-to|down|down-to|reset|status")
	migrateCmd.PersistentFlags().Int64("version", 1, "version")
	migrateCmd.PersistentFlags().String("name", "", "migration name")
}
