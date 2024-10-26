package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

type HttpDecoder interface {
	RetrieveRequestOrBody(string) ([]byte, error)
}

const DataLink string = "https://data.gov.ro/dataset/fbeab450-4092-40c7-b0ef-ce0222e6c17b/resource/3cd79e1f-6b75-4b6f-bf32-903e2b6fad51/download/wifi_cluj.json"

func main() {
	data, err := RetrieveRequest(DataLink)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func RetrieveRequest(link string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		log.Fatal(err)
	}
	reqDump, err := httputil.DumpRequest(req, false)
	if err != nil {
		log.Fatal(err)
	}
	return reqDump, nil
}

func RetrieveBody() {
	panic("not implemented")
}
