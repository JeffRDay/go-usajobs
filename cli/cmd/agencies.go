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
	"errors"
	"fmt"
	"net/http"
	"os"

	usajobs "github.com/JeffRDay/go-usajobs/client"
	"github.com/spf13/cobra"
)

// agenciesCmd represents the agencies command
var agenciesCmd = &cobra.Command{
	Use:   "agencies",
	Short: "list the agency and agency code as tracked by USAJobs",
	Long: `
lists the agencies registered on USAJobs and their agency code. Use the code
in future usajobs search commands to better refine job searches.

Example:
usajobs list agencies | grep Army`,
	Run: func(cmd *cobra.Command, args []string) {
		r, err := executeListAgencies(nil)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}

		for _, i := range r {
			fmt.Println(i)
		}
	},
}

func init() {
	listCmd.AddCommand(agenciesCmd)
}

func executeListAgencies(opt *usajobs.AgencySubelementsOptions) ([]string, error) {

	var err error
	if Client == nil {
		Client, err = usajobs.NewClient("not required", "for this endpoint")
		if err != nil {
			return nil, err
		}
	}

	response, r, err := Client.Agency.WithOptions(opt)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	var results []string
	for _, agencies := range r.CodeList {
		for _, agency := range agencies.ValidValue {
			results = append(results, fmt.Sprintf("%s\t%s", agency.Code, agency.Value))
		}
	}

	return results, nil
}
