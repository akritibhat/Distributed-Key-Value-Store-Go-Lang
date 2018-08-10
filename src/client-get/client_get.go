package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"bytes"
)

func main() {
	var homeURL string = "http://localhost:80/hello"
	var keys string
	keys="Tiger"
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(keys)
	
	fmt.Printf("Key Searching: %v\n", b)

	req, err := http.NewRequest("GET", homeURL, b)
	if err != nil {
		fmt.Printf("GET error http.NewRequest() error: %v\n", err)
		return
	}

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Response error http.Do() error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println(resp.Status)
			    fmt.Println(string(resp_body))

	fmt.Printf("Request Processed")
}