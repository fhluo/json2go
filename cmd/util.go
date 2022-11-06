package cmd

import (
	"bytes"
	gen "github.com/dave/jennifer/jen"
	"github.com/fhluo/json2go/pkg/def"
	"github.com/tidwall/gjson"
)

func Generate(data []byte, packageName string, headerComment string, typeName string) ([]byte, error) {
	file := gen.NewFile(packageName)
	file.HeaderComment(headerComment)
	file.Add(def.Type(typeName, gjson.ParseBytes(data)))

	buf := new(bytes.Buffer)
	if err := file.Render(buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
