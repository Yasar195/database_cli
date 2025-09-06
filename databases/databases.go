package databases

import (
	"database_backup_tool/postgres"
	"fmt"

	"github.com/spf13/cobra"
)

var supported_databases = []string{
	"postgreSQL",
}

var cliMode bool

func GetSupportedDatabases() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Lists all supported database engines",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Supported database engines:")

			for _, db := range supported_databases {
				fmt.Println(db)
			}
		},
	}
}

func CheckDatabaseConnectivity() *cobra.Command {
	return &cobra.Command{
		Use:   "connect",
		Short: "Check database connectivity",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Checking database connectivity...")
			if len(args) > 0 {
				err := postgres.CheckConnectivity(args[0])

				if err != nil {
					fmt.Println("Failed to connect to database")
				} else {
					fmt.Println("Successfully connected to database")
				}
			} else {
				fmt.Println("Please provide a connection string as an argument.")
			}
		},
	}
}
