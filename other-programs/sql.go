package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Inserting data to DB")
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/devcon")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO userdata(image, name, visitorid) VALUES ('https://cdn.vox-cdn.com/thumbor/xuvr8C1o146_anbM4w1JMioGYIM=/0x0:1200x696/920x613/filters:focal(535x82:727x274):format(webp)/cdn.vox-cdn.com/uploads/chorus_image/image/58843699/baby-groot-guardians.0.0.jpg', 'KP', 'OK' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
