package main

import (
	"database_backup_tool/cmds"
)

func main() {
	rootCmd := cmds.GetRootCmd()
	rootCmd.Execute()
}
