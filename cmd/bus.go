package main

type NesBus struct {
	RAM    NesMemory
	ROM    NesRom
	Cycles uint32

	Mapper     NesMapperSectionGroup
	MapperData NesMapperBusData
}

func SetupBus(nes *Nes) {
	nes.Bus.Mapper = CreateBusMapperSectionGroup(nes)
	InitializeMapperSectionGroup(nes, &nes.Bus.Mapper)
}

func BusMemoryRead(nes *Nes, address NesPointer) byte {
	return ReadMapperSectionGroup(nes, &nes.Bus.Mapper, address)
}

func BusMemoryWrite(nes *Nes, address NesPointer, data []byte) {
	WriteMapperSectionGroup(nes, &nes.Bus.Mapper, address, data)
}

func BusMemoryReadAddress(nes *Nes, address NesPointer) NesPointer {
	low := BusMemoryRead(nes, address)
	high := BusMemoryRead(nes, address+1)
	return (NesPointer(high) << 8) | NesPointer(low)
}

func BusMemoryWriteAddress(nes *Nes, address NesPointer, data NesPointer) {
	low := uint8(data & 0xff)
	high := uint8(data >> 8)
	BusMemoryWrite(nes, address, []byte{low})
	BusMemoryWrite(nes, address+1, []byte{high})
}

func BusTick(nes *Nes, cycles uint32) {
	nes.Bus.Cycles += cycles
	frame := PPUTick(nes, cycles*3)
	if frame {
		RenderFrame(nes)
	}
}
