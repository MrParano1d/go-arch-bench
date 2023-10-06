package migrate

import (
	"database/sql"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Up(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	usersSql, err := os.ReadFile("./migrations/users.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(usersSql))
	if err != nil {
		return err
	}

	booksSql, err := os.ReadFile("./migrations/books.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(booksSql))
	if err != nil {
		return err
	}

	return nil

}
