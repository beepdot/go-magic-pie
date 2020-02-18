package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Image     string
	Name      string
	VisitorID string
	Stalls    string
}

func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func echo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var u User
	err = json.Unmarshal(body, &u)
	fileUrl := u.Image
	if err := DownloadFile("/home/keshav/go/src/github.com/thedevsaddam/renderer/demo/tpl/img/"+u.VisitorID+".jpg", fileUrl); err != nil {
		panic(err)
	}

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
	insert, err := db.Query("INSERT INTO userdata(image, name, visitorid, stalls) VALUES (?, ?, ?, ?)", u.Image, u.Name, u.VisitorID, u.Stalls)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", echo)
	port := ":5000"
	log.Println("Listening on port ", port)
	http.ListenAndServe(port, mux)
}
