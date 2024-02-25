package tests

import (
	"testing"

	argsparser "github.com/aadi-1024/configo/args_parser"
)

func TestArgsParser(t *testing.T) {
	a := argsparser.ArgsParser{
		Prefix: '-',
	}
	m, _ := a.GenerateConfig([]string{"-h", "hello world", "woah", "-a"})
	t.Log(m)
}