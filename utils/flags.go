package utils

import (
	"flag"
	"fmt"
	"os"
)

// Flags holds the command-line flags
type Flags struct {
	OutputDir   string
	Name        string
	PackageName string
	JSONFile    string
}

func ParseFlags() (*Flags, error) {
	// 定义命令行参数
	outputDir := flag.String("dir", "output", "代码输出目录")
	name := flag.String("name", "", "实例名称")
	packageName := flag.String("pkg", "", "包名")
	jsonFile := flag.String("json", "./json/default.json", "Json文件")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "帮助：\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *name == "" {
		fmt.Fprintln(os.Stderr, "Error: name参数必须填入")
		os.Exit(1)
	}
	if *packageName == "" {
		*packageName = *name
	}
	return &Flags{
		OutputDir:   *outputDir,
		Name:        *name,
		PackageName: *packageName,
		JSONFile:    *jsonFile,
	}, nil
}
