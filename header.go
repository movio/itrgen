package itrgen

import (
	"bytes"
	"html/template"
)

var header = `/*
 * CODE GENERATED AUTOMATICALLY WITH github.com/movio/itrgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package {{.Package}}

// {{.Name}} Iterator
type {{.Title}}Itr []{{.Name}}
`

func generateHeader(t Type) ([]byte, error) {
	tmpl, err := template.New("header").Parse(header)
	if err != nil {
		return nil, err
	}
	var result bytes.Buffer
	err = tmpl.Execute(&result, t)
	return result.Bytes(), err
}
