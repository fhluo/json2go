package def

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/samber/lo"
)

type Context struct {
	*json.Decoder
	token json.Token
	err   error
}

func From(s string) *Context {
	c := new(Context)
	c.Decoder = json.NewDecoder(bytes.NewBufferString(s))
	c.UseNumber()
	return c
}

func FromBytes(data []byte) *Context {
	c := new(Context)
	c.Decoder = json.NewDecoder(bytes.NewBuffer(data))
	c.UseNumber()
	return c
}

func (c *Context) Error() error {
	return c.err
}

func (c *Context) Token() json.Token {
	return c.token
}

func (c *Context) Next() json.Token {
	c.token, c.err = c.Decoder.Token()
	return c.token
}

func (c *Context) TypeDecl(name string) Type {
	return TypeDecl{
		Name: name,
		Type: c.Type(),
	}
}

func (c *Context) Type() Type {
	if c.err != nil {
		return nil
	}

	switch x := c.Next().(type) {
	case json.Delim:
		switch x {
		case '{':
			defer func() {
				if c.err == nil && !c.More() {
					c.Next()
				}
			}()

			var keys []string
			var types []Type // value types

			for c.More() {
				// key
				switch key := c.Next().(type) {
				case string:
					keys = append(keys, key)
				default:
					c.err = fmt.Errorf("unexpected type")
					return nil
				}
				// value
				types = append(types, c.Type())
			}

			if !ValidNames(keys) {
				m := Map{Value: deduce(types)}
				if ValidIntegers(keys) {
					m.Key = Int{}
				} else {
					m.Key = String{}
				}
				return m
			}

			return Struct{
				Fields: lo.Map(keys, func(key string, i int) *Field {
					return &Field{
						Key:  key,
						Type: types[i],
					}
				}),
			}
		case '[':
			defer func() {
				if c.err == nil && !c.More() {
					c.Next()
				}
			}()

			var types []Type
			for c.More() {
				types = append(types, c.Type())
			}

			return Array{Element: deduce(types)}
		default:
			return nil
		}
	case bool:
		return Bool{}
	case json.Number:
		if isInteger(x) {
			return Int{}
		} else {
			return Float{}
		}
	case string:
		return String{}
	case nil:
		if c.err == nil {
			return Any{}
		} else {
			return nil
		}
	default:
		c.err = fmt.Errorf("unexpected type")
		return nil
	}
}
