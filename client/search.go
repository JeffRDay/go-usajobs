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
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// SearchService is used for interacting with the /search endpoint of the
// usajobs api.
type SearchService struct {
	Client *Client
}

// SearchOptions are most of the url parameters the usajobs api /search api
// supports. SalaryBucket, GradeBucket, and MissionCriticalTags are query
// parameters that are not currently supported, mostly because they are
// duplicative of other options already supported. Submit an issue or PR if
// this is breaking something you would like to do.
type SearchOptions struct {
	Keyword                   string   `url:"Keyword,omitempty"`
	PositionTitle             string   `url:"PositionTitle,omitempty"`
	RemunerationMinimumAmount string   `url:"RemunerationMinimumAmount,omitempty"`
	RemunerationMaximumAmount string   `url:"RemunerationMaximumAmount,omitempty"`
	PayGradeHigh              string   `url:"PayGradeHigh,omitempty"`
	PayGradeLow               string   `url:"PayGradeLow,omitempty"`
	JobCategoryCode           []string `url:"JobCategoryCode,omitempty" del:";"`
	LocationName              []string `url:"LocationName,omitempty" del:";"`
	PostingChannel            []string `url:"PostingChannel,omitempty" del:";"`
	Organization              []string `url:"Organization,omitempty" del:";"`
	PositionOfferingTypeCode  []string `url:"PositionOfferingTypeCode,omitempty" del:";"`
	TravelPercentage          int      `url:"TravelPercentage,omitempty"`
	PositionScheduleTypeCode  []int    `url:"PositionSchedule,omitempty" del:";"`
	RelocationIndicator       bool     `url:"RelocationIndicator,omitempty"`
	SecurityClearanceRequired []int    `url:"SecurityClearanceRequired,omitempty" del:";"`
	SupervisoryStatus         string   `url:"SupervisoryStatus,omitempty"`
	DatePosted                int      `url:"DatePosted,omitempty"`
	JobGradeCode              []int    `url:"JobGradeCode,omitempty" del:";"`
	SortField                 string   `url:"SortField,omitempty"`
	SortDirection             string   `url:"SortDirection,omitempty"`
	Page                      int      `url:"Page,omitempty"`
	ResultsPerPage            int      `url:"ResultsPerPage,omitempty"`
	WhoMayApply               string   `url:"WhoMayApply,omitempty"`
	Radius                    int      `url:"Radius,omitempty"`
	Fields                    string   `url:"Fields,omitempty"`
	HiringPath                []string `url:"HiringPath,omitempty" del:";"`
	PositionSensitivity       []int    `url:"PositionSensitivity,omitempty" del:";"`
	RemoteIndicator           bool     `url:"RemoteIndicator,omitempty"`
}

// SearchResponse is the golang struct implementation of all possible response
// fields from /search. Consumers are responsible for ensuring omitted fields
// do not cause errors in consumer implementations.
type SearchResponse struct {
	LanguageCode     string `json:"LanguageCode,omitempty"`
	SearchParameters struct {
	} `json:"SearchParameters,omitempty"`
	SearchResult struct {
		SearchResultCount    int `json:"SearchResultCount,omitempty"`
		SearchResultCountAll int `json:"SearchResultCountAll,omitempty"`
		SearchResultItems    []struct {
			MatchedObjectID         string `json:"MatchedObjectId,omitempty"`
			MatchedObjectDescriptor struct {
				PositionID              string   `json:"PositionID,omitempty"`
				PositionTitle           string   `json:"PositionTitle,omitempty"`
				PositionURI             string   `json:"PositionURI,omitempty"`
				ApplyURI                []string `json:"ApplyURI,omitempty"`
				PositionLocationDisplay string   `json:"PositionLocationDisplay,omitempty"`
				PositionLocation        []struct {
					LocationName           string  `json:"LocationName,omitempty"`
					CountryCode            string  `json:"CountryCode,omitempty"`
					CountrySubDivisionCode string  `json:"CountrySubDivisionCode,omitempty"`
					CityName               string  `json:"CityName,omitempty"`
					Longitude              float64 `json:"Longitude,omitempty"`
					Latitude               float64 `json:"Latitude,omitempty"`
				} `json:"PositionLocation,omitempty"`
				OrganizationName string `json:"OrganizationName,omitempty"`
				DepartmentName   string `json:"DepartmentName,omitempty"`
				JobCategory      []struct {
					Name string `json:"Name,omitempty"`
					Code string `json:"Code,omitempty"`
				} `json:"JobCategory,omitempty"`
				JobGrade []struct {
					Code string `json:"Code,omitempty"`
				} `json:"JobGrade,omitempty"`
				PositionSchedule []struct {
					Name string `json:"Name,omitempty"`
					Code string `json:"Code,omitempty"`
				} `json:"PositionSchedule,omitempty"`
				PositionOfferingType []struct {
					Name string `json:"Name,omitempty"`
					Code string `json:"Code,omitempty"`
				} `json:"PositionOfferingType,omitempty"`
				QualificationSummary string `json:"QualificationSummary,omitempty"`
				PositionRemuneration []struct {
					MinimumRange     string `json:"MinimumRange,omitempty"`
					MaximumRange     string `json:"MaximumRange,omitempty"`
					RateIntervalCode string `json:"RateIntervalCode,omitempty"`
					Description      string `json:"Description,omitempty"`
				} `json:"PositionRemuneration,omitempty"`
				PositionStartDate            string `json:"PositionStartDate,omitempty"`
				PositionEndDate              string `json:"PositionEndDate,omitempty"`
				PublicationStartDate         string `json:"PublicationStartDate,omitempty"`
				ApplicationCloseDate         string `json:"ApplicationCloseDate,omitempty"`
				PositionFormattedDescription []struct {
					Content          string `json:"Content,omitempty"`
					Label            string `json:"Label,omitempty"`
					LabelDescription string `json:"LabelDescription,omitempty"`
				} `json:"PositionFormattedDescription,omitempty"`
				UserArea struct {
					Details struct {
						MajorDuties       []string `json:"MajorDuties,omitempty"`
						Education         string   `json:"Education,omitempty"`
						Requirements      string   `json:"Requirements,omitempty"`
						Evaluations       string   `json:"Evaluations,omitempty"`
						HowToApply        string   `json:"HowToApply,omitempty"`
						WhatToExpectNext  string   `json:"WhatToExpectNext,omitempty"`
						RequiredDocuments string   `json:"RequiredDocuments,omitempty"`
						Benefits          string   `json:"Benefits,omitempty"`
						BenefitsURL       string   `json:"BenefitsUrl,omitempty"`
						OtherInformation  string   `json:"OtherInformation,omitempty"`
						KeyRequirements   []any    `json:"KeyRequirements,omitempty"`
						JobSummary        string   `json:"JobSummary,omitempty"`
						WhoMayApply       struct {
							Name string `json:"Name,omitempty"`
							Code string `json:"Code,omitempty"`
						} `json:"WhoMayApply,omitempty"`
						LowGrade          string `json:"LowGrade,omitempty"`
						HighGrade         string `json:"HighGrade,omitempty"`
						SubAgencyName     string `json:"SubAgencyName,omitempty"`
						OrganizationCodes string `json:"OrganizationCodes,omitempty"`
					} `json:"Details,omitempty"`
					IsRadialSearch bool `json:"IsRadialSearch,omitempty"`
				} `json:"UserArea,omitempty"`
			} `json:"MatchedObjectDescriptor,omitempty"`
			RelevanceRank float64 `json:"RelevanceRank,omitempty"`
		} `json:"SearchResultItems,omitempty"`
		UserArea struct {
			Refiners struct {
				Organization []struct {
					RefinementName  string `json:"RefinementName,omitempty"`
					RefinementCount string `json:"RefinementCount,omitempty"`
					RefinementToken string `json:"RefinementToken,omitempty"`
					RefinementValue string `json:"RefinementValue,omitempty"`
				} `json:"Organization,omitempty"`
				GradeBucket []struct {
					RefinementName  string `json:"RefinementName,omitempty"`
					RefinementCount string `json:"RefinementCount,omitempty"`
					RefinementToken string `json:"RefinementToken,omitempty"`
					RefinementValue string `json:"RefinementValue,omitempty"`
				} `json:"GradeBucket,omitempty"`
				SalaryBucket []struct {
					RefinementName  string `json:"RefinementName,omitempty"`
					RefinementCount string `json:"RefinementCount,omitempty"`
					RefinementToken string `json:"RefinementToken,omitempty"`
					RefinementValue string `json:"RefinementValue,omitempty"`
				} `json:"SalaryBucket,omitempty"`
				PositionOfferingTypeCode []struct {
					RefinementName  string `json:"RefinementName,omitempty"`
					RefinementCount string `json:"RefinementCount,omitempty"`
					RefinementToken string `json:"RefinementToken,omitempty"`
					RefinementValue string `json:"RefinementValue,omitempty"`
				} `json:"PositionOfferingTypeCode,omitempty"`
				PositionScheduleTypeCode []struct {
					RefinementName  string `json:"RefinementName,omitempty"`
					RefinementCount string `json:"RefinementCount,omitempty"`
					RefinementToken string `json:"RefinementToken,omitempty"`
					RefinementValue string `json:"RefinementValue,omitempty"`
				} `json:"PositionScheduleTypeCode,omitempty"`
				JobCategoryCode []struct {
					RefinementName  string `json:"RefinementName,omitempty"`
					RefinementCount string `json:"RefinementCount,omitempty"`
					RefinementToken string `json:"RefinementToken,omitempty"`
					RefinementValue string `json:"RefinementValue,omitempty"`
				} `json:"JobCategoryCode,omitempty"`
			} `json:"Refiners,omitempty"`
			NumberOfPages  string `json:"NumberOfPages,omitempty"`
			IsRadialSearch bool   `json:"IsRadialSearch,omitempty"`
		} `json:"UserArea,omitempty"`
	} `json:"SearchResult,omitempty"`
}

// NewSearchService instatiates and returns a search service for this client.
func NewSearchService(c *Client) *SearchService {
	ss := new(SearchService)
	ss.Client = c
	return ss
}

// WithOptions executes a request to the usajobs /search endpoint with the
// provided search options. Pass nil if no search options desired.
func (s *SearchService) WithOptions(opt *SearchOptions) (*http.Response, SearchResponse, error) {

	usajobsEndpoint := "/search"
	sr := SearchResponse{}

	requestURL := usajobsEndpoint
	if opt != nil {
		qs, err := query.Values(opt)
		if err != nil {
			return nil, sr, err
		}
		requestURL = fmt.Sprintf("%s?%s", usajobsEndpoint, qs.Encode())
	}

	req, err := s.Client.NewRequest("GET", requestURL)
	if err != nil {
		return nil, sr, err
	}

	response, err := s.Client.Client.Do(req)
	if err != nil {
		return response, sr, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&sr)
	if err != nil {
		return response, sr, err
	}

	return response, sr, nil
}
