package def

import (
	"fmt"
	gen "github.com/dave/jennifer/jen"
	"github.com/fhluo/json2go/pkg/scanner"
	"github.com/fhluo/json2go/pkg/token"
	"github.com/samber/lo"
)

type Context struct {
	*scanner.Scanner
}

func From(s string) *Context {
	c := new(Context)
	c.Scanner = scanner.New(s)
	return c
}

func FromBytes(data []byte) *Context {
	c := new(Context)
	c.Scanner = scanner.New(string(data))
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
	var (
		key string
		t   Type
	)
	for c.More() {
		// key
		_, key, err = c.Scan()
		if err != nil {
			return
		}
		keys = append(keys, key)
		// value
		if t, err = c.Type(); err != nil {
			return
		} else {
			types = append(types, t)
		}
	}

	_, _, err = c.Scan()
	return
}

func (c *Context) objectType() (Type, error) {
	keys, types, err := c.object()
	if err != nil {
		return nil, err
	}

	if !ValidNames(keys) {
		m := Map{Value: deduce(types)}
		if ValidIntegers(keys) {
			m.Key = Int{}
		} else {
			m.Key = String{}
		}
		return m, nil
	}

	return Struct{
		Fields: lo.Map(keys, func(key string, i int) *Field {
			return &Field{
				Key:  key,
				Type: types[i],
			}
		}),
	}, nil
}

func (c *Context) array() (types []Type, err error) {
	var t Type
	for c.More() {
		if t, err = c.Type(); err != nil {
			return
		} else {
			types = append(types, t)
		}
	}

	_, _, err = c.Scan()
	return
}

func (c *Context) arrayType() (Type, error) {
	if types, err := c.array(); err != nil {
		return nil, err
	} else {
		return Array{Element: deduce(types)}, nil
	}
}

func (c *Context) Type() (Type, error) {
	t, _, err := c.Scan()
	if err != nil {
		return nil, err
	}

	switch t {
	case token.LeftBrace:
		return c.objectType()
	case token.LeftBracket:
		return c.arrayType()
	case token.Bool:
		return Bool{}, nil
	case token.Int:
		return Int{}, nil
	case token.Float:
		return Float{}, nil
	case token.String:
		return String{}, nil
	case token.Null:
		return Any{}, nil
	default:
		return nil, fmt.Errorf("unexpected token %s", t)
	}
}
