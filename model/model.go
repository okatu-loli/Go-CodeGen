package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Data struct {
	Fields  []Field  `json:"fields"`
	Structs []Struct `json:"structs"`
	Funcs   []Func   `json:"funcs"`
	Methods []Method `json:"methods"`
}

type Field struct {
	Name        string
	Type        string
	Description string // Description is the description of the field
}

type Struct struct {
	Name        string
	Type        string
	Description string
	Fields      []Field
}

type Func struct {
	Name string
}

type Method struct {
	Name     string
	Receiver string
}

// LoadData loads data from a JSON file
func LoadData(file string) (*Data, error) {
	jsonData, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}
	var data Data
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}
	return &data, nil
}
