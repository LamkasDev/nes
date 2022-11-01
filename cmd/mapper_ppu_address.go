package main

func CreatePPUAddressMapper(start NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypePPUAddress, start, start)
	section.Read = func(nes *Nes, address NesPointer) byte { return 0 }
	section.Write = func(nes *Nes, address NesPointer, data []byte) { PPUWriteAddress(nes, uint8(data[0])) }

	return section
}
