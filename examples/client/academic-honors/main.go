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

	_, r, err := c.AcademicHonors.WithOptions(nil)
	if err != nil {
		panic(err.Error())
	}

	prettyJSON, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(prettyJSON))
}
