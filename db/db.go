package db

import (
	"database/sql"
	"fmt"
	"go-url-shrtr/utils"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error

	if DB, err = sql.Open("mysql", utils.DSN); err != nil {
		panic("Could not connect to databse" + err.Error())
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	createTable()

	fmt.Println("Successfully connected to MySQL!")
}

func createTable() {
	createUrlTable := `CREATE TABLE IF NOT EXISTS urls(
	id INTEGER PRIMARY KEY AUTO_INCREMENT,
	code VARCHAR(255),
	url VARCHAR(255),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	click_count INTEGER
	);`

	if _, err := DB.Exec(createUrlTable); err != nil {
		panic("Could not create url_shrtr table." + err.Error())
	}
}
