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
	"goCobra/utils"
	"fmt"
	"strings"
	"goCobra/models"
)

// editContactCmd represents the editContact command
var editContactCmd = &cobra.Command{
	Use:   "editContact",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args:cobra.MinimumNArgs(5),
	Run: editContact,
}

func editContact(cmd *cobra.Command, args []string) {
	if !utils.IsFileExist() { // Если файла еще нет, выводим инфу об этом и о том, что создание файла происходит при записи нового контакта.
		return
	}

	if !utils.CheckParamsExceptApartment(args[0], args[1], args[2], args[3], args[4]) { // Если аргументы не проходят валидацию (все, кроме квартиры, т.к. она опциональная), выводим инфу об этом и выходим
		return
	}

	// Извлекаем аргументы из командной строки и записываем в адрес
	address := []string{args[2], args[3], args[4]}

	if len(args) > 5 { // Если квартира указана, то
		address = append(address, args[5]) // добавляем квартиру в адрес
		if !utils.Validate(args[5], utils.BuildingOrApartmentFlag) { // и валидируем ее
			utils.PrintValidationMessages()
			return
		}
	}

	// Собираем контакт в структуру и отдаем на изменение (адрес в структуре является строкой)
	contact := models.NewContact(args[0], args[1], strings.Join(address, "-")) // Для удобочитаемости файла элементы адреса будут разделены символом "-"

	if utils.EditContact(contact) {
		fmt.Println("Запись успешно изменена")
	} else {
		fmt.Println("Запись не изменена")
	}
}

func init() {
	rootCmd.AddCommand(editContactCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editContactCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editContactCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
