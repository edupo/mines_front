package mines

import "encoding/json"

type TileState struct {
	id        int
	around    uint8
	flags     uint8
	mines     uint8
	uncovered bool
}

type GameState struct {
	ID    int         `json:"id"`
	Mines int         `json:"mines"`
	Board []TileState `json:"board"`
}

func (mg *Game) State() GameState {
	state := GameState{
		ID:    0,
		Mines: mg.mines,
		Board: make([]TileState, mg.Len()),
	}
	for i, tile := range mg.AllTiles() {
		state.Board[i] = TileState{
			id:        i,
			uncovered: tile.IsUncovered(),
			around:    tile.Around,
			mines:     tile.Mines,
			flags:     tile.Flags,
		}
	}
	return state
}

func (ts TileState) MarshalJSON() ([]byte, error) {
	if ts.uncovered {
		return json.Marshal(&struct {
			ID     int   `json:"id"`
			Around uint8 `json:"around"`
			Mines  uint8 `json:"mines,omitempty"`
		}{
			ID:     ts.id,
			Around: ts.around,
			Mines:  ts.mines,
		})
	} else {
		return json.Marshal(&struct {
			ID    int `json:"id"`
			Flags int `json:"flags,omitempty"`
		}{
			ID:    ts.id,
			Flags: int(ts.flags),
		})
	}
}
