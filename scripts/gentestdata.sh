#!/bin/bash

# Define an array of URLs
CODELIST_ENDPOINTS=(
    "academichonors"
    "academiclevels"
    "actioncodes"
    "agencysubelemnts"
    "applicantsuppliers"
    "applicationstatuses"
    "countries"
    "countrysubdivisions"
    "cyberworkgroupings"
    "cyberworkroles"
    "degreetypecodes"
    "disabilities"
    "documentations"
    "documentformats"
    "ethnicities"
    "federalemploymentstatuses"
    "geoloccodes"
    "gsageoloccodes"
    "hiringpaths"
    "keystandardrequirements"
    "languagecodes"
    "languageproficiencies"
    "locationexpansions"
    "militarystatuscodes"
    "missioncriticalcodes"
    "occupationalseries"
    "payplans"
    "positionofferingtypes"
    "positionopeningstatuses"
    "positionscheduletypes"
    "postalcodes"
    "racecodes"
    "refereetypecodes"
    "remunerationrateintervalcodes"
    "requiredstandarddocuments"
    "securityclearances"
    "servicetypes"
    "specialhirings"
    "travelpercentages"
    "whomayapply"
)

# Iterate over each URL in the array and use curl to send a request
for endpoint in "${CODELIST_ENDPOINTS[@]}"; do
    echo "Sending request to $endpoint"
    curl -H "Host: data.usajobs.gov" \
        -H "User-Agent: $EMAIL" \
        -H "Authorization-Key: $TOKEN" \
        -X GET "https://data.usajobs.gov/api/codelist/${endpoint}" \
        | jq > testdata/${endpoint}-testdata.json
done
