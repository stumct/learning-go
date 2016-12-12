package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const UserAgent = "Golang Reddit Reader"

func main() {
	// Create a request and add the proper headers.
	req, err := http.NewRequest("GET", "https://api.reddit.com/r/aww", nil)
	if err != nil {
		//return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	// Handle the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		//return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		//return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error reading body")
	}
	//fmt.Println(string(body))

	var document Response
	err = json.Unmarshal(body, &document)
	if err != nil {
		panic(err)
	}

	fmt.Println(document)

	defer resp.Body.Close()

}

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

type Item struct {
	Title string
	URL   string
	Score int64
}
