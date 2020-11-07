package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://i.imgur.com/7RRXZBR.jpg")
	if err != nil {
		log.Fatalf("Failed to fetch the image: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read all of the image: %v", err)
	}
	err = ioutil.WriteFile("cat.jpg", body, 0644)
	if err != nil {
		log.Fatalf("Failed to write the image to a file: %v", err)
	}
}
