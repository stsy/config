package main

import (
	"fmt"
	"log"

	"github.com/stsy/config"
)

func main() {
	conf, err := config.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Prints "world" form key "hello"
	fmt.Println(conf["hello"])
}
