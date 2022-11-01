package main

func CreatePPUDataMapper(start NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypePPUData, start, start)
	section.Read = func(nes *Nes, address NesPointer) byte { return 0 }
	section.Write = func(nes *Nes, address NesPointer, data []byte) { PPUWrite(nes, data) }

	return section
}
