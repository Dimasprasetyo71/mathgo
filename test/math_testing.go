package mathtesting

import (
	"testing"

	"github.com/dimasprasetyo71/mathgo/math"
)

func Tambah(t *testing.T) {
	result := math.Tambah(1, 2)
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}


