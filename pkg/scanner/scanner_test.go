package scanner

import (
	"github.com/fhluo/json2go/pkg/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

func scan(s string) (tokens []token.Token) {
	scanner := New(s)

	t, _, err := scanner.Scan()
	if err != nil {
		return
	}

	for t != token.EOF {
		tokens = append(tokens, t)
		t, _, err = scanner.Scan()
		if err != nil {
			return
		}
	}

	return
}

func TestScanner_Scan(t *testing.T) {
	json := `[
  {
    "string": "中文",
    "int": 123,
    "float": 1.0,
    "bool": false,
    "null": null,
    "array": [
      "\\\"", -1, 0, 1, 1.0, 1e3, 1e-3, true, false, null
    ]
  }
]`
	assert.Equal(
		t, scan(json),
		[]token.Token{
			token.LeftBracket,
			token.LeftBrace,
			token.String, token.String,
			token.String, token.Int,
			token.String, token.Float,
			token.String, token.Bool,
			token.String, token.Null,
			token.String, token.LeftBracket,
			token.String, token.Int, token.Int, token.Int, token.Float, token.Float, token.Float, token.Bool, token.Bool, token.Null,
			token.RightBracket,
			token.RightBrace,
			token.RightBracket,
		},
	)
}
