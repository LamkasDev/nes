package main

func CreatePPUOamAddressMapper(start NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypePPUOamAddress, start, start)
	section.Read = func(nes *Nes, address NesPointer) byte { return 0 }
	section.Write = func(nes *Nes, address NesPointer, data []byte) { PPUWriteOAMAddress(nes, uint8(data[0])) }

	return section
}
