package utils

import (
	"api/config"
	"strconv"
	"strings"
	"time"
)

func RequireCheck(listCheck []string, item, errors map[string]string) map[string]string {
	for _, key := range listCheck {
		_, ok := errors[key]

		if !ok && len(item[key]) <= 0 {
			errors[key] = config.GetMessageCode("REQUIRE")
		}
	}

	return errors

}

func MaxLengthCheck(listCheck []string, item, errors map[string]string) map[string]string {
	for _, key := range listCheck {
		itemKey := strings.Split(key, ":")

		maxlenth, _ := strconv.Atoi(itemKey[1])

		_, ok := errors[itemKey[0]]
		if !ok && len(item[itemKey[0]]) > maxlenth {
			errors[itemKey[0]] = config.GetMessageCode("MAX_LENGTH")
		}
	}

	return errors

}

func FixLengthCheck(listCheck []string, item, errors map[string]string) map[string]string {
	for _, key := range listCheck {
		itemKey := strings.Split(key, ":")

		maxlenth, _ := strconv.Atoi(itemKey[1])

		_, ok := errors[itemKey[0]]
		if !ok && len(item[itemKey[0]]) != maxlenth {
			errors[itemKey[0]] = config.GetMessageCode("FIX_LENGTH")
		}
	}

	return errors

}

// Date valid YYYY-MM-DD
func DateFormatCheck(listCheck []string, item, errors map[string]string) map[string]string {
	for _, key := range listCheck {
		_, ok := errors[key]

		_, err := time.Parse("2006-01-02", item[key])

		if !ok && err != nil {
			errors[key] = config.GetMessageCode("FORMAT_DATE")
		}
	}

	return errors

}

func NumberCheck(listCheck []string, item, errors map[string]string) map[string]string {
	for _, key := range listCheck {
		_, ok := errors[key]

		_, err := strconv.Atoi(item[key]);

		if !ok && err != nil {
			errors[key] = config.GetMessageCode("FORMAT_NUMBER")
		}
	}

	return errors

}