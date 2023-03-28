package main

import (
	"database/sql"
	"fmt"

	_ "github.com/sijms/go-ora/v2"
)

func main() {

	connString := "oracle://s101598:o8N2C5Q259@217.173.198.135:1521/tpdb"

	db, err := sql.Open("oracle", connString)
	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}

	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Println("Can't close connection: ", err)
		}
	}()

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging db: %w", err))
	}

}
