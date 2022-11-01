package main

type NesMapperSectionGroup struct {
	Points   map[NesPointer]NesMapperSection
	Sections map[NesMapperSectionType]NesMapperSection
}

func InitializeMapperSectionGroup(nes *Nes, group *NesMapperSectionGroup) {
	for _, point := range group.Points {
		if point.Initialize != nil {
			point.Initialize(nes)
		}
	}
	for _, section := range group.Sections {
		if section.Initialize != nil {
			section.Initialize(nes)
		}
	}
}

func ReadMapperSectionGroup(nes *Nes, group *NesMapperSectionGroup, address NesPointer) byte {
	point, ok := group.Points[address]
	if ok {
		return point.Read(nes, point.GetAddress(nes, address))
	}
	for _, section := range group.Sections {
		if section.Matches(nes, address) {
			return section.Read(nes, section.GetAddress(nes, address))
		}
	}

	LogWarnLn("invalid mem read at %04x", address)
	return 0
}

func WriteMapperSectionGroup(nes *Nes, group *NesMapperSectionGroup, address NesPointer, data []byte) {
	point, ok := group.Points[address]
	if ok {
		point.Write(nes, point.GetAddress(nes, address), data)
		return
	}
	for _, section := range group.Sections {
		if section.Matches(nes, address) {
			section.Write(nes, section.GetAddress(nes, address), data)
			return
		}
	}
	LogWarnLn("illegal write (addr: %04x)", address)
}
