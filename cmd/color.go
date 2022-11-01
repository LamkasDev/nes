package main

import "github.com/veandco/go-sdl2/sdl"

const ColorsSize = 15

var ColorBlack = sdl.Color{R: 0, G: 0, B: 0, A: 255}
var ColorWhite = sdl.Color{R: 255, G: 255, B: 255, A: 255}
var ColorGray = sdl.Color{R: 128, G: 128, B: 128, A: 255}
var ColorRed = sdl.Color{R: 255, G: 0, B: 0, A: 255}
var ColorGreen = sdl.Color{R: 0, G: 255, B: 0, A: 255}
var ColorBlue = sdl.Color{R: 0, G: 0, B: 255, A: 255}
var ColorMagenta = sdl.Color{R: 255, G: 0, B: 255, A: 255}
var ColorYellow = sdl.Color{R: 255, G: 255, B: 0, A: 255}
var ColorCyan = sdl.Color{R: 0, G: 255, B: 255, A: 255}

var Colors = [ColorsSize]sdl.Color{
	ColorBlack,
	ColorWhite,
	ColorGray,
	ColorRed,
	ColorGreen,
	ColorBlue,
	ColorMagenta,
	ColorYellow,
	ColorBlack, // ??
	ColorGray,
	ColorRed,
	ColorGreen,
	ColorBlue,
	ColorMagenta,
	ColorYellow,
}

func GetSDLColor(c uint8) sdl.Color {
	if c >= ColorsSize {
		return ColorCyan
	}

	return Colors[c]
}
