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
		"https://1337x.to/",
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
	if res, err := http.Get(website); err != nil {
		fmt.Println(website, "is down")
		c <- "The site is down"

	} else {
		fmt.Printf("[%d] %s is up\n", res.StatusCode, website)
		c <- "The site is up"
	}

}
