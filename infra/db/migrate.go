package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(db *sqlx.DB, dir string) error {
	migrations := &migrate.FileMigrationSource{
		Dir: dir,
	}

	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	if n == 0 {
		fmt.Println("No new migrations to apply")
		return nil
	}

	fmt.Printf("Applied %d migrations\n", n)
	return nil
}
