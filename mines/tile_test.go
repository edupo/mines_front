package mines_test

import (
	"mines"
	"testing"
)

func TestMinesTile(t *testing.T) {
	ms := &mines.Tile{}
	if v := ms.IsUncovered(); v != false {
		t.Errorf("IsUncovered() = %v, want %v", v, false)
	}
	ms.SetUncovered(true)
	if v := ms.IsUncovered(); !v {
		t.Errorf("IsUncovered() = %v, want %v", v, true)
	}
	ms.SetUncovered(false)
	if v := ms.IsUncovered(); v {
		t.Errorf("IsUncovered() = %v, want %v", v, false)
	}
}
