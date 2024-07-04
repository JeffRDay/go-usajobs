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
	"encoding/csv"
	"errors"
	"net/http"
	"os"

	usajobs "github.com/JeffRDay/go-usajobs/client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// applicationstatusesCmd represents the applicationstatuses command
var applicationstatusesCmd = &cobra.Command{
	Use:   "applicationstatuses",
	Short: "lists the academic honors tracked by usajobs",
	Long: `
lists the application statuses tracked by usajobs and can be used to refine job
searches or inform resume uploads to the website itself.

Example: 
usajobs list applicationstatuses

Output:
┌───────┬────────────────────────────────────────────────────────────────────────────────┐
│ CODE  │ VALUE                                                                          │
├───────┼────────────────────────────────────────────────────────────────────────────────┤
│ 11112 │ Applicant Declined Appointment                                                 │
├───────┼────────────────────────────────────────────────────────────────────────────────┤
│ 11312 │ Minimum Qualification Requirements Met                                         │
├───────┼────────────────────────────────────────────────────────────────────────────────┤
│ 11313 │ Multiple Grades/Status/Agencies – Eligible-Application Not Referred to Selecti │
│       │ ng Official                                                                    │
└───────┴────────────────────────────────────────────────────────────────────────────────┘
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := executeApplicationStatuses()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to execute applicationstatuses command")
		}
	},
}

func init() {
	listCmd.AddCommand(applicationstatusesCmd)
}

func executeApplicationStatuses() error {

	var err error
	if Client == nil {
		Client, err = usajobs.NewClient("not", "required")
		if err != nil {
			return err
		}
	}

	r, data, err := Client.ApplicationStatuses.WithOptions(nil)
	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return errors.New("bad response from usajobs: " + r.Status)
	}

	headersSummary := []string{"CODE", "VALUE"}
	var dataSummary [][]string
	for _, item := range data.CodeList {
		for _, i := range item.ValidValue {
			dataSummary = append(dataSummary, []string{addNewLines(i.Code, 80), addNewLines(i.Value, 80)})
		}
	}

	headersDetails := []string{"ID", "CODE", "VALUE", "LAST_MODIFIED", "IS_DISABLED", "DATE_GENERATED"}
	var dataDetails [][]string
	for _, item := range data.CodeList {
		for _, i := range item.ValidValue {
			dataDetails = append(dataDetails,
				[]string{
					addNewLines(item.ID, 80),
					addNewLines(i.Code, 80),
					addNewLines(i.Value, 80),
					addNewLines(i.LastModified, 80),
					addNewLines(i.IsDisabled, 80),
					addNewLines(data.DateGenerated, 80),
				})
		}
	}

	switch display {
	case "summary":
		err = displayTable(headersSummary, dataSummary)
		if err != nil {
			return err
		}
	case "detail":
		err = displayTable(headersDetails, dataDetails)
		if err != nil {
			return err
		}
	case "csv":
		writer := csv.NewWriter(os.Stdout)

		err := writer.Write(headersDetails)
		if err != nil {
			return err
		}

		err = writer.WriteAll(dataDetails)
		if err != nil {
			return err
		}

		writer.Flush()

		if err := writer.Error(); err != nil {
			return err
		}
	default:
		err = displayTable(headersSummary, dataSummary)
		if err != nil {
			return err
		}
	}

	return nil
}
