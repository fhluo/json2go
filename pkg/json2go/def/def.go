package def

import (
	"bytes"
	"encoding/json/jsontext"
	"fmt"
	"slices"
	"strings"

	gen "github.com/dave/jennifer/jen"
	"github.com/fhluo/json2go/pkg/json2go/conv"
	"github.com/fhluo/json2go/pkg/xiter"
	"github.com/samber/lo"
)

type Context struct {
	conv.CamelCaseConverter
	*jsontext.Decoder
}

func From(s string, allCaps ...string) *Context {
	c := new(Context)
	c.CamelCaseConverter = conv.NewDefaultCamelCaseConverter(allCaps)
	c.Decoder = jsontext.NewDecoder(bytes.NewBufferString(s))
	return c
}

func FromBytes(data []byte, allCaps ...string) *Context {
	c := new(Context)
	c.CamelCaseConverter = conv.NewDefaultCamelCaseConverter(allCaps)
	c.Decoder = jsontext.NewDecoder(bytes.NewBuffer(data))
	return c
}

func (c *Context) Declare(name string) (*gen.Statement, error) {
	t, err := c.Type()
	if err != nil {
		return nil, err
	}
	return gen.Type().Id(name).Add(t.Code()), nil
}

func (c *Context) object() (keys []string, types []Type, err error) {
	if c.PeekKind() != '{' {
		err = fmt.Errorf("expected '{'")
		return
	}
	if _, err = c.ReadToken(); err != nil {
		return
	}

	for c.PeekKind() != '}' {
		// key name
		tok, err := c.ReadToken()
		if err != nil {
			return nil, nil, err
		}
		keys = append(keys, tok.String())
		// value type
		t, err := c.Type()
		if err != nil {
			return nil, nil, err
		}
		types = append(types, t)
	}

	_, err = c.ReadToken()
	return
}

func validNames(items []string) bool {
	for _, item := range items {
		if item == "" {
			return false
		}

		if !(('a' <= item[0] && item[0] <= 'z') || ('A' <= item[0] && item[0] <= 'Z') || item[0] == '_') {
			return false
		}

		for i := 1; i < len(item); i++ {
			if !(('a' <= item[i] && item[i] <= 'z') || ('A' <= item[i] && item[i] <= 'Z') || ('0' <= item[i] && item[i] <= '9') || item[i] == '_') {
				return false
			}
		}
	}

	return true
}

func validIntegers(items []string) bool {
	for _, s := range items {
		if s == "" {
			return false
		}
		if !(s[0] == '+' || s[0] == '-' || ('0' <= s[0] && s[0] <= '9')) {
			return false
		}
		for i := 1; i < len(s); i++ {
			if !('0' <= s[i] && s[i] <= '9') {
				return false
			}
		}
	}
	return true
}

func (c *Context) objectType() (Type, error) {
	keys, types, err := c.object()
	if err != nil {
		return nil, err
	}

	if len(keys) == 0 {
		return Map{
			Key:   String{},
			Value: Any{},
		}, nil
	}

	if !validNames(keys) {
		m := Map{Value: c.deduce(types)}
		if validIntegers(keys) {
			m.Key = Int{}
		} else {
			m.Key = String{}
		}
		return m, nil
	}

	return Struct{
		Fields: slices.Collect(func(yield func(field *Field) bool) {
			for i, key := range keys {
				if !yield(&Field{
					Key:  key,
					Type: types[i],
				}) {
					return
				}
			}
		}),
		CamelCaseConverter: c.CamelCaseConverter,
	}, nil
}

func (c *Context) array() (types []Type, err error) {
	if c.PeekKind() != '[' {
		err = fmt.Errorf("expected '['")
		return
	}
	if _, err = c.ReadToken(); err != nil {
		return
	}

	for c.PeekKind() != ']' {
		t, err := c.Type()
		if err != nil {
			return nil, err
		}
		types = append(types, t)
	}

	_, err = c.ReadToken()
	return
}

func (c *Context) arrayType() (Type, error) {
	if types, err := c.array(); err != nil {
		return nil, err
	} else {
		return Array{Element: c.deduce(types)}, nil
	}
}

func (c *Context) Type() (Type, error) {
	switch c.PeekKind() {
	case '{':
		return c.objectType()
	case '[':
		return c.arrayType()
	case 't', 'f':
		return Bool{}, c.SkipValue()
	case '0':
		tok, err := c.ReadToken()
		if err != nil {
			return nil, err
		}

		if strings.ContainsAny(tok.String(), ".eE") {
			return Float{}, nil
		} else {
			return Int{}, nil
		}
	case '"':
		return String{}, c.SkipValue()
	case 'n':
		return Any{}, c.SkipValue()
	default:
		tok, err := c.ReadToken()
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("unexpected %v: %v", tok.Kind(), tok.String())
	}
}

func is[T Type](t Type) bool {
	_, ok := t.(T)
	return ok
}

func all[T Type](types []Type) bool {
	for _, t := range types {
		if _, ok := t.(T); !ok {
			return false
		}
	}
	return true
}

func remove[T Type](types []Type) []Type {
	i := 0
	for j := 0; j < len(types); j++ {
		if !is[T](types[j]) {
			if i != j {
				types[i] = types[j]
			}
			i++
		}
	}
	return types[:i]
}

func (c *Context) deduceStruct(types []Type) Struct {
	m := make(map[string][]*Field)
	var keys []string

	for _, t := range types {
		s := t.(Struct)
		for _, field := range s.Fields {
			if _, ok := m[field.Key]; !ok {
				m[field.Key] = make([]*Field, 0, len(types))
				keys = append(keys, field.Key)
			}
			m[field.Key] = append(m[field.Key], field)
		}
	}

	s := Struct{
		Fields:             make([]*Field, 0, len(keys)),
		CamelCaseConverter: c.CamelCaseConverter,
	}
	for _, key := range keys {
		fields := m[key]
		field := fields[0]

		field.Type = c.deduce(slices.Collect(
			xiter.Map(
				func(field *Field) Type {
					return field.Type
				},
				slices.Values(fields),
			),
		))
		field.OmitEmpty = len(fields) != len(types)
		if field.OmitEmpty {
			field.Type = field.Type.Nullable()
		}

		s.Fields = append(s.Fields, field)
	}

	return s
}

func (c *Context) deduceMap(types []Type) Type {
	return Map{
		Key: c.deduce(lo.Map(types, func(item Type, _ int) Type {
			return item.(Map).Key
		})),
		Value: c.deduce(lo.Map(types, func(item Type, _ int) Type {
			return item.(Map).Value
		})),
	}
}

func (c *Context) deduce(types []Type) Type {
	n := len(types)
	types = remove[Any](types)
	nullable := len(types) < n

	switch len(types) {
	case 0:
		return Any{}
	case 1:
		if nullable {
			return types[0].Nullable()
		} else {
			return types[0]
		}
	}

	switch types[0].(type) {
	case String:
		if all[String](types) {
			return String{Pointer: nullable}
		}
	case Int:
		if all[Int](types) {
			return Int{Pointer: nullable}
		} else {
			ok := lo.EveryBy(types, func(item Type) bool {
				return is[Int](item) || is[Float](item)
			})
			if ok {
				return Float{Pointer: nullable}
			}
		}
	case Float:
		if all[Float](types) {
			return Float{Pointer: nullable}
		}
	case Bool:
		if all[Bool](types) {
			return Bool{Pointer: nullable}
		}
	case Array:
		if all[Array](types) {
			for i := range types {
				types[i] = types[i].(Array).Element
			}
			return Array{
				Element: c.deduce(types),
			}
		}
	case Struct:
		if all[Struct](types) {
			if nullable {
				return c.deduceStruct(types).Nullable()
			} else {
				return c.deduceStruct(types)
			}
		} else {
			ok := lo.EveryBy(types, func(item Type) bool {
				return is[Struct](item) || is[Map](item)
			})
			if ok {
				for i := range types {
					if s, ok := types[i].(Struct); ok {
						types[i] = Map{
							Key: String{},
							Value: c.deduce(lo.Map(s.Fields, func(field *Field, _ int) Type {
								return field.Type
							})),
						}
					}
				}
				return c.deduceMap(types)
			}
		}
	case Map:
		if all[Map](types) {
			return c.deduceMap(types)
		}
	}

	return Any{}
}
