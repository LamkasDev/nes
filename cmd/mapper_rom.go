package main

func CreateRomMapper(start NesPointer, end NesPointer) NesMapperSection {
	return CreateRomMapperRaw(start, end, func(nes *Nes, address NesPointer) NesPointer {
		return address - start
	})
}

func CreateRomMapperMirror(start NesPointer, end NesPointer, mirror NesPointer) NesMapperSection {
	return CreateRomMapperRaw(start, end, func(nes *Nes, address NesPointer) NesPointer {
		return (address - start) % mirror
	})
}

func CreateRomMapperRaw(start NesPointer, end NesPointer, getAddress NesMapperSectionGetAddress) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypeROM, start, end)
	section.GetAddress = getAddress
	section.Read = func(nes *Nes, address NesPointer) byte { return MemoryRead(&nes.Bus.ROM.Program, address) }
	section.Write = func(nes *Nes, address NesPointer, data []byte) {
		LogWarnLn("illegal write to rom (addr: %04x)", address)
	}

	return section
}
