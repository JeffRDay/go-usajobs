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
package cmd

import (
	"fmt"
	"strings"

	usajobs "github.com/JeffRDay/go-usajobs/client"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for jobs announced on usajobs",
	Long: `
Search for jobs announced on USAJobs. Sub-commands of the search command support all
sorting and filtering functions that are available to job seekers using the USAJobs web
application. 

Example Usage:

    Search using a keyword and multiple locations:
    search --token=$TOKEN --user-agent=$EMAIL --keyword=army --location=Austin,Texas-Portland,Oregon

    `,
	Run: func(cmd *cobra.Command, args []string) {
		opt := setSearchOptions()
		r := executeSearch(&opt)

		for _, job := range r {
			fmt.Println(job)
		}
	},
}

var (
	Client                    *usajobs.Client
	Keyword                   string
	PositionTitle             string
	RemunerationMinimumAmount string
	RemunerationMaximumAmount string
	PayGradeHigh              string
	PayGradeLow               string
	JobCategoryCode           []string
	LocationName              string
	PostingChannel            []string
	Organization              []string
	PositionOfferingTypeCode  []string
	TravelPercentage          int
	PositionScheduleTypeCode  []int
	RelocationIndicator       bool
	SecurityClearanceRequired []int
	SupervisoryStatus         string
	DatePosted                int
	JobGradeCode              []int
	SortField                 string
	SortDirection             string
	Page                      int
	ResultsPerPage            int
	WhoMayApply               string
	Radius                    int
	Fields                    string
	HiringPath                []string
	PositionSensitivity       []int
	RemoteIndicator           bool
)

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().StringVarP(&Keyword, "keyword", "k", "", "[optional] Words used to refine search (ex., Army Software Factory)")
	searchCmd.PersistentFlags().StringVar(&PositionTitle, "title", "", "[optional] filter jobs by position title (ex., IT Specialist)")
	searchCmd.PersistentFlags().StringVar(&RemunerationMinimumAmount, "min-salary", "", "[optional] Sets the lower limit for filtering jobs by salary (ex., 80,000)")
	searchCmd.PersistentFlags().StringVar(&RemunerationMaximumAmount, "max-salary", "", "[optional] Sets the upper limit for filtering jobs by salary (ex., 120,000)")
	searchCmd.PersistentFlags().StringVar(&PayGradeLow, "min-grade", "", "[optional] Sets the lower limit for filtering jobs by pay grade (ex, GS14)")
	searchCmd.PersistentFlags().StringVar(&PayGradeHigh, "max-grade", "", "[optional] Sets the upper limit for filtering jobs by pay grade (ex., GS15)")
	searchCmd.PersistentFlags().StringSliceVarP(&JobCategoryCode, "job-catagory", "j", []string{""}, "[optional] Comma separated list of job codes (ex., 2210, 0854)")
	searchCmd.PersistentFlags().StringVar(&LocationName, "location", "", "[optional] dash (-) separated list of <city,state> (ex., Austin,Texas-Portland,Oregon)")
	searchCmd.PersistentFlags().StringSliceVar(&Organization, "organization", []string{""}, "[optional] Comma separated list of organizations (ex., Immigration and Customs Enforcement,Office of Chief Information Officer)")
	searchCmd.PersistentFlags().StringSliceVar(&PositionOfferingTypeCode, "position-type", []string{}, "[optional] Filter jobs by position type (ex., 15317)")
	searchCmd.PersistentFlags().IntVar(&TravelPercentage, "travel-rate", -1, "[optional] Filter jobs by percent of travel (ex., 25)")
	searchCmd.PersistentFlags().IntSliceVar(&PositionScheduleTypeCode, "position-schedule-type-code", []int{}, "[optional][Comma Separated List] Filter jobs by schedule position type code (ex., 6,2)")
	searchCmd.PersistentFlags().BoolVar(&RelocationIndicator, "relocation", false, "[optional][true/false] Only show jobs that offer relocation assistance if true.")
	searchCmd.PersistentFlags().IntSliceVar(&SecurityClearanceRequired, "clearance", []int{}, "[optional][Comma Separated List] Filter jobs by clearance types (ex., 1,2,3)")
	searchCmd.PersistentFlags().StringVar(&SupervisoryStatus, "supervisory-status", "", "[optional] TBD, don't use yet")
	searchCmd.PersistentFlags().IntVar(&DatePosted, "date-posted", -1, "[optional][0 to 60] Filter jobs that were posted within the number of days specified")
	searchCmd.PersistentFlags().IntSliceVar(&JobGradeCode, "job-grade-code", []int{}, "[optional] Filter for jobs containing the specified Job Grade Codes")
	searchCmd.PersistentFlags().StringVar(&SortField, "sort-by", "", "[optional] Sort results by the specified value.")
	searchCmd.PersistentFlags().StringVar(&SortDirection, "sort-direction", "", "[optional][Asc/Dsc] Ascending or Descending sort order")
	searchCmd.PersistentFlags().IntVar(&ResultsPerPage, "num-results", 500, "[optional][25-500] number of results to return, 0 returns all")
	searchCmd.PersistentFlags().StringVar(&WhoMayApply, "who-may-apply", "", "[optional][All|Public|Status] Filter jobs based on who can apply")
	searchCmd.PersistentFlags().IntVar(&Radius, "radius", -1, "[optional][int] Radius of miles from location to filter jobs")
	searchCmd.PersistentFlags().StringSliceVar(&HiringPath, "hiring-path", []string{}, "[optional][Comma Seperated List]")
	searchCmd.PersistentFlags().IntSliceVar(&PositionSensitivity, "position-sensitivity", []int{}, "[optional][Comma Separated List] Sensitivity Codes to filter jobs by position sensitivity")
	searchCmd.PersistentFlags().BoolVar(&RemoteIndicator, "remote", false, "[optional][true/false] Only shows jobs supporting remote work if true")
}

func setSearchOptions() usajobs.SearchOptions {
	var opt usajobs.SearchOptions

	if Keyword != "" {
		opt.Keyword = Keyword
	}

	if PositionTitle != "" {
		opt.PositionTitle = PositionTitle
	}

	if RemunerationMinimumAmount != "" {
		opt.RemunerationMinimumAmount = RemunerationMinimumAmount
	}

	if RemunerationMaximumAmount != "" {
		opt.RemunerationMaximumAmount = RemunerationMaximumAmount
	}

	if PayGradeLow != "" {
		opt.PayGradeLow = PayGradeLow
	}

	if PayGradeHigh != "" {
		opt.PayGradeHigh = PayGradeHigh
	}

	if len(JobCategoryCode) >= 1 {
		opt.JobCategoryCode = JobCategoryCode
	}

	if LocationName != "" {
		locationSlice := strings.Split(LocationName, "-")
		for _, i := range locationSlice {
			fmt.Println(i)
		}
		opt.LocationName = locationSlice
	}

	if len(Organization) >= 1 {
		opt.Organization = Organization
	}

	if len(PositionOfferingTypeCode) >= 1 {
		opt.PositionOfferingTypeCode = PositionOfferingTypeCode
	}

	if TravelPercentage > -1 {
		opt.TravelPercentage = TravelPercentage
	}

	if len(PositionScheduleTypeCode) >= 1 {
		opt.PositionScheduleTypeCode = PositionScheduleTypeCode
	}

	if RelocationIndicator {
		opt.RelocationIndicator = RelocationIndicator
	}

	if len(SecurityClearanceRequired) >= 1 {
		opt.SecurityClearanceRequired = SecurityClearanceRequired
	}

	if SupervisoryStatus != "" {
		opt.SupervisoryStatus = SupervisoryStatus
	}

	if DatePosted > -1 {
		opt.DatePosted = DatePosted
	}

	if len(JobGradeCode) > 0 {
		opt.JobGradeCode = JobGradeCode
	}

	if SortField != "" {
		opt.SortField = SortField
	}

	if SortDirection != "" {
		opt.SortDirection = SortDirection
	}

	if ResultsPerPage == 0 || ResultsPerPage < 0 || ResultsPerPage > 500 {
		opt.ResultsPerPage = 500
	}

	if ResultsPerPage > 0 && ResultsPerPage < 501 {
		opt.ResultsPerPage = ResultsPerPage
	}

	if WhoMayApply != "" {
		opt.WhoMayApply = WhoMayApply
	}

	if Radius > 0 {
		opt.Radius = Radius
	}

	if len(HiringPath) > 0 {
		opt.HiringPath = HiringPath
	}

	if len(PositionSensitivity) > 0 {
		opt.PositionSensitivity = PositionSensitivity
	}

	if RemoteIndicator {
		opt.RemoteIndicator = RemoteIndicator
	}

	return opt
}

func executeSearch(opt *usajobs.SearchOptions) []string {

	var err error
    if Client == nil {
		Client, err = usajobs.NewClient(userAgent, apiToken)
		if err != nil {
			panic(err.Error())
		}
	}

	r, err := Client.Search.WithOptions(opt)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf(
		"Displaying %v of %v job announcements matching this search\n",
		r.SearchResult.SearchResultCount,
		r.SearchResult.SearchResultCountAll,
	)

	var results []string
	result := fmt.Sprintf(
		"%s\t%s\t%s",
		"PositionID",
		"Organization",
		"PositionTitle",
	)
	results = append(results, result)

	for _, i := range r.SearchResult.SearchResultItems {
		result := fmt.Sprintf(
			"%s\t%s\t%s",
			i.MatchedObjectID,
			i.MatchedObjectDescriptor.DepartmentName,
			i.MatchedObjectDescriptor.PositionTitle,
		)

		results = append(results, result)
	}

	return results
}
