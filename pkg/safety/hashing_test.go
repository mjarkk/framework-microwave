package safety

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAutoPasswordHash(t *testing.T) {
	testPass := "somePassword"
	pass1, salt1 := AutoPasswordHash(testPass)
	pass2, salt2 := AutoPasswordHash(testPass)
	assert.NotEqual(t, "", pass1, "Return password should not be empty")
	assert.NotEqual(t, "", salt1, "Return salt should not be empty")
	assert.NotEqual(t, pass1, pass2, "Same inputs sould not return the same hashes")
	assert.NotEqual(t, salt1, salt2, "Same inputs sould not return the same salt")
}

func TestHash(t *testing.T) {
	assert.Equal(t, Hash("md5", "test"), "098f6bcd4621d373cade4e832627b4f6", "md5 hasing returns valid md5 string")
	assert.Equal(t, Hash("sha512", "test"), "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff", "md5 hasing returns valid md5 string")
	assert.Equal(t, Hash("sha256", "test"), "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "md5 hasing returns valid md5 string")
	assert.Equal(t, Hash("SomeNotExsistingAlgorithm", "test"), "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "md5 hasing returns valid md5 string")
}

func TestValidHashTypes(t *testing.T) {
	for _, item := range ValidHashTypes {
		assert.NotEqual(t, "", item, "ValidHashTypes doesn't contain empty value's")
	}
}
