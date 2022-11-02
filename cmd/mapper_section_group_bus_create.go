package main

func CreateBusMapperSectionGroup(nes *Nes) NesMapperSectionGroup {
	points := make(map[NesPointer]NesMapperSection)
	sections := make(map[NesMapperSectionType]NesMapperSection)

	points[0x2000] = CreatePPUControlMapper(0x2000)
	points[0x2001] = CreatePPUMaskMapper(0x2001)
	points[0x2002] = CreatePPUStatusMapper(0x2002)
	points[0x2003] = CreatePPUOamAddressMapper(0x2003)
	points[0x2004] = CreatePPUOamDataMapper(0x2004)
	points[0x2005] = CreatePPUScrollMapper(0x2005)
	points[0x2006] = CreatePPUAddressMapper(0x2006)
	points[0x2007] = CreatePPUDataMapper(0x2007)
	sections[NesMapperSectionTypePPU] = CreatePPUMapper(0x2008, 0x3fff, 0x2007)
	points[0x4014] = CreatePPUOamDmaMapper(0x4014)
	sections[NesMapperSectionTypeAPU] = CreateAPUMapper(0x4000, 0x4015)
	points[0x4016] = CreateJoypadMapper(NesMapperSectionTypeJoypad, 0x4016, 0)
	points[0x4017] = CreateNullSilentPointMapper(0x4017)

	switch nes.Bus.ROM.Mapper {
	case 0:
		sections[NesMapperSectionTypeRAM] = CreateRamMapperMirror(0x0000, 0x1fff, 0x7ff)
		sections[NesMapperSectionTypeROM] = CreateRomMapperMirror(0x8000, 0xffff, 0x4000)
	case 3:
		sections[NesMapperSectionTypeGraphics] = CreateGraphicsMapper(0x0000, 0x1fff)
		sections[NesMapperSectionTypeROM] = CreateRomMapper(0x5000, 0x7fff)
		sections[NesMapperSectionTypeGraphicsBankSelect] = CreateGraphicsBankSelectMapper(0x8000, 0xffff)
	}

	return NesMapperSectionGroup{Points: points, Sections: sections}
}
