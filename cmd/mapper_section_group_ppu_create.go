package main

func CreatePPUMapperSectionGroup(nes *Nes) NesMapperSectionGroup {
	group := NesMapperSectionGroup{
		Points:   make(map[NesPointer]NesMapperSection),
		Sections: []NesMapperSection{},
	}

	group.Sections = append(group.Sections, CreatePPUPalleteMirrorMapper())
	group.Sections = append(group.Sections, CreatePPUGraphicsMapper(0x0000, 0x1fff))
	group.Sections = append(group.Sections, CreatePPUVramMapper(0x2000, 0x2fff))
	group.Sections = append(group.Sections, CreateNullMapper(0x3000, 0x3eff))
	group.Sections = append(group.Sections, CreatePPUPalleteMapper(0x3f00, 0x3fff))
	group.VRAMStart = 0x2000

	return group
}
