package def

import (
	"golang.org/x/exp/maps"
	"regexp"
	"strings"
)

var (
	allCapsRe      = regexp.MustCompile(`^[A-Z]{2,}\d*$`)
	capitalizedRe  = regexp.MustCompile(`[A-Z][a-z]+\d*`)
	followNonCapRe = regexp.MustCompile(`([a-z\d])([A-Z])`)
	followCapRe    = regexp.MustCompile(`([A-Z])([A-Z][a-z]+)`)
	splitRe        = regexp.MustCompile(`_|\s+`)
)

var (
	ALLCaps = map[string]bool{"ID": true}
	Cache   = make(map[string]string)
)

func GetAcronyms() []string {
	return maps.Keys(ALLCaps)
}

func SetAcronyms(acronyms ...string) {
	ALLCaps = make(map[string]bool)
	for _, acronym := range acronyms {
		ALLCaps[acronym] = true
	}
}

func ToCamelCase(s string) string {
	if r, ok := Cache[s]; ok {
		return r
	}

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

	r := capitalizedRe.ReplaceAllStringFunc(strings.Join(words, ""), func(s string) string {
		if word := strings.ToUpper(s); ALLCaps[word] {
			return word
		}
		return s
	})

	Cache[s] = r
	return r
}

func ToSnakeCase(s string) string {
	s = followNonCapRe.ReplaceAllString(s, "${1}_$2")
	s = followCapRe.ReplaceAllString(s, "${1}_$2")
	return strings.ToLower(s)
}

func ValidNames(items []string) bool {
	for _, item := range items {
		if item == "" {
			return false
		}

		if !(('a' <= item[0] && item[0] <= 'z') || ('A' <= item[0] && item[0] <= 'Z') || item[0] == '_') {
			return false
		}

		for i := 1; i < len(item); i++ {
			if !(('a' <= item[i] && item[i] <= 'z') || ('A' <= item[i] && item[i] <= 'Z') || ('0' <= item[i] && item[i] <= '9') || item[i] == '_') {
				return false
			}
		}
	}

	return true
}

func ValidIntegers(items []string) bool {
	for _, item := range items {
		if !isInteger(item) {
			return false
		}
	}
	return true
}
