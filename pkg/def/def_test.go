package def

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFrom(t *testing.T) {
	stmt, err := From("{}").Declare("T")
	if err != nil {
		t.Fatal(err)
	}

	buffer := new(bytes.Buffer)
	if err = stmt.Render(buffer); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "type T map[string]any", buffer.String())
}
