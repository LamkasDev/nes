package main

func ShowTileBank(nes *Nes, bank NesPointer) {
	tc := 0
	tr := 0

	for i := NesPointer(0); i < 511; i++ {
		if i != 0 && i%25 == 0 {
			tr++
			tc = 0
		}
		ts := bank + (i * 16)
		tile := MemoryReadLen(&nes.Bus.ROM.Graphics, ts, 16)

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
		tc++
	}
}
