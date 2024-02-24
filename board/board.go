/*
Copyright 2024 Eduardo Lezcano Álvarez
*/
package board

import (
	"fmt"
	"math/rand"
	"time"
)

type Board[T any] struct {
	width, height int
	len           int
	data          []T
	rand          *rand.Rand
}

func New[T any](width, height int) *Board[T] {
	size := width * height
	board := Board[T]{
		width:  width,
		height: height,
		len:    size,
		data:   make([]T, size),
	}
	board.Seed(time.Now().UnixNano())
	return &board
}

func (board *Board[T]) Len() int {
	return board.len
}

func (board *Board[T]) Width() int {
	return board.width
}

func (board *Board[T]) Height() int {
	return board.height
}

func (board *Board[T]) IdOk(id int) error {
	if id < 0 || id >= board.len {
		return fmt.Errorf("Id %v out of the Board", id)
	}
	return nil
}

func (board *Board[T]) PosOk(x, y int) error {
	if x < 0 || y < 0 || x >= board.width || y >= board.height {
		return fmt.Errorf("Position %v, %v out of the Board", x, y)
	}
	return nil
}

func (board *Board[T]) RowOk(y int) error {
	if y < 0 || y >= board.height {
		return fmt.Errorf("Row %v out of the Board", y)
	}
	return nil
}

func (board *Board[T]) Id(x, y int) (int, error) {
	if err := board.PosOk(x, y); err != nil {
		return 0, err
	}
	return board.id(x, y), nil
}

func (board *Board[T]) id(x, y int) int {
	return x + y*board.width
}

func (board *Board[T]) Pos(id int) (int, int, error) {
	if err := board.IdOk(id); err != nil {
		return 0, 0, err
	}
	x, y := board.pos(id)
	return x, y, nil
}

func (board *Board[T]) pos(id int) (int, int) {
	return id % board.width, id / board.width
}

func (board *Board[T]) Seed(seed int64) {
	board.rand = rand.New(rand.NewSource(seed))
	board.rand.Seed(seed)
}

func (board *Board[T]) Tile(id int) (*T, error) {
	if err := board.IdOk(id); err != nil {
		return nil, err
	}
	return board.get(id), nil
}

func (board *Board[T]) get(id int) *T {
	return &board.data[id]
}

func (board *Board[T]) PermIds(n int) ([]int, error) {
	if n < 0 || n >= board.Len() {
		return nil, fmt.Errorf("PermIds(%v) out of range", n)
	}
	return rand.Perm(board.Len())[0:n], nil
}

func (board *Board[T]) PermTiles(n int) ([]*T, error) {
	ids, err := board.PermIds(n)
	if err != nil {
		return nil, err
	}
	tiles := make([]*T, n)
	for i, id := range ids {
		tiles[i] = board.get(id)
	}
	return tiles, nil
}

func (board *Board[T]) RowTiles(y int) ([]*T, error) {
	if err := board.RowOk(y); err != nil {
		return nil, err
	}
	tiles := make([]*T, board.width)
	for i, j := 0, y*board.width; i < board.width; i++ {
		tiles[i] = &board.data[j]
		j++
	}
	return tiles, nil
}

func (board *Board[T]) NeighboursIds(id int) ([]int, error) {
	ids := make([]int, 0, 8)
	x, y, err := board.Pos(id)
	if err != nil {
		return nil, err
	}
	// Not in Left column, Not in Right column, Not in Top row, Not in Bottom row
	nl, nr, nt, nb := x != 0, x != board.width-1, y != 0, y != board.height-1

	if nl && nt {
		ids = append(ids, id-board.width-1)
	}
	if nt {
		ids = append(ids, id-board.width)
	}
	if nt && nr {
		ids = append(ids, id-board.width+1)
	}
	if nl {
		ids = append(ids, id-1)
	}
	if nr {
		ids = append(ids, id+1)
	}
	if nl && nb {
		ids = append(ids, id+board.width-1)
	}
	if nb {
		ids = append(ids, id+board.width)
	}
	if nb && nr {
		ids = append(ids, id+board.width+1)
	}
	return ids, nil
}

func (board *Board[T]) NeighboursTiles(id int) ([]*T, error) {
	ids, err := board.NeighboursIds(id)
	if err != nil {
		return nil, err
	}
	tiles := make([]*T, len(ids))
	for i, id := range ids {
		tiles[i] = board.get(id)
	}
	return tiles, nil
}

func (board *Board[T]) AllTiles() []*T {
	tiles := make([]*T, board.len)
	for i := 0; i < board.len; i++ {
		tiles[i] = &board.data[i]
	}
	return tiles
}
