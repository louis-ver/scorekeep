package library

import (
	"encoding/json"
	"net/http"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func DecodeJSON(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(target)

	return err
}

func teamNameToResourceName(teamName string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r)
	}))

	result, _, _ := transform.String(t, teamName)

	result = strings.ReplaceAll(result, " ", "-")
	result = strings.ReplaceAll(result, ".", "")

	return strings.ToLower(result)
}
