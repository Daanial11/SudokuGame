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
)
const (
	screenWidth = 432
	screenHeight = 432
	cellDiff = 48
	firstCellX = 13
	firstCellY = 35
)

var ( 
	grid *ebiten.Image
	arcadeFont font.Face
	board [9][9]int
)


func update(screen *ebiten.Image) error {
	screen.Fill(color.White)

	
	screen.DrawImage(grid, nil)

	for y := 0; y<9; y++ {
		for x := 0; x<9; x++ {
			if board[y][x] != 0 {
				text.Draw(screen, strconv.Itoa(board[y][x]), arcadeFont, firstCellX + (cellDiff*x), firstCellY + (cellDiff*y), color.Black)

			}
		}
	}
	x, y := ebiten.CursorPosition()
	fmt.Println(x, y)
	
	
	
	
	
	return nil
}

func display(b [9][9]int) {
	board = b
	err := ebiten.Run(update, screenHeight, screenHeight, 2, "Suduko")
	check(err)

}

 

func init() {
	var err error
	grid, _, err = ebitenutil.NewImageFromFile("images/grid.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	tt, err1 := truetype.Parse(fonts.ArcadeN_ttf)
	if err1 != nil {
		log.Fatal(err)
	}
	arcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}
