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
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "usajobs",
	Short: "An Unofficial USAJobs CLI",
	Long: `
   ___     ___            _   _    ___     ___        _    ___     ___     ___  
  / __|   / _ \    ___   | | | |  / __|   /   \    _ | |  / _ \   | _ )   / __| 
 | (_ |  | (_) |  |___|  | |_| |  \__ \   | - |   | || | | (_) |  | _ \   \__ \ 
  \___|   \___/   _____   \___/   |___/   |_|_|   _\__/   \___/   |___/   |___/ 


usajobs is an unofficial, open source command-line tool to search
job openings announced on usajobs. The tool includes all the filtering options
available to users of the website via the API.

!note!: You must request an API Token from USAJobs to use this CLI. See Links below.

Links:
Obtain a USAJobs API Token here: https://developer.usajobs.gov/apirequest/
You can view this project at: 

Support:
Bugfix and feature requests should be submitted as an issue to the GitHub 
repository. There is no guarantee bugfix or feature requests will be implemented.

Contributions welcome!
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

var (
	userAgent string
	apiToken  string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&userAgent, "user-agent", "", "email address used when obtaining a usajobs api token (required)")
	rootCmd.PersistentFlags().StringVar(&apiToken, "token", "", "usajobs api token (required)")
	rootCmd.MarkPersistentFlagRequired("user-agent")
	rootCmd.MarkPersistentFlagRequired("token")
}
