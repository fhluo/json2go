// Package json2go provides generating a Go type definition from JSON.
package json2go

import (
	"bytes"
	gen "github.com/dave/jennifer/jen"
	"github.com/tidwall/gjson"
)

func Generate(data []byte, packageName string, headerComment string, typeName string) ([]byte, error) {
	file := gen.NewFile(packageName)
	file.HeaderComment(headerComment)
	file.Add(Type(typeName, gjson.ParseBytes(data)))

	buf := new(bytes.Buffer)
	if err := file.Render(buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
