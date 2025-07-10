package tax

import (
	"testing"
)

func TestValidImportBan(t *testing.T) {
	importName := "fmt123"
	isValid, err := validImportBan(importName)
	if err != nil {
		t.Fatal(err)
	}

	if !isValid {
		t.Errorf("not valid import %s is banned", importName)
	}
}

func TestValidImportMaxCount(t *testing.T) {
	isValid, err := validImportMaxCount(6)
	if err != nil {
		t.Fatal(err)
	}

	if !isValid {
		t.Error("not valid max count imports")
	}
}

func TestGetImportsList(t *testing.T) {
	validImportList := map[string]struct{}{
		"fmt":           {},
		"go/parser":     {},
		"go/token":      {},
		"os":            {},
		"path/filepath": {},
		"testing":       {},
	}

	imports, err := getImportList()
	if err != nil {
		t.Fatal(err)
	}

	for v := range validImportList {
		if _, ok := imports[v]; !ok {
			t.Errorf("not available import: %s", v)
		}
	}
}
