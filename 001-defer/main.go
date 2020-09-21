package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.StatusCode, res.Status, res.ContentLength)

	defer res.Body.Close() // this will not change the value of res.Body even though this execute at last
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", robots)
}
