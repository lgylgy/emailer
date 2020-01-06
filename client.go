package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"
)

const (
	rapidAPI    = "privatix-temp-mail-v1.p.rapidapi.com"
	rapidHost   = "https://privatix-temp-mail-v1.p.rapidapi.com/request"
	charset     = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"
	emailLength = 8
)

type Client struct {
	key   string
	email string
	hash  string
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
	if len(domains) == 0 {
		return domains, errors.New("no domain available")
	}
	return domains, err
}

func selectRamdonDomain(domains []string) string {
	rand.Seed(time.Now().Unix())
	return domains[rand.Intn(len(domains))]
}

func createRamdonString(length int) string {
	randomer := rand.New(rand.NewSource(time.Now().Unix()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[randomer.Intn(len(charset))]
	}
	return string(result)
}

func generateMd5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (c *Client) CreateAddress(domains []string) (string, error) {
	domain := selectRamdonDomain(domains)
	user := createRamdonString(emailLength)
	hash := generateMd5Hash(user + domain)

	cmd := fmt.Sprintf("mail/id/%s", c.hash)
	err := get("GET", rapidHost, cmd, rapidAPI, c.key,
		func(r io.Reader) error {
			return json.NewDecoder(r).Decode(&domains)
		})
	if err != nil {
		c.email = user + domain
		c.hash = hash
	}
	return c.email, nil
}
