package main

import (
	"fmt"
	"net/http"
)

func main() {

	c := make(chan string)

	websites := []string{
		"https://stackoverflow.com/",
		"https://github.com/",
		"https://www.linkedin.com/",
		"http://medium.com/",
		"https://golang.org/",
		"https://www.udemy.com/",
		"https://www.coursera.org/",
		"https://wesionary.team/",
	}

	for _, website := range websites {
		go getWebsite(website, c)
	}

	for msg := range c {
		fmt.Println(msg)
	}

}
func getWebsite(website string, c chan string) {
	if _, err := http.Get(website); err != nil {
		c <- website + "is down"

	} else {
		c <- website + "is up"
	}

}
