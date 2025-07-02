package tax

import (
	"testing"
)

func TestValidImportMaxCount(t *testing.T) {
	isValid, err := validImportMaxCount(6)
	if err != nil {
		t.Fatal(err)
	}

	if !isValid {
		t.Errorf("not valid max count imports")
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
