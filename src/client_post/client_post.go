package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"bytes"
)

const (
	homeURL string = "http://localhost:80/hello"
)

type Keys []KV

func main() {

	keys := Keys{
        KV{Key: "Tiger"},
        {Key: "Fish"},
        {Key: "Cat"},
        {Key: "Dog"},
        {Key: "Elephant"},
        {Key: "Giraffe"},
        {Key: "Shark"},
        {Key: "Buffalo"},
        {Key: "Crow"},
        {Key: "Tiny"},
        {Key: "Owl"},
        {Key: "Panda"},
        {Key: "Camel"},
        {Key: "Zebra"},
        {Key: "Horse"},
        {Key: "Lion"},
    }
	
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(keys)
	
	fmt.Printf("Keys Sent: %v\n", b)

	req, err := http.NewRequest("POST", homeURL, b)
	if err != nil {
		fmt.Printf("Sending Failed, http.NewRequest() error: %v\n", err)
		return
	}

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Response error : http.Do() error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll() error: %v\n", err)
		return
	}

	fmt.Printf("\n%v\n", string(data))
}