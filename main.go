package main

import (
	"fmt"
	"net/http"
)

func main() {

	websites := []string{
		"https://stackoverflow.com/",
		"https://github.com/",
		"http://medium.com/",
		"https://golang.org/",
		"https://www.udemy.com/",
		"https://www.coursera.org/",
		"https://wesionary.team/",
	}

	for _, website := range websites {
		getWebsite(website)
	}

}
func getWebsite(website string) {
	if res, err := http.Get(website); err != nil {
		fmt.Println(website, "is down")

	} else {
		fmt.Printf("[%d] %s is up\n", res.StatusCode, website)
	}

}
