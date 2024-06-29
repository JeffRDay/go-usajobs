package cmd

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	usajobs "github.com/JeffRDay/go-usajobs/client"
)

var academichonorsTestData = "../../testdata/academichonors-testdata.json"

func TestListAcademicHonors(t *testing.T) {
	// Read the JSON file from testdata directory
	file, err := os.Open(academichonorsTestData)
	if err != nil {
		t.Fatalf("could not open test data: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("could not read test data: %v", err)
	}

	// Create a mock server that returns the JSON data
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Log(r.URL.String())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}))
	defer mockServer.Close()

	u, err := url.Parse(mockServer.URL)
	if err != nil {
		t.Fatalf("failed to parse mock server url: %v", err)
	}

	Client, err = usajobs.NewClient("test", "test")
	if err != nil {
		panic(err.Error())
	}
	Client.BaseURL = u

    results, err := executeListAcademicHonors(nil)
	if err != nil {
		t.Fatal(err.Error())
	}
    t.Log(results)
}