package main

import (
	"fmt"
	"goevents/config"
	"goevents/pkg/mysql"

	migrateMySQL "github.com/golang-migrate/migrate/v4/database/mysql"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	// Initialize MYSQL
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.MYSQL.Username, cfg.MYSQL.Password, cfg.MYSQL.Host, cfg.MYSQL.Port, cfg.MYSQL.Database)
	mysql, err := mysql.New(mysqlUrl)
	if err != nil {
		panic(err)
	}
	defer mysql.Close()

	// Migrate
	_, err = migrateMySQL.WithInstance(mysql.DB, &migrateMySQL.Config{})
	if err != nil {
		panic(err)
	}

	// m, err := migrate.NewWithDatabaseInstance(
	// 	"file://infrastructure/db/migrations",
	// 	"mysql", driver)

	// if
}
