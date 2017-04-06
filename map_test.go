package itrgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapToString(t *testing.T) {

	expected := `
// MapMyStruct takes a mapping function, invokes the mapping function
// for each element in itr, and returns the results as a slice.
// If the mapping function returns err, MapMyStruct immediately returns nil, err.
func (itr StringItr) MapMyStruct(fn func(string) (MyStruct, error)) ([]MyStruct, error) {
	results := []MyStruct{}
	for _, i := range itr {
		result, err := fn(i)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
`

	str := MapType{"String", "string", "MyStruct", "MyStruct"}
	code, err := generateMap(str)
	assert.NoError(t, err)
	assert.Equal(t, string(expected), string(code))
}
