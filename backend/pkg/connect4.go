package connect4

const (
	EMPTY = 0
	YELLOW = 1
	RED = 2
)

const BoardWidth = 7
const BoardHeight = 6

type Board struct {
	cells [BoardWidth * BoardHeight]uint8
}

func NewBoard() *Board {
	board := Board{}

	return &board
}

func (board *Board) get(x uint8, y uint8) uint8 {
	return board.cells[y *BoardWidth + x]
}

