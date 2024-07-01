/*
Copyright © 2024 Jeff Day jeffrey.day33@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package usajobs_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	usajobs "github.com/JeffRDay/go-usajobs/client"
)

func TestPositionOpeningStatuses(t *testing.T) {
	testdata := "../testdata/positionopeningstatuses-testdata.json"
	// Read the JSON file from testdata directory
	file, err := os.Open(testdata)
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

		if strings.Contains(r.URL.String(), "/codelist/positionopeningstatuses") {
			// Return status OK for the specific URL
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		} else {
			// Return status Not Found for any other URL
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer mockServer.Close()

	c, err := usajobs.NewClient("test", "test")
	if err != nil {
		t.Fatalf("could not create new usajobs client: %v", err)
	}

	u, err := url.Parse(mockServer.URL)
	if err != nil {
		t.Fatalf("failed to parse mock server url: %v", err)
	}

	c.BaseURL = u

	_, _, err = c.PositionOpeningStatuses.WithOptions(nil)
	if err != nil {
		t.Fatalf("failed to execute search request: %v", err.Error())
	}

}
