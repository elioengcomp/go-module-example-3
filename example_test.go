package example

import (
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	result := Exec()
	if !strings.Contains(result, "This is go-module-example-3 "+version) {
		t.Errorf("Failed to test example package")
	}
}
