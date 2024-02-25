package jsonparser

import (
	"encoding/json"
	"os"
)

type JsonParser struct {

}

func (j *JsonParser) GenerateConfig(path string) (map[string]any, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	f, _ := file.Stat()
	b := make([]byte, f.Size())

	_, err = file.Read(b)
	if err != nil {
		return nil, err
	}

	m := make(map[string]any)
	err = json.Unmarshal(b, &m)

	return m, err
}