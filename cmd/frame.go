package main

const NesFrameWidth = 256
const NesFrameHeight = 256
const NesFrameSize = NesFrameWidth * NesFrameHeight

type NesFrame struct {
	Data [NesFrameSize]uint32
}

func SetFramePixel(frame *NesFrame, x uint32, y uint32, c uint32) {
	base := y*NesFrameWidth + x
	if base < NesFrameSize {
		frame.Data[base] = c
	}
}

func RenderFrame(nes *Nes) {
	nes.Locks.Renderer.Lock()
	defer nes.Locks.Renderer.Unlock()

	bank := ControlRegBackgroundPatternAddr(nes)
	for i := NesPointer(0); i < NesNametableFirst; i++ {
		ti := NesPointer(MemoryRead(&nes.PPU.VRAM, i))
		tc := i % NesNametableTilesWidth
		tr := i / NesNametableTilesHeight
		ts := bank + (ti * NesTileSize)
		tile := MemoryReadLen(&nes.Bus.ROM.Graphics, ts, NesTileSize)
		pallete := GetBackgroundPalette(nes, uint16(tc), uint16(tr))

		for y := 0; y < NesTileHeight; y++ {
			high := tile[y]
			low := tile[y+NesTileHeight]
			for x := (NesTileWidth - 1); x >= 0; x-- {
				v := (1&low)<<1 | (1 & high)
				high >>= 1
				low >>= 1
				c := uint32(0)
				switch v {
				case 0:
					c = NesPPUPallete[nes.PPU.Pallete.Full[0]]
				case 1, 2, 3:
					c = NesPPUPallete[pallete[v]]
				}
				px := uint32(tc*NesTileWidth) + uint32(x)
				py := uint32(tr*NesTileHeight) + uint32(y)
				SetFramePixel(&RendererFrame, px, py, c)
			}
		}
	}

	bank = ControlRegSpritePatternAddr(nes)
	for i := len(nes.PPU.OAM.Full) - 4; i >= 0; i -= 4 {
		ti := NesPointer(MemoryRead(&nes.PPU.OAM, NesPointer(i+1)))
		tx := MemoryRead(&nes.PPU.OAM, NesPointer(i+3))
		ty := MemoryRead(&nes.PPU.OAM, NesPointer(i))
		ts := bank + (ti * NesTileSize)
		tile := MemoryReadLen(&nes.Bus.ROM.Graphics, ts, NesTileSize)

		spByte := MemoryRead(&nes.PPU.OAM, NesPointer(i+2))
		flipVertical, flipHorizontal := false, false
		if spByte>>7&1 == 1 {
			flipVertical = true
		}
		if spByte>>6&1 == 1 {
			flipHorizontal = true
		}
		palleteIndex := spByte & 0b11
		pallete := GetSpritePalette(nes, palleteIndex)

		for y := 0; y < NesTileHeight; y++ {
			high := tile[y]
			low := tile[y+NesTileHeight]
			for x := (NesTileWidth - 1); x >= 0; x-- {
				v := (1&low)<<1 | (1 & high)
				high >>= 1
				low >>= 1
				c := uint32(0)
				switch v {
				case 0:
					continue
				case 1, 2, 3:
					c = NesPPUPallete[pallete[v]]
				}

				px, py := uint32(0), uint32(0)
				if flipHorizontal {
					px = (uint32(tx) + 7) - uint32(x)
				} else {
					px = uint32(tx) + uint32(x)
				}
				if flipVertical {
					py = (uint32(ty) + 7) - uint32(y)
				} else {
					py = uint32(ty) + uint32(y)
				}
				SetFramePixel(&RendererFrame, px, py, c)
			}
		}
	}
}
