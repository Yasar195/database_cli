package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func CheckConnectivity(conn string) error {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}
