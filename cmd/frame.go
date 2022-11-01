package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const NesFrameWidth = 256
const NesFrameHeight = 256
const NesFrameSize = NesFrameWidth * NesFrameHeight

type NesFrame struct {
	Data [NesFrameSize]uint32
}

func SetFramePixel(frame *NesFrame, x uint32, y uint32, c sdl.Color) {
	base := y*NesFrameWidth + x
	if base < NesFrameSize {
		frame.Data[base] = c.Uint32()
	}
}

func RenderFrame(nes *Nes) {
	bank := ControlRegBackgroundPatternAddr(nes)

	for i := NesPointer(0); i < NesNametableFirst; i++ {
		ti := NesPointer(MemoryRead(&nes.PPU.VRAM, i))
		tc := i % NesNametableTilesWidth
		tr := i / NesNametableTilesHeight
		ts := bank + (ti * 16)
		tile := MemoryReadLen(&nes.Bus.ROM.Graphics, ts, 16)
		// pallete := GetBackgroundPalette(nes, uint16(tc), uint16(tr))

		for y := 0; y < NesTileSize; y++ {
			high := tile[y]
			low := tile[y+NesTileSize]
			for x := (NesTileSize - 1); x >= 0; x-- {
				v := (1&high)<<1 | (1 & low)
				high >>= 1
				low >>= 1
				c := ColorRed
				switch v {
				case 0:
					c = NesPPUPallete[0x01]
				case 1:
					c = NesPPUPallete[0x23]
				case 2:
					c = NesPPUPallete[0x27]
				case 3:
					c = NesPPUPallete[0x30]
				}
				px := uint32(tc*NesTileSize) + uint32(x)
				py := uint32(tr*NesTileSize) + uint32(y)
				SetFramePixel(&RendererFrame, px, py, c)
			}
			CycleRenderer(nes)
		}
	}
}
