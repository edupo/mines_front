/*
Copyright 2024 Eduardo Lezcano Álvarez
*/
package mines

const (
	None      uint8 = 0b00000000
	Uncovered uint8 = 0b00000001
)

type Tile struct {
	Mines   uint8
	Around  uint8
	Flags   uint8
	Visible uint8
	State   uint8
	_       uint8
	_       uint8
	_       uint8
}

func (ms *Tile) IsUncovered() bool {
	return (ms.State & Uncovered) != 0
}

func (ms *Tile) SetUncovered(status bool) {
	if status {
		ms.State |= Uncovered
	} else {
		ms.State &= ^Uncovered
	}
}

func (ms *Tile) Repr() byte {
	// if !ms.IsUncovered() {
	// 	return '~'
	// }
	if ms.Mines > 0 {
		return '*'
	}
	if ms.Around > 0 {
		return '0' + ms.Around // TODO: More than 9 around
	}
	return '.'
}
