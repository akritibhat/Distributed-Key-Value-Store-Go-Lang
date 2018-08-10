package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"bytes"
	"flag"
	"strings"
)
type Keys []KV

func main() {
	var homeURL string = "http://localhost:80/hello"

	method := flag.String("method", "GET/POST/PUT", "method") 


	keyV := flag.String("key", "address", "client")
	flag.Parse()

	b := new(bytes.Buffer)
	if *method=="GET"{
		json.NewEncoder(b).Encode(keyV)
	}else{
		keys :=Keys{KV{Key:*keyV},
			}
	json.NewEncoder(b).Encode(keys)
	}

	fmt.Printf("Request Method: ")
	fmt.Printf(*method)
	fmt.Printf("    data: %v\n", b)

	req, err := http.NewRequest(strings.ToUpper(*method), homeURL, b)
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("http.Do() error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

			    fmt.Println(string(resp_body))
	fmt.Printf("Request Processed")
	}