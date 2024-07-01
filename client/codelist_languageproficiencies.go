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
package usajobs

import (
	"net/http"
)

// LanguageProficienciesService is used for interacting with the /codelist/languageproficiencies
// endpoint of the usajobs api.
type LanguageProficienciesService struct {
	Client *Client
}

// NewLanguageProficienciesService instatiates and returns a search service for this client.
func NewLanguageProficienciesService(c *Client) *LanguageProficienciesService {
	as := new(LanguageProficienciesService)
	as.Client = c
	return as
}

// LanguageProficienciesOptions are the url query parameters supported by the
// /codelist/languageproficiencies usajobs api endpoint.
type LanguageProficienciesOptions struct {
	LastModified string `url:"lastmodified,omitempty"`
}

// LanguageProficienciesResponse is the golang struct implementation of all possible response
// fields from /codelist/languageproficiencies. Consumers are responsible for ensuring omitted fields
// do not cause errors in consumer implementations.
type LanguageProficienciesResponse struct {
	CodeList []struct {
		ValidValue []struct {
			Code         string `json:"Code,omitempty"`
			Value        string `json:"Value,omitempty"`
			LastModified string `json:"LastModified,omitempty"`
			IsDisabled   string `json:"IsDisabled,omitempty"`
		} `json:"ValidValue,omitempty"`
		ID string `json:"id,omitempty"`
	} `json:"CodeList,omitempty"`
	DateGenerated string `json:"DateGenerated,omitempty"`
}

// WithOptions executes a request to the usajobs /codelist/languageproficiencies endpoint
// with the provided options. Pass nil if no options desired.
func (as *LanguageProficienciesService) WithOptions(opt *LanguageProficienciesOptions) (*http.Response, *LanguageProficienciesResponse, error) {
	usajobsEndpoint := "/codelist/languageproficiencies"
	responseObject := new(LanguageProficienciesResponse)
	r, object, err := as.Client.NewResponse(usajobsEndpoint, opt, responseObject)
	return r, object.(*LanguageProficienciesResponse), err
}
