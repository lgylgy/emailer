package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func get(verb, host, path, api, key string,
	decode func(io.Reader) error) error {

	u := fmt.Sprintf("%s/%s", host, path)
	log.Println(u)
	log.Println(key)
	rq, err := http.NewRequest(verb, u, nil)
	if err != nil {
		return err
	}
	rq.Header.Set("Content-Type", "application/json")
	rq.Header.Set("x-rapidapi-host", api)
	rq.Header.Set("x-rapidapi-key", key)

	client := http.Client{}
	rsp, err := client.Do(rq)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	var body io.Reader = rsp.Body
	if rsp.StatusCode != http.StatusOK {
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return err
		}
		result := map[string]string{}
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			return err
		}
		text, ok := result["error"]
		if ok {
			return errors.New(text)
		}
		return errors.New(result["message"])
	}
	return decode(body)
}
