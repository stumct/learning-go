package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetImgurAPI(url string, clientid string) *Response {
	// Create a request and add the proper headers.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//return nil, err
	}
	req.Header.Set("Authorisation", "Client-ID "+clientid)
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

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	return &response
}

type Response struct {
	Data []struct {
		Image ImageData
	}
}

type ImageData struct {
	Id            string
	Title         string
	Description   string
	DateTime      int32
	Type          string
	Animated      bool
	Width         int
	Height        int
	Size          int
	View          int
	Bandwidth     int
	NSFW          bool
	Section       string
	Is_Ad         bool
	In_Gallery    bool
	Link          string
	Comment_Count int
	Ups           int
	Downs         int
	Points        int
	Score         int
	Is_Album      bool
}
