package tax

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

// validImportBan checks if there is a forbidden package in the import list
func validImportBan(name string) (bool, error) {
	im, err := getImportList()
	if err != nil {
		return false, err
	}

	if _, ok := im[name]; ok {
		return false, nil
	}
	return true, nil
}

// validImportMaxCount checks if the number of imported packages crosses the limit with count
func validImportMaxCount(count int) (bool, error) {
	im, err := getImportList()
	if err != nil {
		return false, err
	}

	if len(im) > count {
		return false, nil
	}
	return true, nil
}

// getImportList outputs all imports without repetition
func getImportList() (map[string]struct{}, error) {
	fset := token.NewFileSet()
	imports := make(map[string]struct{})

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
				imports[imp.Path.Value[1:len(imp.Path.Value)-1]] = struct{}{}
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return imports, nil
}
