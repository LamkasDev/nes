package main

const NesPPUOamSize = 256

func CreatePPUOamDataMapper(start NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypePPUOamData, start, start)
	section.Initialize = func(nes *Nes) { nes.PPU.OAM.Full = make([]byte, NesPPUOamSize) }
	section.Read = func(nes *Nes, address NesPointer) byte { return PPUReadOamData(nes) }
	section.Write = func(nes *Nes, address NesPointer, data []byte) { PPUWriteOamData(nes, data) }

	return section
}
