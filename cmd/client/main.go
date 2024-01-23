package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/saviynt/gosaviynt"
)

func main() {
	var paramBaseURL = flag.String("baseurl", "https://example.saviyntcloud.com", "base URL")
	var paramUsername = flag.String("username", "admin", "usernaem")
	var paramPassword = flag.String("password", "changeme", "password")

	flag.Parse()

	fmt.Printf("logging into (%s) with user (%s)\n", *paramBaseURL, *paramUsername)

	tok, err := gosaviynt.FetchToken(*paramBaseURL, *paramUsername, *paramPassword)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Retrieved Access Token: %s\n", tok.AccessToken)

	fmt.Println("DONE")
}

func pointer[E any](e E) *E { return &e }
