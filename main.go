package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

type HttpDecoder interface {
	RetrieveRequestOrBody(*http.Response) ([]byte, error)
}

type DataFormat struct {
	Description string `json:"descriere"`
}

// https://data.europa.eu/data/datasets/fbeab450-4092-40c7-b0ef-ce0222e6c17b?locale=en
const DataLink string = "https://data.gov.ro/dataset/fbeab450-4092-40c7-b0ef-ce0222e6c17b/resource/3cd79e1f-6b75-4b6f-bf32-903e2b6fad51/download/wifi_cluj.json"

func main() {
	httpReq := HttpGetter(DataLink)
	defer httpReq.Body.Close()
	data, err := RetrieveRequest(httpReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	dataReturned, err := RetrieveBody(httpReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dataReturned))

	// print only the 'descriere' field
	if err := ExtractLocations(dataReturned); err != nil {
		log.Fatal(err)
	}
}

func ExtractLocations(dataReturned []byte) error {
	var locations []DataFormat
	err := json.Unmarshal(dataReturned, &locations)
	if err != nil {
		log.Fatal(err)
	}
	for _, loc := range locations {
		fmt.Println(loc.Description)
	}
	return nil
}

// RetrieveRequest dumps the HTTP Request
func RetrieveRequest(req *http.Response) ([]byte, error) {
	reqDump, err := httputil.DumpResponse(req, false)
	if err != nil {
		log.Fatal(err)
	}
	return reqDump, nil
}

func HttpGetter(link string) *http.Response {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := client.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	return req
}

// RetrieveBody dumps the body of an HTTP request
func RetrieveBody(req *http.Response) ([]byte, error) {
	respDump, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Failed to read body: %v", err)
	}
	return respDump, nil
}
