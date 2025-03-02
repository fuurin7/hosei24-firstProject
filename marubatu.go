package main

import (
	"fmt"
	"math"
	"time"
)

const (
	Empty = " "
	maru  = "○"
	batu  = "×"
)

type Move struct {
	Row, Col int
}

func setBoard() [3][3]string {
	var board [3][3]string
	for i := range board {
		for j := range board[i] {
			board[i][j] = Empty
		}
	}
	return board
}

func printBoard(board [3][3]string) {
	fmt.Println()
	for i, row := range board {
		fmt.Println(row[0] + "|" + row[1] + "|" + row[2])
		if i < 2 {
			fmt.Println("-+-+-")
		}
	}
	fmt.Println()
}

func judge(board [3][3]string) string {
	for _, row := range board {
		if row[0] == row[1] && row[1] == row[2] && row[0] != Empty {
			return row[0]
		}
	}
	for col := 0; col < 3; col++ {
		if board[0][col] == board[1][col] && board[1][col] == board[2][col] && board[0][col] != Empty {
			return board[0][col]
		}
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != Empty {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] != Empty {
		return board[0][2]
	}
	for _, row := range board {
		for _, cell := range row {
			if cell == Empty {
				return ""
			}
		}
	}
	return "draw"
}

func minimax(board [3][3]string, depth int, isMaximizing bool, ai_Symbol, human_Symbol string) int {
	winner := judge(board)
	if winner == ai_Symbol {
		return 10 - depth
	} else if winner == human_Symbol {
		return depth - 10
	} else if winner == "draw" {
		return 0
	}

	if isMaximizing {
		bestScore := math.Inf(-1)
		for r := 0; r < 3; r++ {
			for c := 0; c < 3; c++ {
				if board[r][c] == Empty {
					board[r][c] = ai_Symbol
					score := minimax(board, depth+1, false, ai_Symbol, human_Symbol)
					board[r][c] = Empty
					bestScore = math.Max(bestScore, float64(score))
				}
			}
		}
		return int(bestScore)
	} else {
		bestScore := math.Inf(1)
		for r := 0; r < 3; r++ {
			for c := 0; c < 3; c++ {
				if board[r][c] == Empty {
					board[r][c] = human_Symbol
					score := minimax(board, depth+1, true, ai_Symbol, human_Symbol)
					board[r][c] = Empty
					bestScore = math.Min(bestScore, float64(score))
				}
			}
		}
		return int(bestScore)
	}
}

func findBestMove(board [3][3]string, ai_Symbol, human_Symbol string) Move {
	bestScore := math.Inf(-1)
	bestMove := Move{-1, -1}
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[r][c] == Empty {
				board[r][c] = ai_Symbol
				score := minimax(board, 0, false, ai_Symbol, human_Symbol)
				board[r][c] = Empty
				if float64(score) > bestScore {
					bestScore = float64(score)
					bestMove = Move{r, c}
				}
			}
		}
	}
	return bestMove
}

func check(board [3][3]string, move Move) bool {
	return move.Row >= 0 && move.Row < 3 && move.Col >= 0 && move.Col < 3 && board[move.Row][move.Col] == Empty
}

func human(board [3][3]string) Move {
	for {
		var p int
		fmt.Println("1|2|3")
		fmt.Println("-+-+-")
		fmt.Println("4|5|6")
		fmt.Println("-+-+-")
		fmt.Println("7|8|9")
		fmt.Println("場所を入力")
		fmt.Printf(">")
		fmt.Scan(&p)
		var move Move
		switch p {
		case 1:
			move = Move{0, 0}
			if !check(board, move) {
				fmt.Println("無効な入力。もう一度入力してください。")
				continue
			}
		case 2:
			move = Move{0, 1}
			if !check(board, move) {
				fmt.Println("無効な入力。もう一度入力してください。")
				continue
			}
		case 3:
			move = Move{0, 2}
			if !check(board, move) {
				fmt.Println("無効な入力。もう一度入力してください。")
				continue
			}
		case 4:
			move = Move{1, 0}
			if !check(board, move) {
				fmt.Println("無効な入力。もう一度入力してください。")
				continue
			}
		case 5:
			move = Move{1, 1}
			if !check(board, move) {
				fmt.Println("無効な入力。もう一度入力してください。")
				continue
			}
		case 6:
			move = Move{1, 2}
			if !check(board, move) {
				fmt.Println("無効な入力。もう一度入力してください。")
				continue
			}
		case 7:
			move = Move{2, 0}
			if !check(board, move) {
				fmt.Println("無効な入力。もう一度入力してください。")
				continue
			}
		case 8:
			move = Move{2, 1}
			if !check(board, move) {
				fmt.Println("無効な入力。もう一度入力してください。")
				continue
			}
		case 9:
			move = Move{2, 2}
			if !check(board, move) {
				fmt.Println("無効な入力。もう一度入力してください。")
				continue
			}
		default:
			fmt.Println("無効な入力。もう一度入力してください。")
			continue
		}
		return move
	}
}

func human_Vs_AI(player, ai string) {
	board := setBoard()
	turn := maru
	printBoard(board)
	for {
		winner := judge(board)
		if winner != "" {
			if winner == "draw" {
				fmt.Println("引き分け！")
			} else {
				if winner == player {
					fmt.Print("あなたの勝ち！\n")
				} else {
					fmt.Print("あなたの負け！\n")
				}
			}
			return
		}
		if turn == player {
			move := human(board)
			board[move.Row][move.Col] = player
		} else {
			fmt.Println("AIの手を考えています...")
			time.Sleep(1 * time.Second)
			move := findBestMove(board, ai, player)
			board[move.Row][move.Col] = ai
		}
		printBoard(board)
		if turn == maru {
			turn = batu
		} else {
			turn = maru
		}
	}
}

func AI_Vs_AI() {
	board := setBoard()
	turn := maru
	printBoard(board)
	for {
		fmt.Printf("%s の番...", turn)
		time.Sleep(1 * time.Second)
		move := findBestMove(board, maru, batu)
		if turn == batu {
			move = findBestMove(board, batu, maru)
		}
		board[move.Row][move.Col] = turn
		printBoard(board)
		winner := judge(board)
		if winner != "" {
			if winner == "draw" {
				fmt.Println("引き分け！")
			} else {
				fmt.Printf("%s の勝ち！\n", winner)
			}
			break
		}
		if turn == maru {
			turn = batu
		} else {
			turn = maru
		}
	}
}

func main() {
	for {
		var mode int
		fmt.Println("ゲームモードを選択（1: 人 vs AI, 2: AI vs AI）")
		fmt.Printf(">")
		fmt.Scan(&mode)
		switch mode {
		case 1:
			var player, ai string
			for {
				var choice int
				fmt.Print("先攻後攻の選択（1: ○, 2: ×）")
				fmt.Printf(">")
				fmt.Scan(&choice)
				switch choice {
				case 1:
					player, ai = maru, batu
					human_Vs_AI(player, ai)
					return
				case 2:
					player, ai = batu, maru
					human_Vs_AI(player, ai)
					return
				default:
					fmt.Println("無効な選択。1または2を入力してください。")
					continue
				}
			}
		case 2:
			AI_Vs_AI()
			return
		default:
			fmt.Println("無効な選択。1または2を入力してください。")
			continue
		}
	}
}
