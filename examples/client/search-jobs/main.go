package main

import (
	"encoding/json"
	"fmt"
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

	// Find awesome work with the Army Software Factory
	opt := usajobs.SearchOptions{
		JobCategoryCode: []string{"2210", "0854"},
	}

	r, err := c.Search.WithOptions(&opt)
	if err != nil {
		panic(err.Error())
	}

	prettyJSON, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(prettyJSON))
}
