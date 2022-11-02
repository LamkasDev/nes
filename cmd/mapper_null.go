package main

func CreateNullMapper(start NesPointer, end NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypeNull, start, end)
	section.Read = func(nes *Nes, address NesPointer) byte {
		LogWarnLn("illegal read from null (addr: %04x)", address)
		return 0
	}
	section.Write = func(nes *Nes, address NesPointer, data []byte) {
		LogWarnLn("illegal write to null (addr: %04x)", address)
	}

	return section
}

func CreateNullSilentMapper(start NesPointer, end NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypeNull, start, end)
	section.Read = func(nes *Nes, address NesPointer) byte { return 0 }
	section.Write = func(nes *Nes, address NesPointer, data []byte) {}

	return section
}

func CreateNullSilentPointMapper(start NesPointer) NesMapperSection {
	return CreateNullSilentMapper(start, start)
}
