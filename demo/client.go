package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

// User json map
type User struct {
	Image     string
	Name      string
	VisitorID string
	Stalls    string
}

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./*.html",
	}
	rnd = renderer.New(opts)
}

func slqDbQuery() *User {
	var usr User
	var max string
	fmt.Println("Querying DB for Max(id)")
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

	queryerr := db.QueryRow("SELECT max(id) FROM userdata").Scan(&max)

	if queryerr != nil {
		panic(err.Error())
	}

	getdata := db.QueryRow("SELECT image, name, visitorid, stalls FROM userdata where id=?", max).Scan(&usr.Image, &usr.Name, &usr.VisitorID, &usr.Stalls)

	if getdata != nil {
		panic(err.Error())
	}

	return &usr
}

func home(w http.ResponseWriter, r *http.Request) {
	usr := slqDbQuery()
	err := rnd.HTML(w, http.StatusOK, "home", usr)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("tpl"))
	http.Handle("/tpl/", http.StripPrefix("/tpl/", fs))
	http.HandleFunc("/", home)
	log.Println("Listening on port: 9000")
	http.ListenAndServe(":9000", nil)
}
