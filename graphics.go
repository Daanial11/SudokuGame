package main

import (
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten"
	_ "image/png"
	"log"
	"image/color"
	"github.com/hajimehoshi/ebiten/ebitenutil" // This is required to draw debug texts.
	"github.com/hajimehoshi/ebiten/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"strconv"
	"fmt"
	
	"math"
	"sync"
)
const (
	screenWidth = 432
	screenHeight = 432
	cellDiff = 48
	firstCellX = 13
	firstCellY = 35
	firstDotX = 21
	firstDotY = 40
)

var ( 
	grid *ebiten.Image
	cursorDot *ebiten.Image
	arcadeFont font.Face
	board [9][9]int
	validBoardValues [9][9]int
	changeFlag SafeBool

	boardMutex sync.Mutex
	

)
type Game struct {
	pressed []ebiten.Key
}


func update(screen *ebiten.Image) error {
	screen.Fill(color.White)

	
	screen.DrawImage(grid, nil)
	
	
	boardMutex.Lock()
	for y := 0; y<9; y++ {
		for x := 0; x<9; x++ {
			if board[y][x] != 0 && validBoardValues[y][x] != 1 {
				text.Draw(screen, strconv.Itoa(board[y][x]), arcadeFont, firstCellX + (cellDiff*x), firstCellY + (cellDiff*y), color.Black)

			} else if board[y][x] != 0 {
				text.Draw(screen, strconv.Itoa(board[y][x]), arcadeFont, firstCellX + (cellDiff*x), firstCellY + (cellDiff*y), color.RGBA{255, 17, 0, 255})
			}
		}
	}
	boardMutex.Unlock()
	x, y := ebiten.CursorPosition()
	

	if x > 0 && y > 0 && x < screenWidth && y < screenHeight {
		op := &ebiten.DrawImageOptions{}
	
		op.GeoM.Scale(0.025, 0.025)

		op.GeoM.Translate(firstDotX + math.Floor(float64(x)/float64(48))*48, firstDotY + math.Floor(float64(y)/float64(48))*48)
		screen.DrawImage(cursorDot, op)
	}
	
	for k := 1; k<10; k++{
		if ebiten.IsKeyPressed(ebiten.Key(k)) && !changeFlag.Get() && board[y/48][x/48] !=k {
			boardMutex.Lock()
			board[y/48][x/48] = k
			boardMutex.Unlock()
			changeFlag.Set(true)

		}
	}
	
	
	
	return nil
}

func display(intialBoard [9][9]int, boardChan chan [9][9]int) {
	fmt.Println("1")
	board = intialBoard
	go boardSender(boardChan)
	err := ebiten.Run(update, screenHeight, screenHeight, 2, "Suduko")
	check(err)

}

func boardSender(boardChan chan [9][9]int){

	for {
		if (changeFlag.Get() == true) {
			fmt.Println("test")
			boardMutex.Lock()
			boardChan <- board
			validBoardValues = <-boardChan
			boardMutex.Unlock()
			
			changeFlag.Set(false)
		}
		
	}
	
}

 

func init() {
	var err error
	grid, _, err = ebitenutil.NewImageFromFile("images/grid.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	cursorDot, _, err = ebitenutil.NewImageFromFile("images/dot.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	
	tt, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}


