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
	"goCobra/models"
	"strings"
	"fmt"
)

// addContactCmd represents the addContact command
var addContactCmd = &cobra.Command{
	Use:   "addContact",
	Short: "addContact phoneNumber name - добавляет новый контакт в файл",
	Long: `Команда добавляет новый контакт в файл, требует передачу минимум 5-ти, максимум 6-ти аргументов:
номерТелефона имяВладельца город улица дом [квартира](в указанном порядке)
параметры в квадратных скобках являются опциональными
Например: addContact +380501234567 Anya Dnipro Gagarina 103a 5`,
	Args: cobra.MinimumNArgs(5),
	Run:  addContact,
}

func addContact(cmd *cobra.Command, args []string) {
	if !utils.CheckParamsExceptApartment(args[0], args[1], args[2], args[3], args[4]) { // Если аргументы не проходят валидацию (все, кроме квартиры, т.к. она опциональная), выводим инфу об этом и выходим
		return
	}

	// Извлекаем аргументы из командной строки и записываем в адрес
	address := []string{args[2], args[3], args[4]}

	if len(args) > 5 { // Если квартира указана, то
		address = append(address, args[5]) // добавляем квартиру в адрес
		if !utils.Validate(args[5], utils.BuildingOrApartmentFlag) { // и валидируем ее. Если не валидно - выходим
			utils.PrintValidationMessages()
			return
		}
	}

	// Собираем контакт в структуру и отдаем на запись (адрес в структуре является строкой)
	contact := models.NewContact(args[0], args[1], strings.Join(address, "-")) // Для удобочитаемости файла элементы адреса будут разделены символом "-"
	if utils.AddContact(contact) {
		fmt.Println("Запись успешно добавлена")
	} else {
		fmt.Println("Запись не добавлена")
	}
}

func init() {
	rootCmd.AddCommand(addContactCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addContactCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addContactCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
