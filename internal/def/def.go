package def

import (
	gen "github.com/dave/jennifer/jen"
	"github.com/tidwall/gjson"
	"golang.org/x/exp/maps"
	"strconv"
)

func Type(name string, json gjson.Result) gen.Code {
	return gen.Type().Id(name).Add(UnnamedType(json))
}

func UnnamedType(json gjson.Result) gen.Code {
	switch t := JSONType(json); t {
	case Object:
		return UnnamedStructOrMap(json)
	case Array:
		return UnnamedSlice(json)
	default:
		return t.Code()
	}
}

func UnnamedTypeFromArray(array []gjson.Result) gen.Code {
	t, null := JSONArrayType(array)

	switch t {
	case Any:
		return t.Code()
	case Array:
		return gen.Index().Add(UnnamedTypeFromArray(reduceArrayResults(array)))
	case Object:
		if null {
			return gen.Op("*").Add(UnnamedStructOrMapFromArray(array))
		} else {
			return UnnamedStructOrMapFromArray(array)
		}
	default:
		if null {
			return gen.Op("*").Add(t.Code())
		} else {
			return t.Code()
		}
	}
}

func UnnamedStructOrMap(json gjson.Result) gen.Code {
	keys := maps.Keys(json.Map())
	if !ValidNames(keys) {
		code := UnnamedTypeFromArray(maps.Values(json.Map()))
		if ValidIntegers(keys) {
			return gen.Map(gen.Int()).Add(code)
		} else {
			return gen.Map(gen.String()).Add(code)
		}
	}

	s := NewUnnamedStruct()

	json.ForEach(func(key, value gjson.Result) bool {
		s.Add(key.String(), UnnamedType(value), false)
		return true
	})

	return s.Code()
}

func UnnamedSlice(json gjson.Result) gen.Code {
	return gen.Index().Add(UnnamedTypeFromArray(json.Array()))
}

func reduceArrayResults(arrays []gjson.Result) []gjson.Result {
	var result []gjson.Result
	for _, array := range arrays {
		result = append(result, array.Array()...)
	}
	return result
}

func reduce(arrays [][]gjson.Result) []gjson.Result {
	var result []gjson.Result
	for _, array := range arrays {
		result = append(result, array...)
	}
	return result
}

func UnnamedStructOrMapFromArray(array []gjson.Result) gen.Code {
	var names []string

	fields := make(map[string][]gjson.Result)

	for _, item := range array {
		item.ForEach(func(key, value gjson.Result) bool {
			name := key.String()

			if _, ok := fields[name]; !ok {
				names = append(names, name)
			}

			fields[name] = append(fields[name], value)

			return true
		})
	}

	if !ValidNames(names) {
		code := UnnamedTypeFromArray(reduce(maps.Values(fields)))
		if ValidIntegers(names) {
			return gen.Map(gen.Int()).Add(code)
		} else {
			return gen.Map(gen.String()).Add(code)
		}
	}

	s := NewUnnamedStruct()

	for _, name := range names {
		fieldValues := fields[name]

		t, null := JSONArrayType(fieldValues)
		omitempty := len(fieldValues) != len(array)

		var code gen.Code

		switch t {
		case Any:
			code = t.Code()
		case Array:
			code = gen.Index().Add(UnnamedTypeFromArray(reduceArrayResults(fieldValues)))
		case Object:
			if null || omitempty {
				code = gen.Op("*").Add(UnnamedStructOrMapFromArray(fieldValues))
			} else {
				code = UnnamedStructOrMapFromArray(fieldValues)
			}
		default:
			if null || omitempty {
				code = gen.Op("*").Add(t.Code())
			} else {
				code = t.Code()
			}
		}

		s.Add(name, code, omitempty)
	}

	return s.Code()
}

func JSONTag(s string) gen.Code {
	return gen.Tag(map[string]string{"json": s})
}

type UnnamedStruct struct {
	Fields []gen.Code

	counter map[string]int
}

func NewUnnamedStruct() *UnnamedStruct {
	return &UnnamedStruct{counter: make(map[string]int)}
}

func (s *UnnamedStruct) naming(key string) string {
	name := ToCamelCase(key)
	if s.counter[name]++; s.counter[name] != 1 {
		name += strconv.Itoa(s.counter[name])
	}
	return name
}

func (s *UnnamedStruct) Add(key string, type_ gen.Code, omitempty bool) {
	var tag gen.Code
	if omitempty {
		tag = JSONTag(key + ",omitempty")
	} else {
		tag = JSONTag(key)
	}

	s.Fields = append(s.Fields, gen.Id(s.naming(key)).Add(type_, tag))
}

func (s *UnnamedStruct) Code() gen.Code {
	return gen.Struct(s.Fields...)
}
