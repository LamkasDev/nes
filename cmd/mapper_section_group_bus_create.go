package main

func CreateBusMapperSectionGroup(nes *Nes) NesMapperSectionGroup {
	group := NesMapperSectionGroup{
		Points:   make(map[NesPointer]NesMapperSection),
		Sections: []NesMapperSection{},
	}

	group.Points[0x2000] = CreatePPUControlMapper(0x2000)
	group.Points[0x2001] = CreatePPUMaskMapper(0x2001)
	group.Points[0x2002] = CreatePPUStatusMapper(0x2002)
	group.Points[0x2003] = CreatePPUOamAddressMapper(0x2003)
	group.Points[0x2004] = CreatePPUOamDataMapper(0x2004)
	group.Points[0x2005] = CreatePPUScrollMapper(0x2005)
	group.Points[0x2006] = CreatePPUAddressMapper(0x2006)
	group.Points[0x2007] = CreatePPUDataMapper(0x2007)
	group.Sections = append(group.Sections, CreatePPUMapper(0x2008, 0x3fff, 0x2007))
	group.Points[0x4014] = CreatePPUOamDmaMapper(0x4014)
	group.Sections = append(group.Sections, CreateAPUMapper(0x4000, 0x4015))
	group.Points[0x4016] = CreateJoypadMapper(NesMapperSectionTypeJoypad, 0x4016, 0)
	group.Points[0x4017] = CreateNullSilentPointMapper(0x4017)

	switch nes.Bus.ROM.Mapper {
	case 0:
		group.Sections = append(group.Sections, CreateRamMapperMirror(0x0000, 0x1fff, 0x7ff))
		group.Sections = append(group.Sections, CreateRomMapperMirror(0x8000, 0xffff, 0x4000))
		group.ROMStart = 0x8000
	case 3:
		group.Sections = append(group.Sections, CreateGraphicsMapper(0x0000, 0x1fff))
		group.Sections = append(group.Sections, CreateRomMapper(0x5000, 0x7fff))
		group.Sections = append(group.Sections, CreateGraphicsBankSelectMapper(0x8000, 0xffff))
		group.ROMStart = 0x5000
	}

	return group
}
