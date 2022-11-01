package main

// https://www.nesdev.org/wiki/PPU_registers
const (
	NesControlRegisterNametable1            = NesControlRegister(0b0000_0001)
	NesControlRegisterNametable2            = NesControlRegister(0b0000_0010)
	NesControlRegisterVRAMAddIncrement      = NesControlRegister(0b0000_0100)
	NesControlRegisterSpritePatternAddr     = NesControlRegister(0b0000_1000)
	NesControlRegisterBackgroundPatternAddr = NesControlRegister(0b0001_0000)
	NesControlRegisterSpriteSize            = NesControlRegister(0b0010_0000)
	NesControlRegisterMasterSlaveSelect     = NesControlRegister(0b0100_0000)
	NesControlRegisterGenerateNMI           = NesControlRegister(0b1000_0000)
)

type NesControlRegister uint8

func ControlRegVRAMAddrIncrement(nes *Nes) uint8 {
	if BitflagHas(uint8(nes.PPU.Regs.Control), uint8(NesControlRegisterVRAMAddIncrement)) {
		return 32
	}
	return 1
}

func ControlRegSpritePatternAddr(nes *Nes) NesPointer {
	if BitflagHas(uint8(nes.PPU.Regs.Control), uint8(NesControlRegisterSpritePatternAddr)) {
		return 0x1000
	}
	return 0
}

func ControlRegBackgroundPatternAddr(nes *Nes) NesPointer {
	if BitflagHas(uint8(nes.PPU.Regs.Control), uint8(NesControlRegisterBackgroundPatternAddr)) {
		return 0x1000
	}
	return 0
}

func ControlRegSpriteSize(nes *Nes) uint8 {
	if BitflagHas(uint8(nes.PPU.Regs.Control), uint8(NesControlRegisterSpriteSize)) {
		return 16
	}
	return 8
}

func ControlRegMasterSlaveSelect(nes *Nes) uint8 {
	if BitflagHas(uint8(nes.PPU.Regs.Control), uint8(NesControlRegisterMasterSlaveSelect)) {
		return 1
	}
	return 0
}

func ControlRegGenerateVBlankNMI(nes *Nes) bool {
	return BitflagHas(uint8(nes.PPU.Regs.Control), uint8(NesControlRegisterGenerateNMI))
}
