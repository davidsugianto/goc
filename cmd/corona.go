/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"encoding/json"
	"io/ioutil"
	"net/http"
    "os"

	"github.com/spf13/cobra"
)

type Response struct {
	Countries	[]Countries	`json:"countries"`
}

type Countries struct {
	Name	string	`json:"name"`
	Iso2	string	`json:"iso2"`
	Iso3	string	`json:"iso3"`
}

// coronaCmd represents the corona command
var coronaCmd = &cobra.Command{
	Use:   "corona",
	Short: "Display command for the Coronavirus Disease 19 Data",
	Long: `Coronavirus Disease 19 Data`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("corona called")
	},
}

// countryCmd represents the coro subcommands
var countryCmd = &cobra.Command{
	Use: "country",
	Short: "Display command for the show all Country affected by covid19",
	Long: `All Country affected by Covid-19`,
	Run: func (cmd *cobra.Command, args []string) {
		getCountry()
	},
}

func init() {
	rootCmd.AddCommand(coronaCmd)

	coronaCmd.AddCommand(countryCmd)
}

func getCountry() {
	response, err := http.Get("https://covid19.mathdro.id/api/countries")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject.Countries); i++  {
		fmt.Println(responseObject.Countries[i].Name)
	}

}
