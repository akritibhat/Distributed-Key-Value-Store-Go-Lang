package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

var processors[5] string
var totalProcessor int
var request int
var nodes []Client
type Keys []KV
var mod int

type KV struct{
    Key      string
    Value 	 string
}

type Client struct{
	Address string

}

func main() {

	http.HandleFunc("/hello", requestHandler)
	http.HandleFunc("/registerServer", newServerHandler)
	
	fmt.Printf("Starting Master Server\n")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Started Master Server, waiting for requests\n")
}

func newServerHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("in server handler")
		var u Client
		json.NewDecoder(r.Body).Decode(&u)
		fmt.Printf(u.Address)
		totalProcessor++
		nodes= append(nodes,u)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
		if totalProcessor==0{
			fmt.Printf("\n No servers available to process. Please try Later")
			return
		}
		switch r.Method {
			case "GET":
		var keyGiven string
		fmt.Printf(keyGiven)
		json.NewDecoder(r.Body).Decode(&keyGiven)
		for i := 0; i<totalProcessor ;i++ {
			b := new(bytes.Buffer)
					json.NewEncoder(b).Encode(keyGiven)	
					req, err := http.NewRequest("GET", nodes[i].Address, b)
					if err != nil {
					fmt.Printf("Error Error http.NewRequest() error: %v\n", err)
					return
					}

					c := &http.Client{}
					resp, err := c.Do(req)
					if err != nil {
					fmt.Printf("GET error http.Do() error: %v\n", err)
					return
				}
				defer resp.Body.Close()
				resp_body, _ := ioutil.ReadAll(resp.Body)

			    		 fmt.Println(resp.Status)
			    fmt.Println(string(resp_body))
			    w.Header().Set("Content-Type", "application/j")

				w.Write([]byte(resp_body))
				if resp.Status == "OK"{
					break
				return
				}
		}
		
		
		
		default:
		
		var keys Keys
		json.NewDecoder(r.Body).Decode(&keys)
	
		var length int  =len(keys)
		fmt.Printf("\n Received Keys: %d\n",length)
		
	//	var keysNeeded int = length
		request=0
		
		for i :=0; i<length ;i++{
			request++
			fmt.Printf("Processing request %d\n", request)
			mod=request % totalProcessor
				fmt.Printf("going to server %d\n", mod)
		
					b := new(bytes.Buffer)
					json.NewEncoder(b).Encode(keys[i])	
					req, err := http.NewRequest("POST", nodes[mod].Address, b)
					if err != nil {
					fmt.Printf("Error sending http.NewRequest() error: %v\n", err)
					return
					}
					c := &http.Client{}
					resp, err := c.Do(req)
					if err != nil {
					fmt.Printf("Post error from node server http.Do() error: %v\n", err)
					return
				}
				defer resp.Body.Close()
			//	nodes[i].Free=nodes[i].Free-1
			//	nodes[i].Used=nodes[i].Used+1
			} 
		}
		
		
	/*	for i := 0; i<totalProcessor ;i++ {
			
			if nodes[i].Free>=keysNeeded {
				for j :=0; j<keysNeeded ; j++{
						b := new(bytes.Buffer)
					json.NewEncoder(b).Encode(keys[keyCounter])	
					keyCounter++
					req, err := http.NewRequest("POST", nodes[i].Address, b)
					if err != nil {
					fmt.Printf("Error sending http.NewRequest() error: %v\n", err)
					return
					}
					
					c := &http.Client{}
					resp, err := c.Do(req)
					if err != nil {
					fmt.Printf("Post error from node server http.Do() error: %v\n", err)
					return
				}
				defer resp.Body.Close()
			}	
				nodes[i].Free=nodes[i].Free-keysNeeded
				nodes[i].Used=nodes[i].Used+keysNeeded
			
				break
			} else {
				
			keysNeeded = keysNeeded-nodes[i].Free
					for j :=0; j<nodes[i].Free ; j++{
						b := new(bytes.Buffer)
					json.NewEncoder(b).Encode(keys[keyCounter])	
					keyCounter++
					req, err := http.NewRequest("POST", nodes[i].Address, b)
					if err != nil {
					fmt.Printf("Post error from node server http.NewRequest() error: %v\n", err)
					return	
					}
					
					c := &http.Client{}
					resp, err := c.Do(req)
					if err != nil {
					fmt.Printf("Post response error from node serve http.Do() error: %v\n", err)
					return
					}
				defer resp.Body.Close()
				}
					nodes[i].Used=nodes[i].Used+nodes[i].Free
					nodes[i].Free=0
					
			}
		}*/
		
	}

