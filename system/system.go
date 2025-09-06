package systems

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CheckForUpdates() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "System update command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Checking for updates...")
		},
	}
}
