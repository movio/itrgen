package itrgen

import (
	"bytes"
	"html/template"
)

var mapTemplate = `
// Map{{.ResultTitle}} takes a mapping function, invokes the mapping function
// for each element in itr, and returns the results as a slice.
// If the mapping function returns err, Map{{.ResultTitle}} immediately returns nil, err.
func (itr {{.Title}}Itr) Map{{.ResultTitle}}(fn func({{.Name}}) ({{.ResultName}}, error)) ([]{{.ResultName}}, error) {
	results := []{{.ResultName}}{}
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

// MapType holds data needed to render the code generation template
type MapType struct {
	Title       string
	Name        string
	ResultTitle string
	ResultName  string
}

func generateMap(t MapType) ([]byte, error) {
	tmpl, err := template.New("map").Parse(mapTemplate)
	if err != nil {
		return nil, err
	}
	var result bytes.Buffer
	err = tmpl.Execute(&result, t)
	return result.Bytes(), err
}
