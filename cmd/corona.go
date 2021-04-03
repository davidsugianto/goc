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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Corona Global Cases
type ResponseGlobal struct {
	Confirmed  Confirmed `json:"confirmed"`
	Recovered  Recovered `json:"recovered"`
	Deaths     Deaths    `json:"deaths"`
	LastUpdate string    `json:"lastUpdate"`
}

type Confirmed struct {
	Value int `json:"value"`
}

type Recovered struct {
	Value int `json:"value"`
}

type Deaths struct {
	Value int `json:"value"`
}

// Corona Country Cases
type ResponseCountryCase struct {
	CConfirmed  CConfirmed `json:"confirmed"`
	CRecovered  CRecovered `json:"recovered"`
	CDeaths     CDeaths    `json:"deaths"`
	CLastUpdate string     `json:"lastUpdate"`
}

type CConfirmed struct {
	Value int `json:"value"`
}

type CRecovered struct {
	Value int `json:"value"`
}

type CDeaths struct {
	Value int `json:"value"`
}

// Corona Country Lists
type ResponseCountries struct {
	Countries []Countries `json:"countries"`
}

type Countries struct {
	Name string `json:"name"`
	Iso2 string `json:"iso2"`
	Iso3 string `json:"iso3"`
}

var country string

// coronaCmd represents the corona command
var coronaCmd = &cobra.Command{
	Use:   "corona",
	Short: "Display command for the Coronavirus Disease 19 Data, Data sources from Muhammad Mustadi's API",
	Long:  `Coronavirus Disease 19 Data`,
	Run: func(cmd *cobra.Command, args []string) {
		country, _ := cmd.Flags().GetString("country")
		if country != "" {
			getCountryCase()
		} else {
			getGlobalCase()
		}
	},
}

// countryCmd represents the coro subcommands
var countryCmd = &cobra.Command{
	Use:   "country",
	Short: "Display command for the show all Country affected by covid19",
	Long:  `All Country affected by Covid-19`,
	Run: func(cmd *cobra.Command, args []string) {
		getCountry()
	},
}

func init() {
	rootCmd.AddCommand(coronaCmd)

	coronaCmd.AddCommand(countryCmd)
	coronaCmd.Flags().StringVarP(&country, "country", "c", "", "Covid-19 affected countries")
}

func getGlobalCase() {
	p := message.NewPrinter(language.English)
	response, err := http.Get("https://covid19.mathdro.id/api")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var responseObject ResponseGlobal
	json.Unmarshal(responseData, &responseObject)

	lastTimeUpdate := (responseObject.LastUpdate)
	t, err := time.Parse(time.RFC3339, lastTimeUpdate)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("--- Coronavirus 19 Disease Global Cases ---")
	p.Printf("Status Confirmed : %d\n", responseObject.Confirmed.Value)
	p.Printf("Status Recovered : %d\n", responseObject.Recovered.Value)
	p.Printf("Status Deaths : %d\n", responseObject.Deaths.Value)
	fmt.Println("Last Updated : ", t)

}

func getCountryCase() {
	p := message.NewPrinter(language.English)
	response, err := http.Get("https://covid19.mathdro.id/api/countries/" + country)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var responseObject ResponseCountryCase
	json.Unmarshal(responseData, &responseObject)

	lastTimeUpdate := (responseObject.CLastUpdate)
	t, err := time.Parse(time.RFC3339, lastTimeUpdate)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("--- Coronavirus 19 Disease " + country + " Cases ---")
	fmt.Println("Country :", country)
	p.Printf("Status Confirmed : %d\n", responseObject.CConfirmed.Value)
	p.Printf("Status Recovered : %d\n", responseObject.CRecovered.Value)
	p.Printf("Status Deaths : %d\n", responseObject.CDeaths.Value)
	fmt.Println("Last Updated : ", t)
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

	var responseObject ResponseCountries
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject.Countries); i++ {
		fmt.Println(responseObject.Countries[i].Name)
	}

}
