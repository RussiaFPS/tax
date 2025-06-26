package tax

import (
	"testing"
)

func TestGetImportsList(t *testing.T) {
	validImportList := map[string]struct{}{
		"fmt":           {},
		"go/parser":     {},
		"go/token":      {},
		"go/types":      {},
		"os":            {},
		"path/filepath": {},
		"testing":       {},
	}

	imports, err := getImportList()
	if err != nil {
		t.Fatal(err)
	}

	for v := range imports {
		if _, ok := validImportList[v]; !ok {
			t.Errorf("not available import: %s", v)
		}
	}
}
