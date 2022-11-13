package def

import (
	"fmt"
	gen "github.com/dave/jennifer/jen"
	"github.com/samber/lo"
	"go/format"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

type Type interface {
	fmt.Stringer
	Code() gen.Code
	Nullable() Type
}

type TypeDecl struct {
	Name string
	Type Type
}

func (t TypeDecl) String() string {
	s := fmt.Sprintf("type %s %s", t.Name, t.Type)
	r, err := format.Source([]byte(s))
	if err != nil {
		return s
	} else {
		return string(r)
	}
}

type Int struct {
	Pointer bool
}

func (i Int) String() string {
	if i.Pointer {
		return "*int"
	} else {
		return "int"
	}
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

func (f Float) String() string {
	if f.Pointer {
		return "*float64"
	} else {
		return "float64"
	}
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

func (b Bool) String() string {
	if b.Pointer {
		return "*bool"
	} else {
		return "bool"
	}
}

func (b Bool) Nullable() Type {
	b.Pointer = true
	return b
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

func (s String) String() string {
	if s.Pointer {
		return "*string"
	} else {
		return "string"
	}
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

func (a Array) String() string {
	return fmt.Sprintf("[]%s", a.Element)
}

func (a Array) Nullable() Type {
	return a
}

func (a Array) Code() gen.Code {
	return gen.Index().Add(a.Element.Code())
}

type Field struct {
	Name      string // field name
	Key       string // json key
	Type      Type
	OmitEmpty bool
	//String    bool // string option: string, bool, int or float
}

func (f Field) String() string {
	return fmt.Sprintf("%s %s `json:\"%s\"`", f.Name, f.Type, f.Options())
}

func (f Field) Options() string {
	options := make([]string, 0, 3)
	options = append(options, f.Key)

	if f.OmitEmpty {
		options = append(options, "omitempty")
	}
	//if f.String {
	//	options = append(options, "string")
	//}
	return strings.Join(options, ",")
}

func (f Field) Code() gen.Code {
	return gen.Id(f.Name).Add(f.Type.Code(), gen.Tag(map[string]string{"json": f.Options()}))
}

type Struct struct {
	Fields  []*Field
	Pointer bool
}

func (s Struct) String() string {
	s.Naming()
	return fmt.Sprintf(
		`struct{
	%s
}`,
		strings.Join(
			lo.Map(s.Fields, func(field *Field, _ int) string {
				return field.String()
			}),
			"\n\t",
		),
	)
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
	if len(s.Fields) == 0 {
		return
	}

	named := !slices.Contains(
		lo.Map(s.Fields, func(field *Field, _ int) string {
			return field.Name
		}),
		"",
	)
	if named {
		return
	}

	if s.Fields[0].Name != "" {
		return
	}

	counter := make(map[string]int)
	for _, field := range s.Fields {
		field.Name = ToCamelCase(field.Key)
		if counter[field.Name]++; counter[field.Name] != 1 {
			field.Name += strconv.Itoa(counter[field.Name])
		}
	}
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

func (m Map) String() string {
	return fmt.Sprintf("map[%s]%s", m.Key, m.Value)
}

func (m Map) Nullable() Type {
	return m
}

func (m Map) Code() gen.Code {
	return gen.Map(m.Key.Code()).Add(m.Value.Code())
}

type Any struct{}

func (a Any) String() string {
	return "any"
}

func (a Any) Nullable() Type {
	return a
}

func (a Any) Code() gen.Code {
	return gen.Any()
}

func is[T Type](t Type) bool {
	_, ok := t.(T)
	return ok
}

func isInteger(s string) bool {
	if s == "" {
		return false
	}
	if !(s[0] == '+' || s[0] == '-' || ('0' <= s[0] && s[0] <= '9')) {
		return false
	}
	for i := 1; i < len(s); i++ {
		if !('0' <= s[0] && s[0] <= '9') {
			return false
		}
	}
	return true
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

func deduceStruct(types []Type) Struct {
	m := make(map[string][]*Field)
	var names []string

	for _, t := range types {
		s := t.(Struct)
		s.Naming()
		for _, field := range s.Fields {
			if _, ok := m[field.Name]; !ok {
				m[field.Name] = make([]*Field, 0, len(types))
				names = append(names, field.Name)
			}
			m[field.Name] = append(m[field.Name], field)
		}
	}

	s := Struct{
		Fields: make([]*Field, 0, len(names)),
	}
	for _, name := range names {
		fields := m[name]
		field := fields[0]

		field.Type = deduce(lo.Map(fields, func(field *Field, _ int) Type {
			return field.Type
		}))
		field.OmitEmpty = len(fields) != len(types)
		if field.OmitEmpty {
			field.Type = field.Type.Nullable()
		}

		s.Fields = append(s.Fields, field)
	}

	return s
}

func deduceMap(types []Type) Type {
	return Map{
		Key: deduce(lo.Map(types, func(item Type, _ int) Type {
			return item.(Map).Key
		})),
		Value: deduce(lo.Map(types, func(item Type, _ int) Type {
			return item.(Map).Value
		})),
	}
}

func deduce(types []Type) Type {
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
				Element: deduce(types),
			}
		}
	case Struct:
		if all[Struct](types) {
			if nullable {
				return deduceStruct(types).Nullable()
			} else {
				return deduceStruct(types)
			}
		} else {
			ok := lo.EveryBy(types, func(item Type) bool {
				return is[Struct](item) || is[Map](item)
			})
			if ok {
				for i := range types {
					if s, ok := types[i].(Struct); ok {
						types[i] = s.Map()
					}
				}
				return deduceMap(types)
			}
		}
	case Map:
		if all[Map](types) {
			return deduceMap(types)
		}
	}

	return Any{}
}
