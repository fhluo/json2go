package json2go

import (
	"github.com/dave/jennifer/jen"
	"github.com/fhluo/json2go/pkg/json2go/def"
)

var DefaultOptions = Options{TypeName: "T"}

type Options struct {
	TypeName     string
	AllCapsWords []string
}

// GenerateCode generates a Go type definition from a JSON string as jen.Code.
func (o Options) GenerateCode(json string) (jen.Code, error) {
	return def.From(json, o.AllCapsWords...).Declare(o.TypeName)
}

// GenerateCodeFromBytes generates a Go type definition from a slice of JSON bytes as jen.Code.
func (o Options) GenerateCodeFromBytes(json []byte) (jen.Code, error) {
	return def.FromBytes(json, o.AllCapsWords...).Declare(o.TypeName)
}

// Generate generates a Go type definition from a JSON string.
func (o Options) Generate(json string) (string, error) {
	code, err := o.GenerateCode(json)
	if err != nil {
		return "", err
	}

	return RenderCode(code)
}

// GenerateFromBytes generates a Go type definition from a slice of JSON bytes.
func (o Options) GenerateFromBytes(json []byte) (string, error) {
	code, err := o.GenerateCodeFromBytes(json)
	if err != nil {
		return "", err
	}

	return RenderCode(code)
}

// GenerateCode generates a Go type definition from a JSON string as jen.Code.
func GenerateCode(json string) (jen.Code, error) {
	return DefaultOptions.GenerateCode(json)
}

// GenerateCodeFromBytes generates a Go type definition from a slice of JSON bytes as jen.Code.
func GenerateCodeFromBytes(json []byte) (jen.Code, error) {
	return DefaultOptions.GenerateCodeFromBytes(json)
}

// Generate generates a Go type definition from a JSON string.
func Generate(json string) (string, error) {
	return DefaultOptions.Generate(json)
}

// GenerateFromBytes generates a Go type definition from a slice of JSON bytes.
func GenerateFromBytes(json []byte) (string, error) {
	return DefaultOptions.GenerateFromBytes(json)
}
