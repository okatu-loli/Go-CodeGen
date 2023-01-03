package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"GO-CodeGen/model"
	"GO-CodeGen/utils"
)

func main() {
	flags, err := utils.ParseFlags()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if _, err := os.Stat(flags.OutputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(flags.OutputDir, 0755); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	jsonData, err1 := ioutil.ReadFile(flags.JSONFile)
	if err1 != nil {
		fmt.Fprintln(os.Stderr, err1)
		os.Exit(1)
	}
	var data model.Data
	if err2 := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Fprintln(os.Stderr, err2)
		os.Exit(1)
	}
	structs := make([]model.Struct, 0)
	for _, field := range data.Fields {
		structs = append(structs, model.Struct{Name: field.Name, Fields: []model.Field{{Name: field.Name, Type: field.Type, Description: field.Description}}})
	}
	templateData := &utils.TemplateData{
		Structs:     structs,
		Name:        flags.Name,
		PackageName: flags.PackageName,
		Fields:      data.Fields,
	}
	if err3 := templateData.Generate("output.go", flags.OutputDir, flags.PackageName); err != nil {
		fmt.Fprintln(os.Stderr, err3)
		os.Exit(1)
	}
}
