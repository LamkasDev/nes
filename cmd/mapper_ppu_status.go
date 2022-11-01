package main

func CreatePPUStatusMapper(start NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypePPUStatus, start, start)
	section.Read = func(nes *Nes, address NesPointer) byte { return byte(PPUReadStatus(nes)) }
	section.Write = func(nes *Nes, address NesPointer, data []byte) {
		LogWarnLn("illegal write to ppu status (b: %02x)", data[0])
	}

	return section
}
