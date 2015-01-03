package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var conninfo string = "user=postgres password=postgres host=91.191.168.68 port=5432 dbname=tahmin-oyun sslmode=disable"

var selectStatement = `
	 select user_id, user_fb_id from tbl_user 
`

func main() {
	fmt.Printf("testing..\n")

	var db *sql.DB
	var err error

	db, err = sql.Open("postgres", conninfo)

	if err != nil {
		fmt.Printf("sql.Open error: %v\n", err)
		return
	}

	defer db.Close()

	err = doSelect(db)
	if err != nil {
		fmt.Printf("select error: %v\n", err)
		return
	}
}

func doSelect(db *sql.DB) error {
	var stmt *sql.Stmt
	var err error

	stmt, err = db.Prepare(selectStatement)
	if err != nil {
		fmt.Printf("db.Prepare error: %v\n", err)
		return err
	}

	var rows *sql.Rows

	rows, err = stmt.Query()
	if err != nil {
		fmt.Printf("stmt.Query error: %v\n", err)
		return err
	}

	defer stmt.Close()

	for rows.Next() {
		var firstname string
		var lastname string

		err = rows.Scan(&firstname, &lastname)
		if err != nil {
			fmt.Printf("rows.Scan error: %v\n", err)
			return err
		}

		fmt.Printf("firstname: %v lastname: %v \n",
			firstname, lastname)
	}

	return nil
}
