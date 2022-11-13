package def

import (
	"bytes"
	"fmt"
	gen "github.com/dave/jennifer/jen"
	"github.com/goccy/go-json"
	"github.com/samber/lo"
)

type Context struct {
	*json.Decoder
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

func (c *Context) Declare(name string) (*gen.Statement, error) {
	t, err := c.Type()
	if err != nil {
		return nil, err
	}
	return gen.Type().Id(name).Add(t.Code()), nil
}

func (c *Context) object() (keys []string, types []Type, err error) {
	for c.More() {
		// key
		token, err := c.Token()
		if err != nil {
			return keys, types, err
		}

		if key, ok := token.(string); !ok {
			return keys, types, fmt.Errorf("unexpected type")
		} else {
			keys = append(keys, key)
		}
		// value
		if t, err := c.Type(); err != nil {
			return keys, types, err
		} else {
			types = append(types, t)
		}
	}

	_, err = c.Token()
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
	for c.More() {
		if t, err := c.Type(); err != nil {
			return types, err
		} else {
			types = append(types, t)
		}
	}

	_, err = c.Token()
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
	token, err := c.Token()
	if err != nil {
		return nil, err
	}

	switch x := token.(type) {
	case json.Delim:
		switch x {
		case '{':
			return c.objectType()
		case '[':
			return c.arrayType()
		default:
			return nil, fmt.Errorf("invalid delim %s", x)
		}
	case bool:
		return Bool{}, nil
	case json.Number:
		if isInteger(x.String()) {
			return Int{}, nil
		} else {
			return Float{}, nil
		}
	case string:
		return String{}, nil
	case nil:
		return Any{}, nil
	default:
		return nil, fmt.Errorf("unexpected type")
	}
}
