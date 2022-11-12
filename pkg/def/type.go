package def

import (
	"encoding/json"
	gen "github.com/dave/jennifer/jen"
	"github.com/fhluo/json2go/internal/def"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

type Type interface {
	Code() gen.Code
}

type Nullable interface {
	Type
	Nullable() Type
}

type TypeDecl struct {
	Name string
	Type Type
}

func (t TypeDecl) Code() gen.Code {
	return gen.Type().Id(t.Name).Add(t.Type.Code())
}

type Int struct {
	Pointer bool
}

func (i Int) Nullable() Type {
	i.Pointer = true
	return i
}

func (i Int) Code() gen.Code {
	if i.Pointer {
		return gen.Op("*").Int()
	}
	return gen.Int()
}

type Float struct {
	Pointer bool
}

func (f Float) Nullable() Type {
	f.Pointer = true
	return f
}

func (f Float) Code() gen.Code {
	if f.Pointer {
		return gen.Op("*").Float64()
	}
	return gen.Float64()
}

type Bool struct {
	Pointer bool
}

func (b Bool) Code() gen.Code {
	if b.Pointer {
		return gen.Op("*").Bool()
	}
	return gen.Bool()
}

type String struct {
	Pointer bool
}

func (s String) Nullable() Type {
	s.Pointer = true
	return s
}

func (s String) Code() gen.Code {
	if s.Pointer {
		return gen.Op("*").String()
	}
	return gen.String()
}

type Array struct {
	Element Type // element type
}

func (a Array) Code() gen.Code {
	return gen.Index().Add(a.Element.Code())
}

type Field struct {
	Name      string // field name
	Key       string // json key
	Type      Type
	OmitEmpty bool
	String    bool // string option: string, bool, int or float
}

func (f Field) Code() gen.Code {
	options := make([]string, 0, 3)
	options = append(options, f.Key)

	if f.OmitEmpty {
		options = append(options, "omitempty")
	}
	if f.String {
		options = append(options, "string")
	}

	return gen.Id(f.Name).Add(f.Type.Code(), gen.Tag(map[string]string{"json": strings.Join(options, ",")}))
}

type Struct struct {
	Fields  []*Field
	Pointer bool

	named bool
}

func (s Struct) Map() Map {
	return Map{
		Key: String{},
		Value: deduce(lo.Map(s.Fields, func(field *Field, _ int) Type {
			return field.Type
		})),
	}
}

func (s Struct) Nullable() Type {
	s.Pointer = true
	return s
}

func (s Struct) Naming() {
	if s.named {
		return
	}

	counter := make(map[string]int)
	for _, field := range s.Fields {
		field.Name = def.ToCamelCase(field.Key)
		if counter[field.Name]++; counter[field.Name] != 1 {
			field.Name += strconv.Itoa(counter[field.Name])
		}
	}

	s.named = true
}

func (s Struct) Code() gen.Code {
	s.Naming()
	codes := lo.Map(s.Fields, func(field *Field, _ int) gen.Code {
		return field.Code()
	})

	if s.Pointer {
		gen.Op("*").Struct(codes...)
	}
	return gen.Struct(codes...)
}

type Map struct {
	Key   Type // key type: string or int
	Value Type // value type
}

func (m Map) Code() gen.Code {
	return gen.Map(m.Key.Code()).Add(m.Value.Code())
}

type Any struct{}

func (a Any) Code() gen.Code {
	return gen.Any()
}

func is[T Type](t Type) bool {
	_, ok := t.(T)
	return ok
}

func isInteger(number json.Number) bool {
	_, err := number.Int64()
	return err == nil
}

func all[T Type](types []Type) bool {
	return lo.EveryBy(types, func(item Type) bool {
		return is[T](item)
	})
}

func deduce(types []Type) Type {
	switch len(types) {
	case 0:
		return Any{}
	case 1:
		return types[0]
	default:

	}

	var nullable bool

	types = lo.Filter(types, func(item Type, _ int) bool {
		if is[Any](item) {
			nullable = true
			return false
		} else {
			return true
		}
	})

	switch len(types) {
	case 0:
		return Any{}
	case 1:
		if t, ok := types[0].(Nullable); ok && nullable {
			return t.Nullable()
		} else {
			return types[0]
		}
	}

	switch types[0].(type) {
	case Int:
		if all[Int](types) {
			return Int{Pointer: nullable}
		} else {
			ok := lo.EveryBy(types, func(item Type) bool {
				return is[Int](item) || is[Float](item)
			})
			if ok {
				return Float{Pointer: nullable}
			} else {
				return Any{}
			}
		}
	case Float:
		if all[Float](types) {
			return Float{Pointer: nullable}
		} else {
			return Any{}
		}
	case Bool:
		if all[Bool](types) {
			return Bool{Pointer: nullable}
		} else {
			return Any{}
		}
	case Array:
		if all[Array](types) {
			return Array{
				Element: deduce(lo.Map(types, func(item Type, _ int) Type {
					return item.(Array).Element
				})),
			}
		} else {
			return Any{}
		}
	case Struct:
		if all[Struct](types) {
			m := make(map[string][]*Field)
			for _, t := range types {
				s := t.(Struct)
				s.Naming()
				for _, field := range s.Fields {
					m[field.Name] = append(m[field.Name], field)
				}
			}

			s := Struct{Pointer: nullable}
			for _, v := range m {
				s.Fields = append(s.Fields, &Field{
					Name: v[0].Name,
					Key:  v[0].Key,
					Type: deduce(lo.Map(v, func(field *Field, _ int) Type {
						return field.Type
					})),
					OmitEmpty: len(v) != len(types),
				})
			}

			return s
		} else {
			ok := lo.EveryBy(types, func(item Type) bool {
				return is[Struct](item) || is[Map](item)
			})
			if ok {
				types = lo.Map(types, func(item Type, _ int) Type {
					if s, ok := item.(Struct); ok {
						return s.Map()
					}
					return item
				})
				return Map{
					Key: deduce(lo.Map(types, func(item Type, _ int) Type {
						return item.(Map).Key
					})),
					Value: deduce(lo.Map(types, func(item Type, _ int) Type {
						return item.(Map).Value
					})),
				}
			} else {
				return Any{}
			}
		}
	case Map:
		if all[Map](types) {
			return Map{
				Key: deduce(lo.Map(types, func(item Type, _ int) Type {
					return item.(Map).Key
				})),
				Value: deduce(lo.Map(types, func(item Type, _ int) Type {
					return item.(Map).Value
				})),
			}
		} else {
			return Any{}
		}
	}

	return nil
}
