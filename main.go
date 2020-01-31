package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createBoard(difficulty string) [9][9]int {
	newBoard := [9][9]int{}
	difficulty = strings.TrimSuffix(difficulty, "\n")
	file, err := ioutil.ReadFile("puzzles/" + difficulty + ".txt")
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter difficulty to play:")
	var difficulty string
	difficulty, err := reader.ReadString('\n')

	check(err)

	createBoard(difficulty)

}
