package connect4

import "testing"

func TestNewBoard(t *testing.T) {
	var board = NewBoard()

	for y := uint8(0); y < BoardHeight; y++ {
		for x := uint8(0); x < BoardWidth; x++ {
			if board.getCell(x, y) != StateEmpty {
				t.Error("Cell is not empty: ", x, y)
			}
		}
	}
}

