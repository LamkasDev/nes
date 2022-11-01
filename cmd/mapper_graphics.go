package main

func CreateGraphicsMapper(start NesPointer, end NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypeGraphics, start, end)
	section.GetAddress = func(nes *Nes, address NesPointer) NesPointer { return address - start }
	section.Read = func(nes *Nes, address NesPointer) byte {
		return MemoryRead(&nes.Bus.ROM.Graphics, nes.Bus.MapperData.GraphicsBank+address)
	}
	section.Write = func(nes *Nes, address NesPointer, data []byte) {}

	return section
}
