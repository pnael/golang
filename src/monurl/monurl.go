package main

import (
	"hash/fnv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// LoopEverySeconds Loop every 5s
const LoopEverySeconds = 5

var url string
var cookies string
var client = &http.Client{}
var lastData string

func makeRequest() {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	//Set Cookies if they were passed as an argument
	if cookies != "" {
		req.Header.Set("Cookie", cookies)
	}

	//Send requests
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	//Save response body into data variable
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// if lastdata is equal to "", it means that it is
	// the first request and we set lastdata to current
	// response Body
	// Otherwise , we compare previous and current html
	if lastData == "" {
		lastData = string(data)
	} else {
		checkChanges(string(data))
		lastData = string(data)
	}
}

func scheduleRequest() {
	// Infinite loop to run every LoopEverySeconds seconds
	for range time.Tick(time.Second * LoopEverySeconds) {
		makeRequest()
	}
}

func checkChanges(newData string) {
	if newData == lastData {
		log.Println("No changes")
		return
	}
	h := fnv.New32a()
	h.Write([]byte(lastData))
	log.Println(h.Sum32())

	log.Println("Changes noticed")
	os.Exit(0)
}

func main() {
	// check if there is an URl
	if len(os.Args) < 2 {
		log.Println("Missing URL")
		os.Exit(0)
	}

	//set URL
	url = os.Args[1]

	//set Cookie if an argument
	if len(os.Args) == 3 {
		cookies = os.Args[2]
	}

	// Make initial request
	makeRequest()
	//Continue with requests. Loop every LoopEverySeconds seconds
	scheduleRequest()
}
