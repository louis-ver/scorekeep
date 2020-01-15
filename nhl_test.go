package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamNameToResourceNameWithPeriod(t *testing.T) {
	expect := "st-louis-blues"
	got := teamNameToResourceName("St. Louis Blues")

	assert.Equal(t, expect, got)
}

func TestTeamNameToResourceNameWithAccent(t *testing.T) {
	expect := "montreal-canadiens"
	got := teamNameToResourceName("Montréal Canadiens")

	assert.Equal(t, expect, got)
}

func TestStringInSlice(t *testing.T) {
	expect := true
	got := stringInSlice("a-string", []string{"a-string"})

	assert.Equal(t, expect, got)
}

func TestStringNotInSlice(t *testing.T) {
	expect := false
	got := stringInSlice("a-string", []string{"not-in-slice"})

	assert.Equal(t, expect, got)
}
