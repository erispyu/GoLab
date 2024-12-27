package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const boardSize = 8 // Checkers board size is 8x8

func main() {
	// Create application instance
	myApp := app.New()
	myWindow := myApp.NewWindow("Checkers")

	// Create a grid for the board
	grid := container.NewGridWithColumns(boardSize)

	// Generate the board
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			// Determine the color of the square
			var squareColor color.Color
			if (row+col)%2 == 0 {
				squareColor = color.RGBA{R: 240, G: 217, B: 181, A: 255} // Light color
			} else {
				squareColor = color.RGBA{R: 181, G: 136, B: 99, A: 255} // Dark color
			}

			// Draw the square
			rect := canvas.NewRectangle(squareColor)
			rect.SetMinSize(fyne.NewSize(50, 50)) // Size of each square
			grid.Add(rect)
		}
	}

	// Add the board to the window
	content := container.NewVBox(
		widget.NewLabel("Checkers"),
		grid,
	)
	myWindow.SetContent(content)

	// Set the window size and display it
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}
