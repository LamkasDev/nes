package main

const NesPPUVramSize = 2048

func CreatePPUVramMapper(start NesPointer, end NesPointer) NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypePPUVram, start, end)
	section.Initialize = func(nes *Nes) { nes.PPU.VRAM.Full = make([]byte, NesPPUVramSize) }
	section.GetAddress = func(nes *Nes, address NesPointer) NesPointer {
		return MirrorVRAMAddress(nes, address)
	}
	section.Read = func(nes *Nes, address NesPointer) byte {
		data := nes.PPU.Buffer
		nes.PPU.Buffer = MemoryRead(&nes.PPU.VRAM, address)
		return data
	}
	section.Write = func(nes *Nes, address NesPointer, data []byte) {
		MemoryWrite(&nes.PPU.VRAM, address, data)
	}

	return section
}
