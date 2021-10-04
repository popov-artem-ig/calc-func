package utils

import (
	"strconv"
)

func IsDigit(c uint8) bool {
	return c >= '0' && c <= '9'
}

func IsFloat(s string) bool {
	const bitSize = 64 // Don't think about it to much. It's just 64 bits.
	_, err := strconv.ParseFloat(s, bitSize)
	if err == nil {
		return true
	}
	return false
}

func RpnStrToArr(rpnString string) []string {
	hoarder := ""                  // Накопитель для сборки/обработки дробных и больших чисел
	var result []string

	for i := range rpnString {
		c := rpnString[i]

		if IsDigit(c) || c == '.' {
			hoarder += string(c)
			if i != len(rpnString)-1 {
				continue
			}
		}

		if len(hoarder) > 0 {
			result = append(result, hoarder)
			hoarder = ""
		}

		if c == ' ' {
			continue
		} else {
			result = append(result, string(c))
		}
	}

	return result
}





