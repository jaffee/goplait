package main

import (
	"fmt"
	"github.com/jaffee/github"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	api := github.Api{"jaffee"}
	user, err := api.User()
	check(err)
	fmt.Printf("%v", user)
}
