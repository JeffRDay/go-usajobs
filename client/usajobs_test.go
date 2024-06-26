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
	"testing"

	usajobs "github.com/JeffRDay/go-usajobs/client"
)

func TestNewClient(t *testing.T) {
	t.Run("test missing user agent", func(t *testing.T) {
		_, err := usajobs.NewClient("", "testtest")
		if err == nil {
			t.Error("expected error, go nil")
			t.FailNow()
		}
	})

	t.Run("test missing apitoken", func(t *testing.T) {
		_, err := usajobs.NewClient("testtest", "")
		if err == nil {
			t.Error("expected error, go nil")
			t.FailNow()
		}
	})

	t.Run("test bad baseurl", func(t *testing.T) {
		_, err := usajobs.NewClient("testtest", "")
		if err == nil {
			t.Error("expected error, go nil")
			t.FailNow()
		}
	})
}

func TestNewRequest(t *testing.T) {
	host := "data.usajobs.gov"
	userAgent := "test"
	apiToken := "test"
	method := "GET"
	urlStr := "/search/"
	expectedUrl := "https://data.usajobs.gov/api/search/"

	c, err := usajobs.NewClient(userAgent, apiToken)
	if err != nil {
		t.Errorf("%s", err.Error())
		t.FailNow()
	}

	r, err := c.NewRequest(method, urlStr)
	if err != nil {
		t.Errorf("%s", err.Error())
		t.FailNow()
	}

	if r.URL.String() != expectedUrl {
		t.Errorf("expected %s, got %s", expectedUrl, r.URL.String())
	}

	// host, user agent, and auth key are required headers
	// by usajobs for interacting with the api.
	if r.Header.Get("Host") != host {
		t.Errorf("Expected Host header to be %s, got %s", host, r.Header.Get("Host"))
	}

	if r.Header.Get("User-Agent") != userAgent {
		t.Errorf("Expected User-Agent header to be %s, got %s", userAgent, r.Header.Get("User-Agent"))
	}

	if r.Header.Get("Authorization-Key") != apiToken {
		t.Errorf("Expected Authorization-Key header to be %s, got %s", apiToken, r.Header.Get("Authorization-Key"))
	}
}
