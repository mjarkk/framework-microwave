package regex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	assert.True(t, Match(`ab?c`, "ac"), "Regex /ab?c/ sould match \"ac\"")
	assert.True(t, !Match(`ab?c`, "aacc"), "Regex /ab?c/ sould NOT match \"aacc\"")
}

func TestNormalMatch(t *testing.T) {
	assert.True(t, NormalMatch(`ab?c`, "ac"), "Regex /ab?c/ sould match \"ac\"")
	assert.True(t, NormalMatch(`ab?c`, "aacc"), "Regex /ab?c/ sould match \"aacc\"")
}
