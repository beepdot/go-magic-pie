package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type User struct {
	Image  string
	Name   string
	Dob    string
	Mobile string
	Email  string
	Stalls string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", echo)
	port := ":5000"
	log.Println("Listening on port ", port)
	http.ListenAndServe(port, mux)
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
	fmt.Printf(u.Image)
	if err := DownloadFile("/home/keshav/go/src/github.com/thedevsaddam/renderer/demo/tpl/img/avatar.jpg", fileUrl); err != nil {
		panic(err)
	}
	data := []byte(u.Image + "\n" + u.Name + "\n" + u.Dob + "\n" + u.Mobile + "\n" + u.Email + "\n" + u.Stalls + "\n")
	fmt.Printf(u.Image, u.Name, u.Dob, u.Mobile, u.Email, u.Stalls)
	writerr := ioutil.WriteFile("input.txt", data, 0777)
	if writerr != nil {
		fmt.Println(err)
	}
}
