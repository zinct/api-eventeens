package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db *sql.DB
}

func New(url string, opts ...Options) (*MySQL, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, fmt.Errorf("pkg/mysql - New - sql.Open: %w", err)
	}

	for _, opt := range opts {
		opt(db)
	}
	return &MySQL{db: db}, nil
}

func (m *MySQL) Close() error {
	return m.db.Close()
}
