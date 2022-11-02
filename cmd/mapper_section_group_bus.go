package main

const (
	NesMapperSectionTypeRAM           = NesMapperSectionType("RAM")
	NesMapperSectionTypePPUControl    = NesMapperSectionType("PPU_CONTROL")
	NesMapperSectionTypePPUMask       = NesMapperSectionType("PPU_MASK")
	NesMapperSectionTypePPUStatus     = NesMapperSectionType("PPU_STATUS")
	NesMapperSectionTypePPUOamAddress = NesMapperSectionType("PPU_OAM_ADDRESS")
	NesMapperSectionTypePPUOamData    = NesMapperSectionType("PPU_OAM_DATA")
	NesMapperSectionTypePPUScroll     = NesMapperSectionType("PPU_SCROLL")
	NesMapperSectionTypePPUAddress    = NesMapperSectionType("PPU_ADDRESS")
	NesMapperSectionTypeJoypad        = NesMapperSectionType("JOYPAD")
	NesMapperSectionTypePPUData       = NesMapperSectionType("PPU_DATA")
	NesMapperSectionTypePPUOamDma     = NesMapperSectionType("PPU_OAM_DMA")
	NesMapperSectionTypePPU           = NesMapperSectionType("PPU")
	NesMapperSectionTypeAPU           = NesMapperSectionType("APU")
	NesMapperSectionTypeROM           = NesMapperSectionType("ROM")

	NesMapperSectionTypeGraphics           = NesMapperSectionType("GFX")
	NesMapperSectionTypeGraphicsBankSelect = NesMapperSectionType("GFX_BANK_SELECT")
)

type NesMapperBusData struct {
	GraphicsBank NesPointer
}
