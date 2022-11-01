package main

func CreatePPUOamDmaMapper(start NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypePPUOamDma, start, start)
	section.Read = func(nes *Nes, address NesPointer) byte { return 0 }
	section.Write = func(nes *Nes, address NesPointer, data []byte) {
		buffer := [256]byte{}
		high := NesPointer(data[0]) << 8
		for i := NesPointer(0); i < 256; i++ {
			buffer[i] = BusMemoryRead(nes, high+i)
		}

		PPUWriteOamData(nes, buffer[:])
	}

	return section
}
