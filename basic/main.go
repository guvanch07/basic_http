package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//getRobots()
	//head()
	//postRobots()
	deleteRobots()
}

func getRobots() {
	resp, err := http.Get("http://www.google.com/robots.txt")

	if err != nil {
		log.Panicln(err)
	}
	// Print HTTP Status
	fmt.Println(resp.Status)

	// Read and display response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	resp.Body.Close()
}

func head() {
	resp, err := http.Head("https://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}

	resp.Body.Close()
	fmt.Println(resp.Status)
}

func postRobots() {
	form := url.Values{}
	form.Add("foo", "bar")
	resp, err := http.Post(
		"https://www.google.com/robots.txt",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)
}

func deleteRobots() {
	req, err := http.NewRequest("DELETE", "https://www.google.com/robots.txt", nil)
	if err != nil {
		log.Panicln(err)
	}
	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)
}
