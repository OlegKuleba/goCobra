package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"goCobra/models"
)

// Флаги, использующиеся для валидации
const FileName string = "phoneNumbers.txt"
const PhoneFlag string = "phone"
const NameFlag string = "name"
const CityFlag string = "city"
const StreetFlag string = "street"
const BuildingOrApartmentFlag string = "buildingOrApartment"

var PhoneRegexp *regexp.Regexp = regexp.MustCompile("^[+]380([0-9]{9})$")
var NameRegexp *regexp.Regexp = regexp.MustCompile("^[0-9A-Za-z]{3,15}$")
var CityRegexp *regexp.Regexp = regexp.MustCompile("^[A-Za-z]{3,15}$")
var StreetRegexp *regexp.Regexp = regexp.MustCompile("^[0-9A-Za-z]{4,20}$")
var BuildingOrApartmentRegexp *regexp.Regexp = regexp.MustCompile("^[0-9A-Za-z]{1,5}$")

func AddContact(contact models.Contact) bool {
	f, err := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	Check(err)
	defer f.Close()

	// Считываем данные из файла для проверки уникальности номера
	tmpF, err := ioutil.ReadFile(FileName)
	Check(err)
	rows := strings.Split(string(tmpF), ":")
	// Перебираем все строки из файла
	for _, row := range rows {
		// Если № совпадает (т.е. уже есть в файле), то сообщаем об этом и выходим ничего не меняя
		if strings.Contains(row, contact.Phone) {
			fmt.Println("Запись с таким номером уже существует. Для изменения данных используйте команду editContact")
			return false
		}
	}
	// Записываем в файл подготовленную строку
	f.WriteString(prepareString(contact) + "\n")
	return true
}

// Для удобочитаемости разбиваем поля структуры символом ":"
func prepareString(contact models.Contact) string {
	arr := []string{contact.Phone, contact.Name, contact.Address}
	return strings.Join(arr, ":")
}

func FindAll() {
	f, err := ioutil.ReadFile(FileName)
	Check(err)

	// Если файл пустой - выходим
	if IsFileEmpty(len(f)) {
		return
	}
	fmt.Println("Найдены записи:")
	fmt.Println(string(f))
}

func FindByNumber(number string) {
	f, err := ioutil.ReadFile(FileName)
	Check(err)
	rows := strings.Split(string(f), "\n")

	// Если файл пустой - выходим
	if IsFileEmpty(len(rows)) {
		return
	}

	// Перебираем все строки из файла
	for _, row := range rows {
		// Если текущий номер (из файла) совпал с введенным - выводим инфу о нем на дисплей
		if strings.Contains(row, number) {
			fmt.Println("Найдена запись:", row)
			return
		}
	}
	fmt.Println("Запись с таким номером не существует")
}

func EditContact(contact models.Contact) bool {
	f, err := ioutil.ReadFile(FileName)
	Check(err)
	rows := strings.Split(string(f), "\n")

	// Если файл пустой - выходим
	if IsFileEmpty(len(rows)) {
		return false
	}

	var rowItems []string
	// Перебираем все строки из файла
	for idx, row := range rows {
		// Разбиваем каждую строку на поля
		rowItems = strings.Split(row, ":")

		// Если текущий номер (из файла) совпал с новым (введенным) - входим в блок для замены информации по текущему номеру
		if strings.EqualFold(rowItems[0], contact.Phone) {
			rowItems[1] = contact.Name
			rowItems[2] = contact.Address
			rows[idx] = strings.Join(rowItems, ":")
			// Вписываем в найденную строку новые данные
			output := strings.Join(rows, "\n")
			// Записываем все в файл
			err = ioutil.WriteFile(FileName, []byte(output), 0644)
			Check(err)
			fmt.Println("Найдена запись:", row)
			fmt.Println("Изменена на:", rows[idx])
			return true
		}
	}
	fmt.Println("Запись с таким номером не существует")
	return false
}

func DeleteContact(number string) bool {
	f, err := ioutil.ReadFile(FileName)
	Check(err)
	rows := strings.Split(string(f), "\n")

	// Если файл пустой - выходим
	if IsFileEmpty(len(rows)) {
		return false
	}

	var rowItem []string
	var changedRows []string

	// Перебираем все строки из файла
	for idx, row := range rows {
		// Разбиваем каждую строку на поля
		rowItem = strings.Split(row, ":")

		// Если текущий номер (из файла) совпал с введенным - выводим инфу о нем на дисплей и удаляем
		if strings.EqualFold(rowItem[0], number) {
			fmt.Println("Найдена запись:", row)
			// Делаем срез массива от начального элемента до текущего
			changedRows = append(changedRows, rows[:idx]...)
			// Увеличиваем индекс на 1 для исключения текущего значения
			idx++
			// Добавляем в срез массива данные от текущего+1 элемента до последнего элемента
			changedRows = append(changedRows, rows[idx:]...)
			output := strings.Join(changedRows, "\n")
			// Пишем все в файл
			err = ioutil.WriteFile(FileName, []byte(output), 0644)
			Check(err)
			return true
		}
	}

	fmt.Println("Запись с таким номером не существует")
	return false
}

func CheckParamsExceptApartment(phone, name, city, street, building string) bool { // Если аргументы не проходят валидацию (все, кроме квартиры, т.к. она опциональная), выводим инфу об этом
	if Validate(phone, PhoneFlag) && Validate(name, NameFlag) && Validate(city, CityFlag) && Validate(street, StreetFlag) && Validate(building, BuildingOrApartmentFlag) {
		return true
	}
	PrintValidationMessages()
	return false
}

func IsFileEmpty(length int) bool {
	// Используется 2, т.к. в файле есть пустая строка (перенос)
	if length < 2 {
		fmt.Println("В файле нет записей")
		return true
	}
	return false
}

func IsFileExist() bool {
	if _, err := os.Stat(FileName); os.IsNotExist(err) {
		fmt.Println("Файл с номерами пока отсутствует. Для его создания совершите добавление записи")
		return false
	}
	return true
}

func Validate(data string, flag string) bool {
	switch flag {
	case PhoneFlag:
		return PhoneRegexp.MatchString(data)
	case NameFlag:
		return NameRegexp.MatchString(data)
	case CityFlag:
		return CityRegexp.MatchString(data)
	case StreetFlag:
		return StreetRegexp.MatchString(data)
	case BuildingOrApartmentFlag:
		return BuildingOrApartmentRegexp.MatchString(data)
	default:
		return false
	}
	//return false
}

func PrintValidationMessages() string {
	fmt.Println("Попробуйте еще раз. При этом введите верные данные")
	fmt.Println("Формат номера телефона +380xxYYYYYYY")
	fmt.Println("Имя - только буквы A-z и/или цифры (от 3 до 15 символов)")
	fmt.Println("Город - только буквы A-z (от 3 до 15 символов)")
	fmt.Println("Улица - только буквы A-z и/или цифры (от 4 до 20 символов)")
	fmt.Println("Дом/квартира - только буквы A-z и/или цифры (от 1 до 5 символов)")
	return "PrintValidationMessages()"
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
