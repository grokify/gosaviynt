package main

/*
2024/01/21 16:48:06 405 Method Not Allowed
exit status 1
johnwang@MacBook-Pro-6 simple % go run main.go
2024/01/21 16:51:50 412 Precondition Failed
*/

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/saviynt/gosaviynt"
	"github.com/saviynt/gosaviynt/saviynt"
)

func main() {
	var paramBaseURL = flag.String("baseurl", "https://example.saviyntcloud.com", "base URL. Do not add trailing slash")
	var paramUsername = flag.String("username", "admin", "username")
	var paramPassword = flag.String("password", "changeme", "password")
	var paramAnalyticsName = flag.String("analyticsname", "changeme", "analyticsname")

	flag.Parse()

	client, err := gosaviynt.NewClient(*paramBaseURL, *paramUsername, *paramPassword)
	if err != nil {
		log.Fatal(err)
	}

	apiURL := *paramBaseURL + gosaviynt.RelURLECM + gosaviynt.RelURLAPI

	apiClient := saviynt.NewAPIClient(&saviynt.Configuration{
		Servers:    saviynt.ServerConfigurations{{URL: apiURL}},
		HTTPClient: client})

	req := apiClient.AnalyticsAPI.FetchRuntimeControlsDataV2(context.Background())

	r2 := saviynt.FetchRuntimeControlsDataV2Request{
		Analyticsname: paramAnalyticsName,
		Attributes:    map[string]any{"timeFrame": "10000"},
		Max:           gosaviynt.Pointer("50"),
		Offset:        gosaviynt.Pointer("0"),
	}
	fmt.Printf("%v\n", r2)

	req2 := req.FetchRuntimeControlsDataV2Request(r2)

	resp, err := req2.Execute()
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(b))

	fmt.Println("DONE")
}
