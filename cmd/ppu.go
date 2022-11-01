package main

const NesNametableSize = 0x400
const NesNametableSizeDouble = NesNametableSize * 2
const NesPPUScanlineCycles = 341
const NesPPUScanlineVBlank = 241
const NesPPUScanlineMax = 262

const NesPPUMirrorPalette = 0x10
const NesPPUMirrorVRAM = 0b10111111111111

type NesPPU struct {
	Pallete NesMemory
	VRAM    NesMemory
	OAM     NesMemory
	Regs    NesPPURegisters
	Buffer  uint8

	Scanline     uint16
	Cycles       uint32
	NMIInterrupt bool

	Mapper NesMapperSectionGroup
}

type NesPPURegisters struct {
	Control    NesControlRegister
	Mask       NesMaskRegister
	Status     NesStatusRegister
	OAMAddress NesOAMAddressRegister
	Address    NesAddressRegister
	Scroll     NesScrollRegister
}

func SetupPPU(nes *Nes) {
	nes.PPU = NesPPU{
		Regs: NesPPURegisters{
			Address: AddressRegCreate(),
		},
		Mapper: CreatePPUMapperSectionGroup(nes),
	}
	InitializeMapperSectionGroup(nes, &nes.PPU.Mapper)
}

func PPURead(nes *Nes) uint8 {
	address := AddressRegGet(&nes.PPU.Regs.Address)
	PPUIncrementVRAMAddress(nes)
	return ReadMapperSectionGroup(nes, &nes.PPU.Mapper, address)
}

func PPUWrite(nes *Nes, data []byte) {
	address := AddressRegGet(&nes.PPU.Regs.Address)
	WriteMapperSectionGroup(nes, &nes.PPU.Mapper, address, data)
	PPUIncrementVRAMAddress(nes)
}

func PPUWriteControl(nes *Nes, data NesControlRegister) {
	nmi := ControlRegGenerateVBlankNMI(nes)
	nes.PPU.Regs.Control = data
	if !nmi && ControlRegGenerateVBlankNMI(nes) && StatusRegHasVBlankStart(nes) {
		nes.PPU.NMIInterrupt = true
	}
}

func PPUWriteMask(nes *Nes, data NesMaskRegister) {
	nes.PPU.Regs.Mask = data
}

func PPUReadStatus(nes *Nes) NesStatusRegister {
	data := nes.PPU.Regs.Status
	StatusRegClearVBlankStart(nes)
	AddressRegReset(&nes.PPU.Regs.Address)
	ScrollRegReset(nes)
	return data
}

func PPUWriteOAMAddress(nes *Nes, data uint8) {
	nes.PPU.Regs.OAMAddress = NesOAMAddressRegister(data)
}

func PPUReadOamData(nes *Nes) byte {
	return MemoryRead(&nes.PPU.OAM, NesPointer(nes.PPU.Regs.OAMAddress))
}

func PPUWriteOamData(nes *Nes, data []byte) {
	for _, b := range data {
		MemoryWrite(&nes.PPU.OAM, NesPointer(nes.PPU.Regs.OAMAddress), []byte{b})
		PPUWriteOAMAddress(nes, WrappingAdd8(uint8(nes.PPU.Regs.OAMAddress), 1))
	}
}

func PPUWriteScroll(nes *Nes, data uint8) {
	ScrollRegWrite(nes, data)
}

func PPUWriteAddress(nes *Nes, data uint8) {
	AddressRegUpdate(&nes.PPU.Regs.Address, data)
}

func PPUIncrementVRAMAddress(nes *Nes) {
	AddressRegIncrement(&nes.PPU.Regs.Address, ControlRegVRAMAddrIncrement(nes))
}

func PPUPollNMIInterrupt(nes *Nes) bool {
	v := nes.PPU.NMIInterrupt
	if v {
		nes.PPU.NMIInterrupt = false
	}
	return v
}

func PPUTick(nes *Nes, cycles uint32) bool {
	nes.PPU.Cycles += cycles
	if nes.PPU.Cycles >= NesPPUScanlineCycles {
		nes.PPU.Cycles -= NesPPUScanlineCycles
		nes.PPU.Scanline++
		if nes.PPU.Scanline == NesPPUScanlineVBlank {
			StatusRegSetVBlankStart(nes)
			StatusRegClearSpriteZeroHit(nes)
			if ControlRegGenerateVBlankNMI(nes) {
				nes.PPU.NMIInterrupt = true
			}
		}
		if nes.PPU.Scanline == NesPPUScanlineMax {
			nes.PPU.Scanline = 0
			nes.PPU.NMIInterrupt = false
			StatusRegClearSpriteZeroHit(nes)
			StatusRegClearVBlankStart(nes)
			return true
		}
	}

	return false
}
