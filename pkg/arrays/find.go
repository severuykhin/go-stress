package arrays

import "regexp"

func FindString(value string, slice []string) int {
	index := -1
	for i, v := range slice {
		if v == value {
			index = i
			break
		}
	}
	return index
}

func FindStringByRegex(regex *regexp.Regexp, slice []string) int {
	index := -1
	for i, v := range slice {
		if regex.Match([]byte(v)) {
			index = i
			break
		}
	}
	return index
}
