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

// FederalEmploymentStatusesService is used for interacting with the /codelist/federalemploymentstatuses
// endpoint of the usajobs api.
type FederalEmploymentStatusesService struct {
	Client *Client
}

// NewFederalEmploymentStatusesService instatiates and returns a search service for this client.
func NewFederalEmploymentStatusesService(c *Client) *FederalEmploymentStatusesService {
	as := new(FederalEmploymentStatusesService)
	as.Client = c
	return as
}

// FederalEmploymentStatusesOptions are the url query parameters supported by the
// /codelist/federalemploymentstatuses  usajobs api endpoint.
type FederalEmploymentStatusesOptions struct {
	LastModified string `url:"lastmodified,omitempty"`
}

// FederalEmploymentStatusesResponse is the golang struct implementation of all possible response
// fields from /codelist/federalemploymentstatuses. Consumers are responsible for ensuring omitted fields
// do not cause errors in consumer implementations.
type FederalEmploymentStatusesResponse struct {
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

// WithOptions executes a request to the usajobs /codelist/federalemploymentstatuses endpoint
// with the provided options. Pass nil if no optons desired.
func (as *FederalEmploymentStatusesService) WithOptions(opt *FederalEmploymentStatusesOptions) (*http.Response, *FederalEmploymentStatusesResponse, error) {
	usajobsEndpoint := "/codelist/federalemploymentstatuses"
	responseObject := new(FederalEmploymentStatusesResponse)
	r, object, err := as.Client.NewResponse(usajobsEndpoint, opt, responseObject)
	return r, object.(*FederalEmploymentStatusesResponse), err
}
