package def

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	allCapsRe      = regexp.MustCompile(`^[A-Z]{2,}\d*$`)
	capitalizedRe  = regexp.MustCompile(`[A-Z][a-z]+\d*`)
	followNonCapRe = regexp.MustCompile(`([a-z\d])([A-Z])`)
	followCapRe    = regexp.MustCompile(`([A-Z])([A-Z][a-z]+)`)
	splitRe        = regexp.MustCompile(`_|\s+`)
)

var ALLCaps = map[string]bool{"ID": true}

func SetAcronyms(acronyms ...string) {
	for _, acronym := range acronyms {
		ALLCaps[acronym] = true
	}
}

func ToCamelCase(s string) string {
	words := splitRe.Split(s, -1)

	for i, word := range words {
		if len(word) != 0 {
			if allCapsRe.MatchString(word) {
				words[i] = word[:1] + strings.ToLower(word[1:])
			} else {
				words[i] = strings.ToUpper(word[:1]) + word[1:]
			}
		}
	}

	return capitalizedRe.ReplaceAllStringFunc(strings.Join(words, ""), func(s string) string {
		if word := strings.ToUpper(s); ALLCaps[word] {
			return word
		}
		return s
	})
}

func ToSnakeCase(s string) string {
	s = followNonCapRe.ReplaceAllString(s, "${1}_$2")
	s = followCapRe.ReplaceAllString(s, "${1}_$2")
	return strings.ToLower(s)
}

var nameRe = regexp.MustCompile(`^[a-zA-Z_]\w*$`)

func validNames(items []string) bool {
	for _, item := range items {
		if !nameRe.MatchString(item) {
			return false
		}
	}
	return true
}

func validIntegers(items []string) bool {
	for _, item := range items {
		if _, err := strconv.Atoi(item); err != nil {
			return false
		}
	}
	return true
}
