
<p align="center">
    <a href="https://github.com/JeffRDay/go-usajobs/releases" alt="Release">
        <img src="https://img.shields.io/github/v/release/jeffrday/go-usajobs" /></a>
</p>

# go-usajobs

An unoffocial golang http client library and CLI for the usajobs api.

## Use Case

Software Engineers, Platform Engineers, and other folks in the IT community live
closer to the command-line interface and/or within developer environments. The 
goal for this project is to help bring US Government jobs announced on USAJobs
closer to this group of people while providing a minor contribution to the open 
source community to enable others to further expand the reach of these job
announcements.

## Getting Started

You can use the CLI to start querying the API immediately or integrate
the Client library into your own Golang applications. In both cases, you must
first obtain a developer API token from USAJobs (see below).

## Obtaining an API Token

Go to [USAJobs](https://developer.usajobs.gov/apirequest/) to register for
a developer API token. The process is fully automated and you should receive
an email from them with a token quickly.

### Obtaining the latest CLI binary

Please download the latest binary for your operating system from this repository's
release page!

### Adding the Client to your Go project

```bash
go get github.com/JeffRDay/go-usajobs/client
```

### USAJobs CLI Example

```bash
export TOKEN=<TOKEN>
export USER=<EMAIL>

./usajobs search --keyword=army --token=$TOKEN --user-agent=$USER --min-salary=80,000
```

### USAJobs API Client Example

```go
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

	// Find awesome work with the Army Software Factory
	opt := usajobs.SearchOptions{
		JobCategoryCode: []string{"2210", "0854"},
	}

	_, r, err := c.Search.WithOptions(&opt)
	if err != nil {
		panic(err.Error())
	}

	prettyJSON, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(prettyJSON))
}
```

## Support

- [X] /search
- [ ] /historicjoa (usajobs needs to fix)
- [X] /codelist/academichonors
- [X] /codelist/academiclevels
- [X] /codelist/agencysubelements
- [ ] /codelist/announcementclosingtype (404, usajobs needs to remove from docs)
- [-] /codelist/applicantsuppliers
- [-] /codelist/applicationstatuses
- [-] /codelist/countries
- [-] /codelist/countrysubdivisions
- [-] /codelist/cyberworkgroupings
- [-] /codelist/cyberworkroles
- [-] /codelist/degreetypecode
- [-] /codelist/disabilities
- [-] /codelist/documentations
- [-] /codelist/documentformats
- [-] /codelist/ethnicity
- [-] /codelist/federalemploymentstatuses
- [-] /codelist/geoloccodes
- [-] /codelist/gsageoloccodes
- [-] /codelist/hiringpaths
- [-] /codelist/keystandardrequirements
- [-] /codelist/languagecodes
- [-] /codelist/languageproficiency
- [-] /codelist/locationexpansions
- [-] /codelist/militarystatuscodes
- [-] /codelist/missioncriticalcodes
- [-] /codelist/occupationalseries
- [-] /codelist/payplans
- [-] /codelist/positionofferingtypes
- [-] /codelist/positionopeningstatuses
- [-] /codelist/positionscheduletypes
- [-] /codelist/postalcodes
- [-] /codelist/racecodes
- [-] /codelist/refereetypecodes
- [-] /codelist/remunerationrateintervalcodes
- [-] /codelist/requiredstandarddocuments
- [-] /codelist/securityclearances
- [-] /codelist/servicetypes
- [-] /codelist/specialhirings
- [-] /codelist/travelpercentages
- [-] /codelist/whomayapply

## Other Tools

- [fedscope](https://www.fedscope.opm.gov/ibmcognos/bi/?perspective=home)

## Contributing

Contributions are welome! Please file a PR for additions or changes. 
