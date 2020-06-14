package main

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"fmt"
	"github.com/jwangsadinata/go-multimap/setmultimap"
	
	
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type cordinates struct{
	Y int
	X int
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

	//Check rows	
	for y := 0; y < 9; y++ {
		m := setmultimap.New()
		for x := 0; x < 9; x++ {
			
			m.Put(board[y][x], x)  
			validBoard[y][x] = 0
		}
		for n:= 1; n < 10; n++{
			cords, exist := m.Get(n)
			
			if exist {
				
				if len(cords)>1 {
					for _, v := range cords {
						
						
						validBoard[y][v.(int)] = 1
					}
					
					
				} 
			}
		}
		
	}

	//Check columns
	for x := 0; x < 9; x++ {
		m := setmultimap.New()
		for y := 0; y < 9; y++ {
			
			m.Put(board[y][x], y)  
			
		}
		for n:= 1; n < 10; n++{
			cords, exist := m.Get(n)
			
			if exist {
				
				if len(cords)>1 {
					for _, v := range cords {
						
						
						validBoard[v.(int)][x] = 1
					}
					
					
				} 
			}
		}
		
	}

	//Check boxes
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++{
			m := setmultimap.New()
			for r := 0; r < 3; r++ {
				m.Put(board[(y*3) + r][x*3], cordinates{Y: (y*3) + r, X:x*3})
				m.Put(board[(y*3) + r][(x*3) + 1], cordinates{Y: (y*3) + r, X:(x*3) + 1})
				m.Put(board[(y*3) + r][(x*3) + 2], cordinates{Y: (y*3) + r, X:(x*3) + 2})
			}
			
			for n:= 1; n < 10; n++{
				cords, exist := m.Get(n)
				if exist {
					if len(cords)>1 {
						for _, v := range cords {
							d := v.(cordinates)
							validBoard[d.Y][d.X] = 1
							
						}
						
						
					} 
				}
			}
		}
	}
	
	
	fmt.Println(validBoard)
	return validBoard
}

func main() {

	boardChannel := make(chan [9][9]int)
	go boardLogic(boardChannel)
	display(createBoard(), boardChannel)

}
