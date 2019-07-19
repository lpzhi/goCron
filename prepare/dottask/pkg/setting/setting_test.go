package setting

import (
	"testing"
)

func TestGetDefuultDatabase(t *testing.T) {
	if actual := GetDefuultDatabase(); actual != "test_central" {
		t.Errorf("GetDefuultDatabase( )"+"got %v;expected %v", actual, "test_central")
	}
}
