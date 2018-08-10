package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"flag"
	"strings"
	"bytes"
	"crypto/sha1"
)

var myPort string
var storedKey int

var myKeys []KV

func main() {
	
	port := flag.String("ip", "address", "client")  //Register node with main server
	//memory :=flag.Int("mem",10,"Memory")
	flag.Parse()
	myPort = strings.ToLower(*port)
	registerServer(myPort)	
	
													
	http.HandleFunc("/hello", requestHandler)		//Listen to port
	fmt.Printf("Starting server \n")
	
	if err := http.ListenAndServe(myPort, nil); err != nil {
		log.Fatal(err)
		return
	}
	
	fmt.Printf("Server started.. waiting for requests \n")
}

func registerServer(port string) { 
	
	storedKey=0
	port="http://localhost"+port
	port=port+"/hello"
	fmt.Printf(port)
	ur := Client{Address: port }
	br := new(bytes.Buffer)
	json.NewEncoder(br).Encode(ur)
	
	fmt.Printf("\nRegisterig service with main server:  %v\n", br)
	
	var vport string = "http://localhost:80/registerServer"
	
	req, err := http.NewRequest("GET", vport,br)
	 if err != nil {
        fmt.Printf("Error registering http.NewRequest() error: %v\n", err)
        return
    } 
	 c := &http.Client{}
	 resp, err := c.Do(req)
		if err != nil {
		fmt.Printf("Response error http.Do() error: %v\n", err)
		return
	}
	defer resp.Body.Close()
 }


func requestHandler(w http.ResponseWriter, r *http.Request) {
	
		switch r.Method {
			case "GET":
		var keyGiven string
		json.NewDecoder(r.Body).Decode(&keyGiven)
		
		for j :=0; j<storedKey ; j++{
			if myKeys[j].Key==keyGiven{
		
		
		myResponse,_ := json.Marshal(myKeys[j].Value)
		w.Header().Set("Content-Type", "application/j")
		w.WriteHeader(http.StatusOK)
		w.Write(myResponse)
			}
		}
		default:
		var u KV
		json.NewDecoder(r.Body).Decode(&u)
		hash := sha1.New()
		hash.Write([]byte(u.Key))
		u.Value = string(hash.Sum(nil))
		fmt.Printf("\n Key Stored: ")
		fmt.Printf(u.Key)
		fmt.Printf(" Value Stored: ")
		fmt.Printf(u.Value)
		myKeys = append(myKeys,u)
		storedKey++
		}
}