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

// ApplicantSuppliersService is used for interacting with the /codelist/applicantsuppliers
// endpoint of the usajobs api.
type ApplicantSuppliersService struct {
	Client *Client
}

// NewApplicantSuppliersService instatiates and returns a search service for this client.
func NewApplicantSuppliersService(c *Client) *ApplicantSuppliersService {
	as := new(ApplicantSuppliersService)
	as.Client = c
	return as
}

// ApplicantSuppliersOptions are the url query parameters supported by the
// /codelist/applicantsuppliers usajobs api endpoint.
type ApplicantSuppliersOptions struct {
	LastModified string `url:"lastmodified,omitempty"`
}

// ApplicantSuppliersResponse is the golang struct implementation of all possible response
// fields from /codelist/applicantsuppliers. Consumers are responsible for ensuring omitted fields
// do not cause errors in consumer implementations.
type ApplicantSuppliersResponse struct {
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

// WithOptions executes a request to the usajobs /codelist/academichonors endpoint
// with the provided options. Pass nil if no options desired.
func (as *ApplicantSuppliersService) WithOptions(opt *ApplicantSuppliersOptions) (*http.Response, *ApplicantSuppliersResponse, error) {
	usajobsEndpoint := "/codelist/applicantsuppliers"
	responseObject := new(ApplicantSuppliersResponse)
	r, object, err := as.Client.NewResponse(usajobsEndpoint, opt, responseObject)
	return r, object.(*ApplicantSuppliersResponse), err
}
