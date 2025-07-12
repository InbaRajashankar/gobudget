package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// StringToDateValues converts a M/D/Y string into a [3]int with the integer values.
// It also performs verification if the M and D values are valid.
func StringToDateValues(date_str string) ([3]int, error) {

	// Parse the string into an array of iingeters
	error_msg := "String %s cannot be parsed into a date!"
	parts := strings.Split(date_str, "/")
	if len(parts) != 3 {
		return [3]int{}, fmt.Errorf(error_msg+"Length != 3.", date_str)
	}

	var date_ints [3]int
	var err error
	for i := 0; i < 3; i++ {
		date_ints[i], err = strconv.Atoi(parts[i])
		if err != nil {
			return [3]int{}, fmt.Errorf(error_msg+"Failed cast to int.", date_str)
		}
	}

	// Check if the month is valid.
	if date_ints[0] < 1 || date_ints[0] > 12 {
		return [3]int{}, fmt.Errorf(error_msg+" Invalid Month.", date_str)
	}

	// Check if the day is valid.
	switch date_ints[0] {
	case 1, 3, 5, 7, 8, 10, 12: // 31-day months.
		if date_ints[1] < 1 || date_ints[1] > 31 {
			return [3]int{}, fmt.Errorf(error_msg+" Invalid Day.", date_str)
		}
	case 4, 6, 9, 11: // 33-day months.
		if date_ints[1] < 1 || date_ints[1] > 30 {
			return [3]int{}, fmt.Errorf(error_msg+" Invalid Day.", date_str)
		}
	case 2: // February...ignoring handling for leap years. :)
		if date_ints[1] < 1 || date_ints[1] > 29 {
			return [3]int{}, fmt.Errorf(error_msg+" Invalid Day.", date_str)
		}
	default:
		return [3]int{}, fmt.Errorf(error_msg+" Invalid Day.", date_str)
	}

	return date_ints, nil
}
