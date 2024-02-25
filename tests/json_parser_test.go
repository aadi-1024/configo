package tests

import (
	"testing"

	jsonparser "github.com/aadi-1024/configo/json_parser"
	"github.com/google/go-cmp/cmp"
)

func TestJsonParse(t *testing.T) {
	p := jsonparser.JsonParser{}
	m, e := p.GenerateConfig("test.json")

	if e != nil {
		t.Errorf(e.Error())
	}

	x := make(map[string]any)
	x["hi"] = 2.0
	x["hello"] = "hey there"

	if !cmp.Equal(x, m) {
		t.Errorf("values didnt match: %v and %v", x, m)
	}
	t.Log(m)
}