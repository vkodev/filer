package sqlite3

import (
	"encoding/json"
	"github.com/vkodev/filer/common"
	"testing"
)

func TestModelConvertation(t *testing.T) {
	f := common.MakeNewFile("test.jpeg")
	m := &FileModel{}

	m.FromFile(f)

	if m.OriginalFilename != f.OriginalFilename {
		t.Errorf("Expected %s === %s", f.OriginalFilename, m.OriginalFilename)
	}

	f1 := m.ToFile()

	b1, err := json.Marshal(f)
	if err != nil {
		t.Error(err)
	}
	b2, err := json.Marshal(f1)
	if err != nil {
		t.Error(err)
	}

	if string(b1) != string(b2) {
		t.Errorf("Expected f1==f2 \n %s \n %s", b1, b2)
	}

}
