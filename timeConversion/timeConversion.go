package timeconversion

import (
	"fmt"
	"strconv"
	"strings"
)

func splitTime(arg string) (int32, int32, int32) {
	splitArr := strings.Split(arg, ":")
	h, _ := strconv.Atoi(splitArr[0])
	m, _ := strconv.Atoi(splitArr[1])
	s, _ := strconv.Atoi(splitArr[2])
	return int32(h), int32(m), int32(s)
}

func processAndConvert(arg int32) string {
	var convertedTimeValue string
	if arg > 9 && arg < 100 {
		convertedTimeValue = fmt.Sprintf("%01d", arg)
	} else if arg <= 9 {
		convertedTimeValue = fmt.Sprintf("%02d", arg)
	}
	return convertedTimeValue
}

func processHours(hourValue int32, suffix string) int32 {
	var hourValueConverted int32
	if suffix == "AM" {
		if hourValue == 12 {
			hourValueConverted = 0
		} else {
			hourValueConverted = hourValue
		}
	} else if suffix == "PM" {
		if hourValue == 12 {
			hourValueConverted = 12
		} else {
			hourValueConverted = hourValue + 12
		}
	}
	return hourValueConverted
}

func timeConversion(s string) string {
	var suffixTrimmedTime string
	var suffix string
	if strings.Contains(s, "AM") {
		suffixTrimmedTime = strings.TrimSuffix(s, "AM")
		suffix = "AM"
	} else if strings.Contains(s, "PM") {
		suffixTrimmedTime = strings.TrimSuffix(s, "PM")
		suffix = "PM"
	}

	hours, minutes, seconds := splitTime(suffixTrimmedTime)
	h := processHours(hours, suffix)
	return fmt.Sprintf("%s:%s:%s", processAndConvert(h), processAndConvert(minutes), processAndConvert(seconds))
}

func TimeConversionHelper(s string) string {
	return timeConversion(s)
}
