package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, ""+
			`Generate tempory email address and fetch messages using TEMPMAIL APi.
`)
		flag.PrintDefaults()
	}
	key := flag.String("key", "", "RapidAPI application key")
	frequency := flag.Int("frequency", 20, "email fetching frequency (seconds)")
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

	log.Println("Generate random address:")
	email, err := client.CreateAddress(domains)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("-> ", email)

	log.Println("Fetch email:")
	ticker := time.NewTicker(time.Second * time.Duration(*frequency))
	defer ticker.Stop()
	for range ticker.C {
		msg, emails, err := client.FetchEmail()
		if err != nil {
			log.Fatal(err)
		}
		if len(emails) != 0 {
			log.Printf("*** %d emails received\n", len(emails))
		} else {
			log.Println("-> ", msg)
		}
	}
}
