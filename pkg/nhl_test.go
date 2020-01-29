package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamNameToResourceNameWithPeriod(t *testing.T) {
	expect := "st-louis-blues"
	got := TeamNameToResourceName("St. Louis Blues")

	assert.Equal(t, expect, got)
}

func TestTeamNameToResourceNameWithAccent(t *testing.T) {
	expect := "montreal-canadiens"
	got := TeamNameToResourceName("Montréal Canadiens")

	assert.Equal(t, expect, got)
}
