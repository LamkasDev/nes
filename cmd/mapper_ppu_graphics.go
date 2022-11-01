package main

func CreatePPUGraphicsMapper(start NesPointer, end NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypeGraphics, start, end)
	section.Read = func(nes *Nes, address NesPointer) byte {
		data := nes.PPU.Buffer
		nes.PPU.Buffer = MemoryRead(&nes.Bus.ROM.Graphics, address)
		return data
	}
	section.Write = func(nes *Nes, address NesPointer, data []byte) {
		LogWarnLn("illegal write to ppu graphics (addr: %04x)", address)
	}

	return section
}
