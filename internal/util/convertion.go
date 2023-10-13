package util

import "strconv"

func ConvertStringToInt32(value string) int32 {
	data, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return int32(data)
}

func ConvertStringToFloat64(value string) float64 {
	data, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	return data
}
