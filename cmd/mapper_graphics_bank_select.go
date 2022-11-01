package main

func CreateGraphicsBankSelectMapper(start NesPointer, end NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypeGraphicsBankSelect, start, end)
	section.Read = func(nes *Nes, address NesPointer) byte {
		return uint8(nes.Bus.MapperData.GraphicsBank / 8192)
	}
	section.Write = func(nes *Nes, address NesPointer, data []byte) {
		nes.Bus.MapperData.GraphicsBank = (address - start) * 8192
	}

	return section
}
