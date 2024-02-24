package board

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestBoard(w, h int) *Board[int] {
	board := New[int](w, h)
	board.Seed(42)
	for i := 0; i < board.Len(); i++ {
		board.data[i] = i
	}
	return board
}

func TestNewBoard(t *testing.T) {
	board := newTestBoard(5, 10)
	if w, v := board.Width(), 5; w != v {
		t.Errorf("New(5, 10).Width() = %v, want %v", w, v)
	}
	if h, v := board.Height(), 10; h != v {
		t.Errorf("New(5, 10).Height() = %v, want %v", h, v)
	}
	if l, v := board.Len(), 50; l != v {
		t.Errorf("New(5, 10).Len() = %v, want %v", l, v)
	}
	if board.data == nil || reflect.TypeOf(board.data[0]).String() != "int" {
		t.Fatalf("New(5, 10).data = %v, want []int", reflect.TypeOf(board.data).String())
	}
}

func TestPos(t *testing.T) {
	board := newTestBoard(10, 10)
	for _, tt := range []struct {
		id, x, y int
		err      error
	}{
		{-1, 0, 0, fmt.Errorf("")},
		{0, 0, 0, nil},
		{1, 1, 0, nil},
		{15, 5, 1, nil},
		{99, 9, 9, nil},
		{100, 0, 0, fmt.Errorf("")},
	} {
		if x, y, err := board.Pos(tt.id); x != tt.x ||
			y != tt.y || reflect.TypeOf(err) != reflect.TypeOf(tt.err) {
			t.Errorf("Pos(%v) = %v, %v, %v, want %v, %v, %v",
				tt.id, x, y, err, tt.x, tt.y, tt.err)
		}
	}
}

func TestId(t *testing.T) {
	board := newTestBoard(5, 5)
	for _, tt := range []struct {
		x, y, id int
		err      error
	}{
		{-1, 0, 0, fmt.Errorf("")},
		{0, -1, 0, fmt.Errorf("")},
		{0, 0, 0, nil},
		{1, 0, 1, nil},
		{4, 4, 24, nil},
		{4, 5, 0, fmt.Errorf("")},
		{6, 5, 0, fmt.Errorf("")},
	} {
		if id, err := board.Id(tt.x, tt.y); id != tt.id ||
			reflect.TypeOf(err) != reflect.TypeOf(tt.err) {
			t.Errorf("Id(%v, %v) = %v, %v, want %v, %v",
				tt.x, tt.y, id, err, tt.id, tt.err)
		}
	}

	if _, err := board.Id(7, 15); err == nil {
		t.Error("id isn't out of range")
	}
}

func TestGet(t *testing.T) {
	board := newTestBoard(10, 10)
	for _, tt := range []struct {
		id    int
		ptile *int
		err   error
	}{
		{-1, nil, fmt.Errorf("")},
		{0, &board.data[0], nil},
		{62, &board.data[62], nil},
		{99, &board.data[99], nil},
		{100, nil, fmt.Errorf("")},
	} {
		if ptile, err := board.Tile(tt.id); ptile != tt.ptile ||
			reflect.TypeOf(err) != reflect.TypeOf(tt.err) {
			t.Errorf("Get(%v) = %v, want %v, %v",
				tt.id, ptile, tt.ptile, tt.err)
		}
	}
}

func TestRow(t *testing.T) {
	board := newTestBoard(10, 10)
	for _, tt := range []struct {
		row   int
		i     int
		ptile *int
		err   error
	}{
		{-1, 0, nil, nil},
		{0, 0, &board.data[0], nil},
		{0, 1, &board.data[1], nil},
		{9, 9, &board.data[99], nil},
		{10, 0, nil, nil},
	} {
		if row, err := board.RowTiles(tt.row); row != nil &&
			tt.i > 0 && tt.i < len(row)-1 && row[tt.i] != tt.ptile && err != tt.err {
			t.Errorf("RowTiles(%v)[%v] = %v, want %v",
				tt.row, tt.i, row[tt.i], tt.ptile)
		}

	}
}

func TestNeighbours(t *testing.T) {
	board := newTestBoard(3, 3)
	d := board.data
	for _, tt := range []struct {
		id int
		ns []*int
	}{
		{0, []*int{&d[1], &d[3], &d[4]}},
		{1, []*int{&d[0], &d[2], &d[3], &d[4], &d[5]}},
		{2, []*int{&d[1], &d[4], &d[5]}},
		{3, []*int{&d[0], &d[1], &d[4], &d[6], &d[7]}},
		{4, []*int{&d[0], &d[1], &d[2], &d[3], &d[5], &d[6], &d[7], &d[8]}},
		{5, []*int{&d[1], &d[2], &d[4], &d[7], &d[8]}},
		{6, []*int{&d[3], &d[4], &d[7]}},
		{7, []*int{&d[3], &d[4], &d[5], &d[6], &d[8]}},
		{8, []*int{&d[4], &d[5], &d[7]}},
	} {
		ns, err := board.NeighboursTiles(tt.id)
		if err != nil {
			t.Fatalf("This is not supposed to happen...")
		}
		for i, n := range ns {
			if n != tt.ns[i] {
				t.Errorf("NeighbourTiles(%v)[%v] = %v, want %v", tt.id, i, n, tt.ns[i])
			}
		}
	}
}

func TestPerm(t *testing.T) {
	board := newTestBoard(7, 11)
	for _, tt := range []struct {
		n   int
		l   int
		err error
	}{
		{-1, 0, fmt.Errorf("")},
		{0, 0, nil},
		{1, 1, nil},
		{board.Len() - 1, board.Len() - 1, nil},
		{board.Len(), 0, fmt.Errorf("")},
	} {

		if ids, err := board.PermIds(tt.n); len(ids) != tt.l ||
			reflect.TypeOf(err) != reflect.TypeOf(tt.err) {
			t.Errorf("PermIds(%v) = [%v],%v, want [%v],%v",
				tt.n, len(ids), err, tt.n, tt.err)
		}

	}
}

func TestPermTiles(t *testing.T) {
	board := newTestBoard(7, 11)
	for _, tt := range []struct {
		n   int
		l int
		err error
	}{
		{-1, 0, fmt.Errorf("")},
		{0, 0, nil},
		{1, 1, nil},
		{board.Len() - 1,board.Len() - 1,  nil},
		{board.Len(), 0, fmt.Errorf("")},
	} {

		if tiles, err := board.PermTiles(tt.n); len(tiles) != tt.l ||
			reflect.TypeOf(err) != reflect.TypeOf(tt.err) {
			t.Errorf("PermTiles(%v) = [%v],%v, want [%v],%v ",
				tt.n, len(tiles), err, tt.n, tt.err)
		}

	}
}

func TestAllTiles(t *testing.T) {
	board := newTestBoard(7, 11)
	if l := len(board.AllTiles()); l != board.len {
		t.Errorf("len(AllTiles()) = %v, want %v", l, board.len)
	}
}
