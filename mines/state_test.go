package mines

import (
	"encoding/json"
	"testing"
)

func TestStateMines(t *testing.T) {
	mines := 10
	board := NewGame(10, 10, mines)
	state := board.State()
	if state.Mines != mines {
		t.Errorf("Mines = %v, want %v", state.Mines, mines)
	}
}

func TestStateID(t *testing.T) {
	t.Skipf("ID not implemented")
}

func TestStateBoard(t *testing.T) {
	board := NewGame(10, 10, 10)
	state := board.State()
	for i, tile := range state.Board {
		board_tile, _ := board.Tile(tile.id)
		if v, ttv := tile.uncovered, board_tile.IsUncovered(); v != ttv {
			t.Errorf("Board[%v].uncovered = %v, want %v", i, v, ttv)
		}
		if v, ttv := tile.around, board_tile.Around; v != ttv {
			t.Errorf("Board[%v].around = %v, want %v", i, v, ttv)
		}
		if v, ttv := tile.mines, board_tile.Mines; v != ttv {
			t.Errorf("Board[%v].mines = %v, want %v", i, v, ttv)
		}
		if v, ttv := tile.flags, board_tile.Flags; v != ttv {
			t.Errorf("Board[%v].flags = %v, want %v", i, v, ttv)
		}
	}
}

func TestStateMarshall(t *testing.T) {
	board := NewGame(10, 10, 10)
	board.Uncover(15)
	_, err := json.Marshal(board.State())
	if err != nil {
		t.Error(err)
	}
}
