package converter

import (
	"flagger/micro"
	"strconv"
)

func IntToString(input int) string {
	return strconv.Itoa(input)
}

func FloatToString(input float64) string {
	return strconv.FormatFloat(input, 'f', -1, 64)
}

func BoolToString(input bool) string {
	return strconv.FormatBool(input)
}

func StringToInt(input string) int {
	i, err := strconv.Atoi(input)
	micro.CheckError(err)
	return i
}

func StringToFloat(input string) float64 {
	f, err := strconv.ParseFloat(input, 64)
	micro.CheckError(err)
	return f
}

func StringToBool(input string) bool {
	b, err := strconv.ParseBool(input)
	micro.CheckError(err)
	return b
}
