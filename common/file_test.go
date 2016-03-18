package common

import "testing"

func TestIsImage(t *testing.T) {
	exts := [...]string{".png", ".jpeg", ".GIF", "BMP", "tiff"}

	for _, val := range exts {
		if !IsImage(val) {
			t.Errorf("Expected %s is image", val)
		}
	}
}

func TestFileIsImage(t *testing.T) {
	f := MakeNewFile("test.jpeg")

	if f.IsImage == false {
		t.Errorf("Expected test.jpeg is image")
	}
}
