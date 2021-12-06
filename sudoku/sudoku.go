package main

import (
	"os"
)

// Sudoku final sudoku
var Sudoku [9][9]rune

func main() {

	args := os.Args[1:]
	if validateArgs(args) {
		sudoku := extractSudoku(args)

		if solve(sudoku, sudoku) {
			if checkSum(Sudoku) {
				printSudoku(Sudoku)
			} else {
				print("Error")
			}
		} else {
			print("Error")
		}
	} else {
		print("Error")
	}

}

// imprimer le Sudoku
func printSudoku(sudoku [9][9]rune) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			print(sudoku[i][j] - 48)
			if j != 8 {
				print(" ")
			}
		}
		println()
	}
}

// fonction permmettant de resoudre le sudoku (avec les différents retour à faire)
func solve(sudoku [9][9]rune, prevSudoku [9][9]rune) bool {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == '.' {

				possibilities := possibilities(i, j, sudoku)

				if len(possibilities) == 0 {
					sudoku = copyIt(prevSudoku)
					return false
				}

				for _, sVar := range possibilities {

					sudoku[i][j] = sVar

					if !solve(sudoku, prevSudoku) {
						continue
					} else {
						return true
					}
				}
				return false
			}
		}
	}
	Sudoku = copyIt(sudoku)
	return true
}

// Tous les types de possibilitées
func possibilities(i, j int, sudoku [9][9]rune) []rune {

	var possibilities [9]rune
	var iterPos = 0
	for pos := '1'; pos <= '9'; pos++ {
		if foundInSquare(i, j, pos, sudoku) {
			continue
		}
		if foundInRow(i, j, pos, sudoku) {
			continue
		}
		if foundInCol(i, j, pos, sudoku) {
			continue
		}

		possibilities[iterPos] = pos
		iterPos++
	}

	return possibilities[0:iterPos]
}

// fonction carré (toutes les possibilitées)
func foundInSquare(i, j int, pos rune, arr [9][9]rune) bool {
	ki := i / 3
	kj := j / 3
	for ii := 3 * ki; ii < 3*(ki+1); ii++ {
		for jj := 3 * kj; jj < 3*(kj+1); jj++ {
			if arr[ii][jj] == pos {
				return true
			}
		}
	}
	return false
}

// fonction rangée (toutes les possibilitées)
func foundInRow(i, j int, pos rune, arr [9][9]rune) bool {

	for ii := 0; ii < 9; ii++ {
		if ii != j {
			if arr[i][ii] == pos {
				return true
			}
		}
	}
	return false
}

// fonction colonne (toutes les possibilitées)
func foundInCol(i, j int, pos rune, arr [9][9]rune) bool {

	for ii := 0; ii < 9; ii++ {
		if ii != i {
			if arr[ii][j] == pos {
				return true
			}
		}
	}
	return false
}

func copyIt(arr [9][9]rune) [9][9]rune {

	var sudoku [9][9]rune

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sudoku[i][j] = arr[i][j]
		}
	}
	return sudoku
}

// créer le tableau sudoku
func extractSudoku(args []string) [9][9]rune {

	var sudoku [9][9]rune

	for i := 0; i < 9; i++ {
		runes := strToRunes(args[i])
		for j := 0; j < 9; j++ {

			sudoku[i][j] = runes[j]
		}
	}
	return sudoku
}

// permet de passer de string à Rune
func strToRunes(str string) []rune {
	runes := make([]rune, len(str))
	for i, runa := range str {
		runes[i] = runa
	}
	return runes
}

// permet la validification des arguments
func validateArgs(args []string) bool {

	emptyCount := 0

	for _, w := range args {
		for _, s := range strToRunes(w) {
			if s == '.' {
				emptyCount++
			} else {
				if s < '1' || s > '9' {
					return false
				}
			}
		}
	}

	if emptyCount == 81 {
		return false
	}

	if len(args) != 9 {
		return false
	}
	for _, w := range args {
		if len(w) != 9 {
			return false
		}
	}
	return true
}

// permet de controler l'ensemble du tableau
func checkSum(sudoku [9][9]rune) bool {

	rowSum := 0

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			rowSum += toInt(sudoku[i][j])
		}
		if rowSum != 45 {
			return false
		}
		rowSum = 0
	}

	colSum := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			colSum += toInt(sudoku[j][i])
		}
		if colSum != 45 {
			return false
		}
		colSum = 0

	}

	allCellsSum := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			allCellsSum += toInt(sudoku[i][j])
		}
	}
	return allCellsSum == 405
}
// Transforme la Rune en int 
func toInt(r rune) int {
	counter := 1
	for i := '1'; i <= '9'; i++ {
		if r == i {
			return counter
		}
		counter++

	}
	return counter
}

