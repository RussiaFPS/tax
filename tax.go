package tax_go

import (
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

func GetNamesPackages() {
	fset := token.NewFileSet()
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil || !info.Mode().IsRegular() || filepath.Ext(info.Name()) != ".go" {
			return nil // пропускаем файлы, не относящиеся к Go
		}

		file, err := parser.ParseFile(fset, path, nil, parser.AllErrors|parser.ImportsOnly)
		if err != nil {
			log.Printf("Ошибка парсинга файла %q: %v\n", path, err)
			return nil
		}

		for _, imp := range file.Imports {
			if imp.Path.Value != "" && imp.Path.Value[0] == '"' { // проверка валидности пути импорта
				log.Printf("Импортированный пакет: %s\n", imp.Path.Value[1:len(imp.Path.Value)-1])
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
