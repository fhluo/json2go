package cmd

import (
	"bytes"
	gen "github.com/dave/jennifer/jen"
	"github.com/fhluo/json2go/pkg/def"
	"github.com/fhluo/json2go/pkg/scanner"
)

func Generate(data []byte, packageName string, headerComment string, typeName string, conv def.CamelCaseConverter) ([]byte, error) {
	file := gen.NewFile(packageName)
	file.HeaderComment(headerComment)

	c := def.Context{
		Scanner:            scanner.New(string(data)),
		CamelCaseConverter: conv,
	}

	t, err := c.Declare(typeName)
	if err != nil {
		return nil, err
	}

	file.Add(t)

	buf := new(bytes.Buffer)
	if err = file.Render(buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
