package cmds

import (
	"database_backup_tool/databases"
	systems "database_backup_tool/system"
	"fmt"

	"github.com/spf13/cobra"
)

func GetRootCmd() *cobra.Command {

	var rootCmd = &cobra.Command{
		Use:   "dBackup",
		Short: "A simple cli application that handles database backups",
		Long:  "This is a command line application that will handle all the database data backup easily and securely",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to your database backup utility")
		},
	}

	rootCmd.AddCommand(getDatabasecmd())
	rootCmd.AddCommand(getSystemFunctions())

	return rootCmd
}

func getDatabasecmd() *cobra.Command {
	dbCmd := &cobra.Command{
		Use:   "db",
		Short: "Database related cli utilities",
	}

	dbCmd.AddCommand(databases.GetSupportedDatabases())
	dbCmd.AddCommand(databases.CheckDatabaseConnectivity())

	return dbCmd
}

func getSystemFunctions() *cobra.Command {
	systemCmd := &cobra.Command{
		Use:   "system",
		Short: "Internal system commands",
	}

	systemCmd.AddCommand(systems.CheckForUpdates())

	return systemCmd
}
