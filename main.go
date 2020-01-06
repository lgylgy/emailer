package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, ""+
			`Generate tempory email address and fetch messages using TEMPMAIL APi.
`)
		flag.PrintDefaults()
	}
	key := flag.String("key", "", "RapidAPI application key")
	flag.Parse()

	client := NewClient(*key)

	log.Println("List Domains:")
	domains, err := client.ListDomains()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range domains {
		log.Println("-> ", v)
	}
}
