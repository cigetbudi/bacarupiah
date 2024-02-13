package bacarupiah

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	satuan = []string{"", "Satu", "Dua", "Tiga", "Empat", "Lima", "Enam", "Tujuh", "Delapan", "Sembilan"}
)

func GetTerbilang(number int) string {
	if number == 0 {
		return "Nol Rupiah"
	}

	result := ""

	if number >= 1000000 {
		result += BacaAngka(number/1000000) + " Juta "
		number %= 1000000
	}

	if number >= 100000 {
		if number < 200000 {
			result += "Seratus "
		} else {
			result += satuan[number/100000] + " Ratus "
		}
		number %= 100000
	}

	if number >= 1000 {
		result += BacaAngka(number/1000) + " Ribu "
		number %= 1000
	}

	if number >= 100 {
		result += satuan[number/100] + " Ratus "
		number %= 100
	}

	if number >= 20 {
		result += satuan[number/10] + " Puluh "
		number %= 10
	}

	if number > 0 {
		if number < 10 {
			result += satuan[number]
		} else if number == 10 {
			result += "Sepuluh"
		} else if number == 11 {
			result += "Sebelas"
		} else if number < 20 {
			result += satuan[number%10] + " Belas"
		}
	}

	return result
}

func BacaAngka[numberAcceptedType int | float64](numberGenerics numberAcceptedType) string {

	switch numberType := reflect.TypeOf(numberGenerics).String(); numberType {
	case "int":
		number := int(numberGenerics)

		result := GetTerbilang(number)

		return strings.TrimSpace(result)
		break
	case "float64":
		numberString := fmt.Sprintf("%v", numberGenerics)
		numberSplit := strings.Split(numberString, ".")

		number, _ := strconv.Atoi(numberSplit[0])
		koma, _ := strconv.Atoi(numberSplit[1])

		result := ""

		result += GetTerbilang(number)
		result += " Koma "
		result += GetTerbilang(koma)

		return strings.TrimSpace(result)
		break
	}

	return ""
}

func BacaRupiah[numberAcceptedType int | float64](numberGenerics numberAcceptedType) string {
	return BacaAngka(numberGenerics) + " Rupiah"
}
