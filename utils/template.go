package utils

import (
	"GO-CodeGen/model"
	"os"
	"path/filepath"
	"text/template"
)

// TemplateData 是传递给 template 的数据
type TemplateData struct {
	Structs     []model.Struct
	Funcs       []model.Func
	Methods     []model.Method
	Name        string
	PackageName string
	Fields      []model.Field
}

// 定义一个接口，包含一个方法
type Generator interface {
	Generate(outputFile string, outputDir string, packageName string) error
}

// 实现 Generator 接口的 Generate 方法
func (g *TemplateData) Generate(outputFile string, outputDir string, packageName string) error {
	tmpl, err := template.ParseFiles("templates/template.tmpl")
	if err != nil {
		return err
	}
	// 将输出文件的名称和输出文件夹拼接在一起
	outputPath := filepath.Join(outputDir, outputFile)
	file, creatErr := os.Create(outputPath)
	if creatErr != nil {
		return creatErr
	}
	defer file.Close()
	return tmpl.ExecuteTemplate(file, "template", g)
}
