package usecases

import (
	"fmt"
	"os"
	"time"

	e "github.com/muzudho/kifuwarabe-uec12/entities"
	p "github.com/muzudho/kifuwarabe-uec12/presenter"
)

// SelfplayV9 - コンピューター同士の対局。
func SelfplayV9(board e.IBoard, printBoardType1 func(e.IBoard), printBoardType2 func(e.IBoard, int)) {
	color := 1

	for {
		fUCT := 1
		if color == 1 {
			fUCT = 0
		}
		z := e.GetComputerMoveV9(board, color, fUCT, printBoardType1)
		e.AddMovesV8(board, z, color, printBoardType2)
		if z == 0 && e.Moves > 1 && e.Record[e.Moves-2] == 0 {
			break
		}
		if e.Moves > 300 {
			break
		} // too long
		color = e.FlipColor(color)
	}

	p.PrintSgf(e.Moves, e.Record)
}

// TestPlayoutV9 - 試しにプレイアウトする。
func TestPlayoutV9(board e.IBoard, printBoardType1 func(e.IBoard), printBoardType2 func(e.IBoard, int)) {
	e.FlagTestPlayout = 1
	board.Playout(1, printBoardType1)
	printBoardType2(board, e.Moves)
	p.PrintSgf(e.Moves, e.Record)
}

// PlayComputerMoveV9a - コンピューター・プレイヤーの指し手。 GoGoV9a から呼び出されます。
func PlayComputerMoveV9a(board e.IBoard, color int, fUCT int, printBoardType1 func(e.IBoard), printBoardType2 func(e.IBoard, int)) int {
	var z int
	st := time.Now()
	e.AllPlayouts = 0
	if fUCT != 0 {
		z = e.GetBestUctV9a(board, color, printBoardType1)
	} else {
		z = board.PrimitiveMonteCalro(color, printBoardType1)
	}
	t := time.Since(st).Seconds()
	fmt.Fprintf(os.Stderr, "%.1f sec, %.0f playoutV9/sec, play_z=%2d,moves=%d,color=%d,playouts=%d\n",
		t, float64(e.AllPlayouts)/t, e.Get81(z), e.Moves, color, e.AllPlayouts)
	e.AddMoves9a(board, z, color, t, printBoardType2)
	return z
}

// TestPlayoutV9a - 試しにプレイアウトする。
func TestPlayoutV9a(board e.IBoard, printBoardType1 func(e.IBoard), printBoardType2 func(e.IBoard, int)) {
	e.FlagTestPlayout = 1
	board.Playout(1, printBoardType1)
	printBoardType2(board, e.Moves)
	p.PrintSgf(e.Moves, e.Record)
}

func UndoV9() {
	// Unimplemented.
}