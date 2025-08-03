package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GetRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "database backup utility",
		Short: "A simple cli application that handles database backups",
		Long:  "This is a command line application that will handle all the database data backup easily and securely",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to your database backup utility")
		},
	}

	return rootCmd
}
