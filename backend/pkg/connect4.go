package connect4

import (
	"fmt"
)

type BoardState = uint8
type Player = uint8

const (
	StateEmpty	BoardState = 0
	StateYellow BoardState = 1
	StateRed 	BoardState = 2
)

const (
	PlayerYellow Player = 1
	PlayerRed    Player = 2
)

const BoardWidth = 7
const BoardHeight = 6

type Board struct {
	cells [BoardWidth * BoardHeight]uint8 `json:"cells"`
}

func NewBoard() *Board {
	board := Board{}

	return &board
}

func (board *Board) getCell(x uint8, y uint8) uint8 {
	return board.cells[y *BoardWidth + x]
}

func (board *Board) setCell(x uint8, y uint8, state BoardState) {
	board.cells[y * BoardWidth + x] = state
}

func (board *Board) playAt(x uint8, y uint8, player Player) bool {
	cell := board.getCell(x, y)

	if cell != StateEmpty {
		return false
	}

	if y != 0 && board.getCell(x, y - 1) != StateEmpty {
		// Cannot place a piece above an empty space
		return false
	}

	switch player {
	case PlayerYellow:
		board.setCell(x, y, StateYellow)
		break
	case PlayerRed:
		board.setCell(x, y, StateRed)
		break
	default:
		panic(fmt.Sprintf("Wrong value passed to playAt: %v", player))
		return false
	}

	return true
}

func (board *Board) isWinningRow(x uint8, y uint8) bool {
	cell := board.getCell(x, y)
	count := 0

	for position := int16(x); position >= 0; position-- {
		if board.getCell(uint8(position), y) != cell {
			break
		}
		count++
	}

	for position := x; position < BoardWidth; position++ {
		if board.getCell(position, y) != cell {
			break
		}
		count++
	}

	return count >= 4
}

func (board *Board) isWinningColumn(x uint8, y uint8) bool {
	cell := board.getCell(x, y)
	count := 0

	for position := int16(y); position >= 0; position-- {
		if board.getCell(x, uint8(position)) != cell {
			break
		}
		count++
	}

	for position := x; position < BoardHeight; position++ {
		if board.getCell(position, y) != cell {
			break
		}
		count++
	}

	return count >= 4
}

func (board *Board) isWinningDiagonal(x uint8, y uint8) bool {
	cell := board.getCell(x, y)
	countFirstDiagonal := 0
	intX := int(x)
	intY := int(y)

	for delta := 0; delta < 4; delta++ {
		if intX + delta < BoardWidth && intY + delta < BoardHeight &&
			board.getCell(uint8(intX + delta), uint8(intY + delta)) == cell {
			countFirstDiagonal++
		}
		if intX - delta >= 0 && intY - delta >= 0 &&
			board.getCell(uint8(intX - delta), uint8(intY - delta)) == cell {
			countFirstDiagonal++
		}
	}

	if countFirstDiagonal >= 4 {
		return true
	}

	countSecondDiagonal := 0

	for delta := 0; delta < 4; delta++ {
		if intX + delta < BoardWidth && intY - delta >= 0 &&
			board.getCell(uint8(intX + delta), uint8(intY - delta)) == cell {
			countSecondDiagonal++
		}
		if intX - delta >= 0 && intY + delta < BoardHeight &&
			board.getCell(uint8(intX - delta), uint8(intY + delta)) == cell {
			countSecondDiagonal++
		}
	}

	return countSecondDiagonal >= 4
}

func (board *Board) isWinningPosition(x uint8, y uint8) bool {
	return board.isWinningRow(x, y) || board.isWinningColumn(x, y) || board.isWinningDiagonal(x, y)
}