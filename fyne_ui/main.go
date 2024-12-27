package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const (
	boardSize            = 8          // Board size (8x8)
	squareSize           = 50         // Size of each square
	pieceScale           = 0.8        // Piece size relative to square size
	crownScale           = 0.36       // Crown size relative to square size
	crownWithHeightRatio = 1.3        // Crown width to height ratio
	crownStrokeWidth     = 2          // Crown stroke width
	lightSquareColor     = 0xf0d9b5ff // Light square color
	darkSquareColor      = 0xb58863ff // Dark square color
	crownColor           = 0xffd700ff // Gold crown color
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Checkers")

	board := container.NewWithoutLayout()
	createBoard(board)

	// Add pieces
	addPiece(board, 0, 1, color.Black, false) // Black normal piece
	addPiece(board, 1, 2, color.Black, true)  // Black queen
	addPiece(board, 2, 3, color.Black, false) // Black normal piece
	addPiece(board, 5, 0, color.White, false) // White normal piece
	addPiece(board, 6, 1, color.White, true)  // White queen
	addPiece(board, 7, 2, color.White, false) // White normal piece

	myWindow.SetContent(board)
	myWindow.Resize(fyne.NewSize(boardSize*squareSize, boardSize*squareSize))
	myWindow.ShowAndRun()
}

func createBoard(board *fyne.Container) {
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			squareColor := getSquareColor(row, col)
			rect := canvas.NewRectangle(squareColor)
			rect.Resize(fyne.NewSize(squareSize, squareSize))
			rect.Move(fyne.NewPos(float32(col*squareSize), float32(row*squareSize)))
			board.Add(rect)
		}
	}
}

func getSquareColor(row, col int) color.Color {
	if (row+col)%2 == 0 {
		return decodeColor(lightSquareColor)
	}
	return decodeColor(darkSquareColor)
}

func addPiece(board *fyne.Container, row, col int, pieceColor color.Color, isQueen bool) {
	piece := canvas.NewCircle(pieceColor)
	pieceDiameter := float32(squareSize) * pieceScale
	piece.Resize(fyne.NewSize(pieceDiameter, pieceDiameter))
	piece.Move(fyne.NewPos(
		float32(col*squareSize)+(float32(squareSize)-pieceDiameter)/2,
		float32(row*squareSize)+(float32(squareSize)-pieceDiameter)/2,
	))
	board.Add(piece)

	if isQueen {
		addCrown(board, row, col)
	}
}

func addCrown(board *fyne.Container, row, col int) {
	// X-axis increase from left to right
	// Y-axis increase from top to bottom
	centerX := float32(col*squareSize) + float32(squareSize)/2
	centerY := float32(row*squareSize) + float32(squareSize)/2
	crownHeight := float32(squareSize) * crownScale
	crownWidth := crownHeight * crownWithHeightRatio

	leftBottomCorner := fyne.NewPos(centerX+crownWidth/2, centerY+crownHeight/2)
	rightBottomCorner := fyne.NewPos(centerX-crownWidth/2, centerY+crownHeight/2)
	centerSpike := fyne.NewPos(centerX, centerY-crownHeight/2)
	leftSpike := fyne.NewPos(centerX+crownWidth*2/3, centerY-crownHeight/4)
	rightSpike := fyne.NewPos(centerX-crownWidth*2/3, centerY-crownHeight/4)
	leftIndent := fyne.NewPos(centerX+crownWidth/3, centerY)
	rightIndent := fyne.NewPos(centerX-crownWidth/3, centerY)

	lines := []canvas.Line{
		{Position1: leftBottomCorner, Position2: leftSpike},
		{Position1: leftSpike, Position2: leftIndent},
		{Position1: leftIndent, Position2: centerSpike},
		{Position1: centerSpike, Position2: rightIndent},
		{Position1: rightIndent, Position2: rightSpike},
		{Position1: rightSpike, Position2: rightBottomCorner},
		{Position1: rightBottomCorner, Position2: leftBottomCorner},
	}

	// Draw lines connecting the points
	for _, line := range lines {
		line.StrokeWidth = crownStrokeWidth
		line.StrokeColor = decodeColor(crownColor)
		board.Add(&line)
	}

	// Draw jewels on the spikes
	jewelRadius := crownHeight * 0.16
	jewels := []canvas.Circle{
		{Position1: leftSpike, Position2: leftSpike},
		{Position1: rightSpike, Position2: rightSpike},
		{Position1: centerSpike, Position2: centerSpike},
	}
	for _, jewel := range jewels {
		jewel.FillColor = decodeColor(crownColor)
		jewel.Resize(fyne.NewSize(jewelRadius*2, jewelRadius*2))
		jewel.Move(fyne.NewPos(jewel.Position1.X-jewelRadius, jewel.Position1.Y-jewelRadius))
		board.Add(&jewel)
	}
}

func decodeColor(hex uint32) color.Color {
	return color.RGBA{
		R: uint8(hex >> 24),
		G: uint8((hex >> 16) & 0xFF),
		B: uint8((hex >> 8) & 0xFF),
		A: uint8(hex & 0xFF),
	}
}
