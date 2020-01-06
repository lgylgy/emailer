package main

import (
	"encoding/json"
	"io"
)

const (
	rapidAPI  = "privatix-temp-mail-v1.p.rapidapi.com"
	rapidHost = "https://privatix-temp-mail-v1.p.rapidapi.com/request"
)

type Client struct {
	key string
}

func NewClient(key string) *Client {
	return &Client{
		key: key,
	}
}

func (c *Client) ListDomains() ([]string, error) {
	domains := []string{}
	err := get("GET", rapidHost, "domains/", rapidAPI, c.key,
		func(r io.Reader) error {
			return json.NewDecoder(r).Decode(&domains)
		})
	return domains, err
}
