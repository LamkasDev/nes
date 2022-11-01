package main

func CreatePPUControlMapper(start NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypePPUControl, start, start)
	section.Read = func(nes *Nes, address NesPointer) byte { return 0 }
	section.Write = func(nes *Nes, address NesPointer, data []byte) { PPUWriteControl(nes, NesControlRegister(data[0])) }

	return section
}
