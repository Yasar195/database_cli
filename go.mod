module database_backup_tool

go 1.24.4

replace database_backup_tool/cmds => ./cmds

require database_backup_tool/cmds v0.0.0-00010101000000-000000000000

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cobra v1.9.1 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
)
