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

func TestFindMatch(t *testing.T) {
	matcher := `((aa)|(bb)|(cc))`
	testS := "_aabbcc_"
	assert.Equal(t, FindMatch(testS, matcher, 1), "aa", "FindMatch finds corect match (group 1)")
	assert.Equal(t, FindMatch(testS, matcher, 2), "aa", "FindMatch finds corect match (group 2)")
	assert.Equal(t, FindMatch(testS, matcher, 3), "", "Find no match on group 3")
}

func TestFindAllMatch(t *testing.T) {
	matcher := `((aa)|(bb)|(cc))`
	testS := "_aabbcc_"
	assert.Equal(t, FindAllMatch(testS, matcher, 1), []string{"aa", "bb", "cc"}, "FindAllMatch finds corect match #1")
}

func TestReplace(t *testing.T) {
	testVal := "foo bar"
	replacedVal := Replace(testVal, " baz", `\sbar`)
	assert.NotEqual(t, testVal, replacedVal, "replaced variable is not the same as `foo bar`")
	assert.Equal(t, "foo baz", replacedVal, "new replaced val is `foo baz`")
}
