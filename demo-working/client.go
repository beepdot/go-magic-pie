package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

// User json map
type User struct {
	Image  string
	Name   string
	Dob    string
	Mobile string
	Email  string
	Stalls string
}

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./*.html",
	}
	rnd = renderer.New(opts)
}

func readFileWithReadString(fn string) *User {
	var line string
	var usr User
	file, err := os.Open(fn)
	defer file.Close()
	if err != nil {
		fmt.Printf("Error")
		usr.Image = "https://st3.depositphotos.com/1029662/14068/v/1600/depositphotos_140681792-stock-illustration-website-error-500-internal-server.jpg"
		usr.Name = "Exception"
		usr.Dob = "01/01/1900"
		usr.Mobile = "0000000000"
		usr.Email = "admin@devops.org"
		usr.Stalls = "0000"
		return &usr
	}
	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var count = 0
	for {
		line, err = reader.ReadString('\n')
		//		fmt.Printf(line)
		if count == 0 {
			usr.Image = strings.Trim(line, "\n")
		} else if count == 1 {
			usr.Name = strings.Trim(line, "\n")
		} else if count == 2 {
			usr.Dob = strings.Trim(line, "\n")
		} else if count == 3 {
			usr.Mobile = strings.Trim(line, "\n")
		} else if count == 4 {
			usr.Email = strings.Trim(line, "\n")
		} else if count == 5 {
			usr.Stalls = strings.Trim(line, "\n")
		} else {
		}
		if err != nil {
			break
		}
		count++
	}
	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}
	return &usr
}

func home(w http.ResponseWriter, r *http.Request) {
	usr := readFileWithReadString("input.txt")
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
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", home)
	//port := ":9000"
	//log.Println("Listening on port ", port)
	//http.ListenAndServe(port, mux)
}
