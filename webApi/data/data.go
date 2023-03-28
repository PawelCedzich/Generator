package data

import (
	"database/sql"
	"fmt"

	_ "github.com/sijms/go-ora/v2"
)

type Dane struct {
	val1 string
}

type BodyData struct {
	Val1, Val2 string
	Option     any
}

func newData(val1 string) Dane {
	return Dane{
		val1: val1,
	}
}

func DbGenerate(bodyData BodyData) {

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

	fmt.Println(bodyData.Val1)
	fmt.Println(bodyData.Val2)
	fmt.Println(bodyData.Option)
	//rows, err := db.Query("SELECT table_name, s101598 FROM user_tables ORDER BY owner, table_name")
	//defer rows.Close()

}
