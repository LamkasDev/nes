package main

const NesPPUPalleteSize = 32

func CreatePPUPalleteMapper(start NesPointer, end NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypePPUPalette, start, end)
	section.Initialize = func(nes *Nes) { nes.PPU.Pallete.Full = make([]byte, NesPPUPalleteSize) }
	section.GetAddress = func(nes *Nes, address NesPointer) NesPointer {
		return address - start
	}
	section.Read = func(nes *Nes, address NesPointer) byte {
		return MemoryRead(&nes.PPU.Pallete, address)
	}
	section.Write = func(nes *Nes, address NesPointer, data []byte) {
		MemoryWrite(&nes.PPU.Pallete, address, data)
	}

	return section
}
