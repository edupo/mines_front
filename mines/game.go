package mines

import (
	"board"
)

type Game struct {
	*board.Board[Tile]
	mines int
}

func NewGame(width, height, mines int) *Game {
	return &Game{
		board.New[Tile](width, height),
		mines,
	}
}

func (g *Game) InitializeClassic() error {
	ids, err := g.PermIds(g.mines)
	if err != nil {
		return err
	}
	for _, id := range ids {
		tile, _ := g.Tile(id)
		tile.Mines = 1
		ns, err := g.NeighboursTiles(id)
		if err != nil {
			return err
		}
		for _, n := range ns {
			n.Around += 1
		}
	}
	return nil
}

func (g *Game) Run(c Command) error {
	switch c.Action {
	case ActionUncover:
		return g.Uncover(c.Id)
	case ActionFlag:
		return g.Flag(c.Id)
	}
	return nil
}

func (g *Game) Flag(id int) error {
	tile, err := g.Tile(id)
	if err != nil {
		return err
	}
	if tile.Flags > 0 {
		tile.Flags = 0
	} else {
		tile.Flags = 1
	}
	return nil
}

func (g *Game) Uncover(id int) error {
	tile, err := g.Tile(id)
	if err != nil {
		return err
	}
	if tile.Around > 0 {
		tile.SetUncovered(true)
		return nil
	}
	if err := g.uncover(id); err != nil {
		return err
	}
	return nil
}

func (g *Game) uncover(id int) error {
	tile, _ := g.Tile(id)
	if tile.IsUncovered() {
		return nil
	}
	tile.SetUncovered(true)
	if tile.Around > 0 {
		return nil
	}
	ids, err := g.NeighboursIds(id)
	if err != nil {
		return err
	}
	for _, id = range ids {
		g.uncover(id)
	}
	return nil
}

func (g *Game) Print() error {
	for y := range g.Height() {
		str := make([]byte, g.Width())
		row, err := g.RowTiles(y) 
		if err != nil {
			return err
		}
		for x, tile := range row{
			str[x] = tile.Repr()
		}
		println(string(str))
	}
	return nil
}
