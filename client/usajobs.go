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
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	usajobsBaseApiUrl = "https://data.usajobs.gov/api"
	usajobsHost       = "data.usajobs.gov"
)

// Client contains the HTTP client for making apis calls and services
// for interacting with the different usajobs api endpoints
type Client struct {

	// HTTP client used to communicate with the API.
	Client *http.Client

	// base url to uasjabs api (include /api/)
	BaseURL *url.URL

	// host sets the host header value for the client when communicating
	// with the usajobs api. This is a required field.
	// ref: https://developer.usajobs.gov/tutorials/search-jobs
	Host string

	// User agent used when communicating with the usajobs api. This is a
	// required field.
	UserAgent string

	// Api token for authenticating with the usajobs api. This is a required
	// field and can be requested from usajobs: https://developer.usajobs.gov/general/quick-start
	ApiToken string

	// services used for communicating with different aspects of the
	// usajobs api.
	Search                        *SearchService
	Agency                        *AgencySubelementsService
	AcademicHonors                *AcademicHonorsService
	AcademicLevels                *AcademicLevelsService
	ApplicantSuppliers            *ApplicantSuppliersService
	ApplicationStatuses           *ApplicationStatusesService
	Countries                     *CountriesService
	CountrySubdivisions           *CountrySubdivisionsService
	CyberWorkGroupings            *CyberWorkGroupingsService
	CyberWorkRoles                *CyberWorkRolesService
	DegreeTypeCode                *DegreeTypeCodeService
	Disabilities                  *DisabilitiesService
	Documentations                *DocumentationsService
	DocumentFormats               *DocumentFormatsService
	Ethnicities                   *EthnicitiesService
	FederalEmploymentStatuses     *FederalEmploymentStatusesService
	GeoLocCodes                   *GeoLocCodesService
	GsaGeoLocCodes                *GsaGeoLocCodesService
	HiringPaths                   *HiringPathsService
	KeyStandardRequirements       *KeyStandardRequirementsService
	LanguageCodes                 *LanguageCodesService
	LanguageProficiencies         *LanguageProficienciesService
	LocationExpansions            *LocationExpansionsService
	MilitaryStatusCodes           *MilitaryStatusCodesService
	MissionCriticalCodes          *MissionCriticalCodesService
	OccupationalSeries            *OccupationalSeriesService
	PayPlans                      *PayPlansService
	PositionOfferingTypes         *PositionOfferingTypesService
	PositionOpeningStatuses       *PositionOpeningsStatusesService
	PositionScheduleTypes         *PositionScheduleTypesService
	PostalCodes                   *PostalCodesService
	RaceCodes                     *RaceCodesService
	RefereeTypeCodes              *RefereeTypeCodesService
	RemunerationRateIntervalCodes *RemunerationRateIntervalCodesService
	RequiredStandardDocuments     *RequiredStandardDocumentsService
	SecurityClearances            *SecurityClearancesService
	ServiceTypes                  *ServiceTypesService
	SpecialHirings                *SpecialHiringsService
	TravelPercentages             *TravelPercentagesService
	WhoMayApply                   *WhoMayApplyService
}

// NewClient requires a user agent and api token string variables and returns
// a Client object or error. The user agent string is the email address provided
// to usajobs when requesting an api token.
func NewClient(userAgent, apiToken string) (*Client, error) {
	if userAgent == "" || apiToken == "" {
		return nil, errors.New("user agent and api token values required")
	}

	u, err := url.Parse(usajobsBaseApiUrl)
	if err != nil {
		return nil, errors.New("failed to parse string url to Url")
	}

	h := http.Client{}

	c := Client{
		Client:    &h,
		BaseURL:   u,
		UserAgent: userAgent,
		ApiToken:  apiToken,
		Host:      usajobsHost,
	}

	c.Search = NewSearchService(&c)
	c.Agency = NewAgencySubelementsService(&c)
	c.AcademicHonors = NewAcademicHonorsService(&c)
	c.AcademicLevels = NewAcademicLevelsService(&c)
	c.ApplicantSuppliers = NewApplicantSuppliersService(&c)
	c.ApplicationStatuses = NewApplicationStatusesService(&c)
	c.Countries = NewCountriesService(&c)
	c.CountrySubdivisions = NewCountrySubdivisionsService(&c)
	c.CyberWorkGroupings = NewCyberWorkGroupingsService(&c)
	c.CyberWorkRoles = NewCyberWorkRolesService(&c)
	c.DegreeTypeCode = NewDegreeTypeCodeService(&c)
	c.Disabilities = NewDisabilitiesService(&c)
	c.Documentations = NewDocumentationsService(&c)
	c.DocumentFormats = NewDocumentFormatsService(&c)
	c.Ethnicities = NewEthnicitiesService(&c)
	c.FederalEmploymentStatuses = NewFederalEmploymentStatusesService(&c)
	c.GeoLocCodes = NewGeoLocCodesService(&c)
	c.GsaGeoLocCodes = NewGsaGeoLocCodesService(&c)
	c.HiringPaths = NewHiringPathsService(&c)
	c.KeyStandardRequirements = NewKeyStandardRequirementsService(&c)
	c.LanguageCodes = NewLanguageCodesService(&c)
	c.LanguageProficiencies = NewLanguageProficienciesService(&c)
	c.LocationExpansions = NewLocationExpansionsService(&c)
	c.MilitaryStatusCodes = NewMilitaryStatusCodesService(&c)
	c.MissionCriticalCodes = NewMissionCriticalCodesService(&c)
	c.OccupationalSeries = NewOccupationalSeriesService(&c)
	c.PayPlans = NewPayPlansService(&c)
	c.PositionOfferingTypes = NewPositionOfferingTypesService(&c)
	c.PositionOpeningStatuses = NewPositionOpeningsStatusesService(&c)
	c.PositionScheduleTypes = NewPositionScheduleTypesService(&c)
	c.PostalCodes = NewPostalCodesService(&c)
	c.RaceCodes = NewRaceCodesService(&c)
	c.RefereeTypeCodes = NewRefereeTypeCodesService(&c)
	c.RemunerationRateIntervalCodes = NewRemunerationRateIntervalCodesService(&c)
	c.RequiredStandardDocuments = NewRequiredStandardDocumentsService(&c)
	c.SecurityClearances = NewSecurityClearancesService(&c)
	c.ServiceTypes = NewServiceTypesService(&c)
	c.SpecialHirings = NewSpecialHiringsService(&c)
	c.TravelPercentages = NewTravelPercentagesService(&c)
	c.WhoMayApply = NewWhoMayApplyService(&c)

	return &c, nil
}

// NewRequest accepts a method string for the http method (ex., GET) and the
// url string. The url string is just the endpoint and query parameters without
// the baseurl. New Request also sets the required host headers for interacting
// with the usajobs api. Returns the created request or an error.
func (c *Client) NewRequest(method, urlStr string) (*http.Request, error) {

	combined := c.BaseURL.String() + urlStr

	req, err := http.NewRequest(method, combined, nil)
	if err != nil {
		return nil, err
	}

	// set the required host headers for communicating with the usajobs api
	req.Header.Set("Host", c.Host)
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization-Key", c.ApiToken)
	return req, nil
}

func (c *Client) NewResponse(endpoint string, opt interface{}, resp interface{}) (*http.Response, interface{}, error) {

	requestURL := endpoint
	if opt != nil {
		qs, err := query.Values(opt)
		if err != nil {
			return nil, resp, err
		}
		requestURL = fmt.Sprintf("%s?%s", endpoint, qs.Encode())
	}

	req, err := c.NewRequest("GET", requestURL)
	if err != nil {
		return nil, resp, err
	}

	response, err := c.Client.Do(req)
	if err != nil {
		return nil, resp, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return response, resp, err
	}

	return response, resp, nil
}
