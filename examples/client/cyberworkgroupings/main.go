package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	usajobs "github.com/JeffRDay/go-usajobs/client"
)

func main() {
	userAgent := os.Getenv("EMAIL")
	token := os.Getenv("TOKEN")

	c, err := usajobs.NewClient(userAgent, token)
	if err != nil {
		panic(err.Error())
	}

	httpResponse, r, err := c.CyberWorkGroupings.WithOptions(nil)
	if err != nil {
		panic(err.Error())
	}

	if httpResponse.StatusCode != http.StatusOK {
		fmt.Printf("received non-200 response code: %d\n", httpResponse.StatusCode)
		os.Exit(1)
	}

	prettyJSON, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(prettyJSON))
}
