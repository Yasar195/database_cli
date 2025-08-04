package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var supported_databases = []string{
	"postgreSQL",
}

func GetRootCmd() *cobra.Command {

	var rootCmd = &cobra.Command{
		Use:   "dBackup",
		Short: "A simple cli application that handles database backups",
		Long:  "This is a command line application that will handle all the database data backup easily and securely",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to your database backup utility")
		},
	}

	rootCmd.AddCommand(getSupportedDatabases())
	rootCmd.AddCommand(getSystemFunctions())

	return rootCmd
}

func getSupportedDatabases() *cobra.Command {
	dbCmd := &cobra.Command{
		Use:   "db",
		Short: "Databse related cli utilities",
	}

	dbCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List's all the supported database engines",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Supported database engines:")
			for _, db := range supported_databases {
				fmt.Println(db)
			}
		},
	})

	return dbCmd
}

func getSystemFunctions() *cobra.Command {
	systemCmd := &cobra.Command{
		Use:   "system",
		Short: "Internal system commands",
	}

	systemCmd.AddCommand(&cobra.Command{
		Use:   "update",
		Short: "System update command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Checking for updates...")
		},
	})

	return systemCmd
}
