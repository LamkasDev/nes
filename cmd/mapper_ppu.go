package main

func CreatePPUMapper(start NesPointer, end NesPointer, mirror NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypePPU, start, end)
	section.GetAddress = func(nes *Nes, address NesPointer) NesPointer { return address % mirror }
	section.Read = func(nes *Nes, address NesPointer) byte { return BusMemoryRead(nes, address) }
	section.Write = func(nes *Nes, address NesPointer, data []byte) { BusMemoryWrite(nes, address, data) }

	return section
}
