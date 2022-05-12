package json2go

import (
	"regexp"
	"strconv"
	"strings"
)

var nameRegexp = regexp.MustCompile(`^[a-zA-Z_]\w*$`)

func toFieldName(s string) string {
	words := strings.Split(s, "_")

	for i, word := range words {
		if len(word) != 0 {
			words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}

	return strings.Join(words, "")
}

func fieldNameConverter() func(s string) string {
	counter := make(map[string]int)

	return func(s string) string {
		result := toFieldName(s)
		counter[result]++

		if counter[result] == 1 {
			return result
		} else {
			return result + strconv.Itoa(counter[result])
		}
	}
}

func validNames(items []string) bool {
	for _, item := range items {
		if !nameRegexp.MatchString(item) {
			return false
		}
	}
	return true
}

func validInts(items []string) bool {
	for _, item := range items {
		if _, err := strconv.Atoi(item); err != nil {
			return false
		}
	}
	return true
}
