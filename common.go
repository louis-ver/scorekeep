package main

import (
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
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
