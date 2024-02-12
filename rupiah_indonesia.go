package bacarupiah

import (
	"strings"
)

func RupiahIndonesia(number int) string {

	var (
		satuan = []string{"", "Satu", "Dua", "Tiga", "Empat", "Lima", "Enam", "Tujuh", "Delapan", "Sembilan"}
	)

	if number == 0 {
		return "Nol Rupiah"
	}

	result := ""

	if number >= 1000000 {
		result += RupiahIndonesia(number/1000000) + " Juta "
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
		result += RupiahIndonesia(number/1000) + " Ribu "
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

	return strings.TrimSpace(result)
}
