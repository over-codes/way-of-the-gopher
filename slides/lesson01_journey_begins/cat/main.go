package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://i.imgur.com/7RRXZBR.jpg")
	body, _ := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile("cat.jpg", body, 0644)
}
