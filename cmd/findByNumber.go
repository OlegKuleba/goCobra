// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/OlegKuleba/goCobra/utils"
)

// findByNumberCmd represents the findByNumber command
var findByNumberCmd = &cobra.Command{
	Use:   "findByNumber",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run:  findByNumber,
}

func findByNumber(cmd *cobra.Command, args []string) {
	if !utils.IsFileExist() { // Если файла еще нет, выводим инфу об этом и о том, что создание файла происходит при записи нового контакта.
		return
	}
	if !utils.Validate(args[0], utils.PhoneFlag) { // Если аргумент не проходит валидацию, выводим инфу об этом и выходим
		utils.PrintValidationMessages()
		return
	}

	utils.FindByNumber(args[0])
}

func init() {
	rootCmd.AddCommand(findByNumberCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findByNumberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findByNumberCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
