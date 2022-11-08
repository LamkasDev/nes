package main

func CreateRamMapper(start NesPointer, end NesPointer) NesMapperSection {
	return CreateRamMapperRaw(start, end, func(nes *Nes, address NesPointer) NesPointer {
		return address
	})
}

func CreateRamMapperMirror(start NesPointer, end NesPointer, mirror NesPointer) NesMapperSection {
	return CreateRamMapperRaw(start, end, func(nes *Nes, address NesPointer) NesPointer {
		return address % mirror
	})
}

func CreateRamMapperRaw(start NesPointer, end NesPointer, getAddress NesMapperSectionGetAddress) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypeRAM, start, end)
	section.Initialize = func(nes *Nes) { nes.Bus.RAM.Full = make([]byte, GetMapperSectionSize(section)) }
	section.GetAddress = getAddress
	section.Read = func(nes *Nes, address NesPointer) byte {
		return MemoryRead(&nes.Bus.RAM, address)
	}
	section.Write = func(nes *Nes, address NesPointer, data []byte) {
		MemoryWrite(&nes.Bus.RAM, address, data)
	}

	return section
}
