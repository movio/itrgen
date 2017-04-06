package itrgen

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Type represents the golang type that you are generating functions for
type Type struct {
	Package string
	Title   string
	Name    string
	Maps    []string
}

// NewType constructs a Type, with the Title attribute set to TitleCase `name`.
// i.e. generated methods for `mytype` will be exported as `MytypeMap`.
func NewType(pkg string, name string, maps []string) Type {
	return Type{
		pkg,
		strings.Title(name),
		name,
		maps,
	}
}

// Generate creates iterator code for Type t and writes it to ./(typename)_itr.go
func Generate(t Type) error {
	// Get code
	code, err := generateCode(t)
	if err != nil {
		return err
	}

	// Write file
	f := filepath.Join(".", fmt.Sprintf("%v_itr.go", strings.ToLower(t.Name)))
	write(f, code)
	return nil
}

func generateCode(t Type) ([]byte, error) {
	// Get code
	code, err := generateHeader(t)
	if err != nil {
		return nil, err
	}

	for _, v := range t.Maps {
		c, err := generateMap(MapType{t.Title, t.Name, strings.Title(v), v})
		if err != nil {
			return nil, err
		}

		code = append(code, c...)
	}

	return code, nil
}
