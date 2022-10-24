package main

import (
	"database/sql"
	"fmt"

	_ "github.com/sijms/go-ora/v2"
)

// connParam Oracle username/password connection
var connParam = map[string]string{
	"user": "system",
	"pass": "myPassword",
	"sid":  "mySID",
	"port": "1521",
	"host": "192.168.0.0",
}

func main() {
	// Oracle string connection "oracle://<user>:<pass>@<host>:<port>/<sid>"
	connString := "oracle://" + connParam["user"] + ":" + connParam["pass"] + "@" + connParam["host"] + ":" + connParam["port"] + "/" + connParam["sid"]

	// Open connection
	db, err := sql.Open("oracle", connString)
	if err != nil {
		panic(fmt.Errorf("error in sql.Open %w", err))
	}
	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Println("Can't close connection: ", err)
		}
	}()

	// Test db connection
	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging db: %w", err))
	}

	// Executing query on database
	rows, err := db.Query("SELECT sysdate FROM dual")
	if err != nil {
		fmt.Println("Error running query: ", err)
	}
	defer rows.Close()

	// Get string from affected rows
	var theDate string
	for rows.Next() {
		rows.Scan(&theDate)
	}
	fmt.Printf("The date is: %s\n", theDate)
}
