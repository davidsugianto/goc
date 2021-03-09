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

	"github.com/spf13/cobra"
)

var env, cloud string

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Display command for deploy to multiple cloud and multiple env",
	Long: `The subcommands directive 'goc deploy' for deploy to multiple cloud and multiple environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		env, _:= cmd.Flags().GetString("env")
		cloud, _:= cmd.Flags().GetString("cloud")

		fmt.Println("Deploying current application artifacts to "+env+" environment on "+cloud+" cloud")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	
	deployCmd.Flags().StringVarP(&env, "env", "e", "", "Enter the environment to deploy applications")
	deployCmd.Flags().StringVarP(&cloud, "cloud", "c", "aws", "Enter the Cloud Provider")
	
	deployCmd.MarkFlagRequired("env")
	deployCmd.MarkFlagRequired("cloud")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
