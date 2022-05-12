package json2go

import (
	gen "github.com/dave/jennifer/jen"
	"github.com/tidwall/gjson"
	"golang.org/x/exp/maps"
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
	if !validNames(keys) {
		code := UnnamedTypeFromArray(maps.Values(json.Map()))
		if validInts(keys) {
			return gen.Map(gen.Int()).Add(code)
		} else {
			return gen.Map(gen.String()).Add(code)
		}
	}

	return gen.StructFunc(func(group *gen.Group) {
		convert := fieldNameConverter()

		json.ForEach(func(key, value gjson.Result) bool {
			var code gen.Code

			switch t := JSONType(value); t {
			case Object:
				code = UnnamedStructOrMap(value)
			case Array:
				code = UnnamedSlice(value)
			default:
				code = t.Code()
			}

			name := key.String()
			group.Id(convert(name)).Add(code).Tag(map[string]string{"json": name})

			return true
		})
	})
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

	if !validNames(names) {
		code := UnnamedTypeFromArray(reduce(maps.Values(fields)))
		if validInts(names) {
			return gen.Map(gen.Int()).Add(code)
		} else {
			return gen.Map(gen.String()).Add(code)
		}
	}

	return gen.StructFunc(func(group *gen.Group) {
		convert := fieldNameConverter()

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

			tag := make(map[string]string)
			if omitempty {
				tag["json"] = name + ",omitempty"
			} else {
				tag["json"] = name
			}

			group.Id(convert(name)).Add(code).Tag(tag)
		}
	})
}
