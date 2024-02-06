package yq

import (
	"testing"
)

func TestParseYAML(t *testing.T) {
	filename := "quotes.yaml"
	r, err := ParseYAML(filename)
	if err != nil {
		t.Errorf("error parsing yaml with err: %v", err)
	}
	PrintNodeTree(&r.Node, 0)
}
