package triangle

import (
	"testing"
)

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{12, 35, 0},
	}

	for _, tt := range tests {
		if actual := Calctriangel(tt.a, tt.b); actual != tt.c {
			t.Errorf("Calctriangel(%d ,%d )"+"got %d;expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}
