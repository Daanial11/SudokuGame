package main

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	
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

func main() {

	
	display(createBoard())

}
