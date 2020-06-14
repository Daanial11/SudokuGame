package main

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"fmt"
	
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}



func createBoard() [9][9]int {
	newBoard := [9][9]int{}

	file, err := ioutil.ReadFile(filepath.FromSlash("puzzles/easy.txt"))
	check(err)

	boardData := strings.Split(string(file), "\n")

	i := 0

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			n, err := strconv.Atoi(string(boardData[0][i]))
			check(err)
			newBoard[y][x] = n
			i++
		}
	}

	return newBoard
}

func boardLogic(boardChan chan [9][9]int ){
	fmt.Println("test")
	for {
		select {
		case board := <- boardChan:
			boardChan <- validBoardCheck(board)
		default:
					
		}
	}
	
}

func validBoardCheck(board [9][9]int) [9][9]int{
	validBoard := [9][9]int{}

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			validBoard[y][x] = 0
		}
	}
	
	validBoard[1][3] = 1
	validBoard[5][7] = 1
	validBoard[3][2] = 1

	return validBoard
}

func main() {

	boardChannel := make(chan [9][9]int)
	go boardLogic(boardChannel)
	display(createBoard(), boardChannel)

}
