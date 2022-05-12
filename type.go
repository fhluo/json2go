package json2go

import (
	gen "github.com/dave/jennifer/jen"
	"github.com/tidwall/gjson"
	"golang.org/x/exp/maps"
	"strconv"
)

type T int

const (
	Null T = iota
	Bool
	Int
	Float
	String
	Object
	Array
	Any
)

func (t T) Code() gen.Code {
	switch t {
	case Bool:
		return gen.Bool()
	case Int:
		return gen.Int()
	case Float:
		return gen.Float64()
	case String:
		return gen.String()
	default: // Null, Object , Array, Any
		return gen.Any()
	}
}

func JSONType(json gjson.Result) T {
	switch json.Type {
	case gjson.Null:
		return Null
	case gjson.False, gjson.True:
		return Bool
	case gjson.Number:
		if _, err := strconv.Atoi(json.Raw); err == nil {
			return Int
		} else {
			return Float
		}
	case gjson.String:
		return String
	case gjson.JSON:
		switch {
		case json.IsObject():
			return Object
		case json.IsArray():
			return Array
		default:
			return Any
		}
	default:
		return Any
	}
}

func countArrayTypes(results []gjson.Result) map[T]int {
	counter := make(map[T]int)

	for _, result := range results {
		switch t := JSONType(result); t {
		case Null:
			counter[Null]++
		case Int:
			if counter[Float] == 0 {
				counter[Int]++
			}
		case Float:
			if counter[Int] > 0 {
				counter[Float] += counter[Int]
				delete(counter, Int)
			}
			counter[Float]++
		default:
			counter[t]++
		}
	}

	return counter
}

func JSONArrayType(array []gjson.Result) (T, bool) {
	counter := countArrayTypes(array)

	if counter[Any] > 0 || counter[Null] == len(array) {
		return Any, true
	}

	switch {
	case len(counter) == 1:
		return maps.Keys(counter)[0], false
	case len(counter) == 2 && counter[Null] > 0:
		delete(counter, Null)
		return maps.Keys(counter)[0], true
	default:
		return Any, true
	}
}
