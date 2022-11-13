package cmd

import (
	"bytes"
	gen "github.com/dave/jennifer/jen"
	"github.com/fhluo/json2go/pkg/def"
)

func Generate(data []byte, packageName string, headerComment string, typeName string) ([]byte, error) {
	file := gen.NewFile(packageName)
	file.HeaderComment(headerComment)

	t, err := def.FromBytes(data).Declare(typeName)
	if err != nil {
		return nil, err
	}

	file.Add(t)

	buf := new(bytes.Buffer)
	if err := file.Render(buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
