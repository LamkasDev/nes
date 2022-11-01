package main

func CreatePPUPalleteMirrorMapper() NesMapperSection {
	section := CreateMapperSection(NesMapperSectionTypePPUPaletteMirror, 0x3f10, 0x3f1c)
	section.Matches = func(nes *Nes, address NesPointer) bool {
		return address == 0x3f10 || address == 0x3f14 || address == 0x3f18 || address == 0x3f1c
	}
	section.GetAddress = func(nes *Nes, address NesPointer) NesPointer {
		return address - 0x3f00 - 0x10
	}
	section.Read = func(nes *Nes, address NesPointer) byte {
		return MemoryRead(&nes.PPU.Pallete, address)
	}
	section.Write = func(nes *Nes, address NesPointer, data []byte) {
		MemoryWrite(&nes.PPU.Pallete, address, data)
	}

	return section
}
