package main

func CreatePPUMapperSectionGroup(nes *Nes) NesMapperSectionGroup {
	points := make(map[NesPointer]NesMapperSection)
	sections := make(map[NesMapperSectionType]NesMapperSection)

	sections[NesMapperSectionTypePPUPaletteMirror] = CreatePPUPalleteMirrorMapper()
	sections[NesMapperSectionTypePPUGraphics] = CreatePPUGraphicsMapper(0x0000, 0x1fff)
	sections[NesMapperSectionTypePPUVram] = CreatePPUVramMapper(0x2000, 0x2fff)
	sections[NesMapperSectionTypeNull] = CreateNullMapper(0x3000, 0x3eff)
	sections[NesMapperSectionTypePPUPalette] = CreatePPUPalleteMapper(0x3f00, 0x3fff)

	return NesMapperSectionGroup{Points: points, Sections: sections}
}
