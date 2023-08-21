package util

import "strconv"

func ValidateIdNumber(value string) (int64, bool) {
	var maxLength int = 20
	if isEmpty(value) {
		return 0, false
	}

	if len(value) > maxLength {
		return 0, false
	}

	int64Value, err := convertToInteger64(value)
	if err != nil {
		return 0, false
	}
	return int64Value, true
}

func isEmpty(value string) bool {
	return len(value) == 0
}

func convertToInteger64(value string) (int64, error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return int64(intValue), err
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
