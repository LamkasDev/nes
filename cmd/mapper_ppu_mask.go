package main

func CreatePPUMaskMapper(start NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypePPUMask, start, start)
	section.Read = func(nes *Nes, address NesPointer) byte { return 0 }
	section.Write = func(nes *Nes, address NesPointer, data []byte) { PPUWriteMask(nes, NesMaskRegister(data[0])) }

	return section
}
