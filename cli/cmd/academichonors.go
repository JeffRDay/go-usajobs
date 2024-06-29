/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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

// academichonorsCmd represents the academichonors command
var academichonorsCmd = &cobra.Command{
	Use:   "academic-honors",
	Short: "list the academic honors tracked by USAJobs",
	Long: `
lists the academic honors registered on USAJobs and their agency code. Use the code
in future usajobs search commands to better refine job searches.

Example:
usajobs list academic-honors`,
	Run: func(cmd *cobra.Command, args []string) {
		r, err := executeListAcademicHonors(nil)
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
	listCmd.AddCommand(academichonorsCmd)
}

func executeListAcademicHonors(opt *usajobs.AcademicHonorsOptions) ([]string, error) {

	var err error
	if Client == nil {
		Client, err = usajobs.NewClient("not required", "for this endpoint")
		if err != nil {
			return nil, err
		}
	}

	response, r, err := Client.AcademicHonors.WithOptions(opt)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	var results []string
	for _, honors := range r.CodeList {
		for _, honor := range honors.ValidValue {
			results = append(results, fmt.Sprintf("%s", honor.Code))
		}
	}

	return results, nil
}
