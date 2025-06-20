package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func getImportList() ([]string, error) {
	fset := token.NewFileSet()
	imports := make([]string, 0)

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil || !info.Mode().IsRegular() || filepath.Ext(info.Name()) != ".go" {
			return nil
		}

		file, err := parser.ParseFile(fset, path, nil, parser.AllErrors|parser.ImportsOnly)
		if err != nil {
			return fmt.Errorf("File parsing error %q: %v\n", path, err)
		}

		for _, imp := range file.Imports {
			if imp.Path.Value != "" && imp.Path.Value[0] == '"' {
				imports = append(imports, imp.Path.Value[1:len(imp.Path.Value)-1])
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return imports, nil
}
