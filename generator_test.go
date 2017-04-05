package itrgen

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	str := NewType("main", "string", []string{"string", "MyStruct"})
	code, err := generateCode(str)
	assert.NoError(t, err)
	assert.Equal(t, string(getExpected(t, "string")), string(code))
}

func TestStruct(t *testing.T) {
	str := NewType("custom", "MyStruct", []string{"string", "int64"})
	code, err := generateCode(str)
	assert.NoError(t, err)
	assert.Equal(t, string(getExpected(t, "struct")), string(code))
}

func getExpected(t *testing.T, name string) []byte {
	dat, err := ioutil.ReadFile("./expected_" + name + ".txt")
	if err != nil {
		t.Error("could not find expected file for " + name)
	}

	return dat
}
