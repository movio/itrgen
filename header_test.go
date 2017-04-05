package itrgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicHeader(t *testing.T) {

	expected := `/*
 * CODE GENERATED AUTOMATICALLY WITH github.com/movio/itrgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package main

// string Iterator
type StringItr []string
`

	stringItr := NewType("main", "string", nil)
	res, err := generateHeader(stringItr)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(res))
}
