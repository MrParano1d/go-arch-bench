package plainsql

import (
	"database/sql"
	"fmt"

	serverstd "github.com/mrparano1d/archs-go/pkg/server_std"
	"github.com/mrparano1d/archs-go/pkg/session"

	_ "github.com/lib/pq"
)

func Serve(dsn string) error {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	sessionStorage := session.NewInMemoryStorage()
	authStorage := NewSqlAuthStorage(db)
	booksStorage := NewSqlBooksStorage(db)



}
