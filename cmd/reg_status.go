package main

// https://www.nesdev.org/wiki/PPU_registers
const (
	NesStatusRegisterNotUsed        = NesControlRegister(0b0000_0001)
	NesStatusRegisterNotUsed2       = NesControlRegister(0b0000_0010)
	NesStatusRegisterNotUsed3       = NesControlRegister(0b0000_0100)
	NesStatusRegisterNotUsed4       = NesControlRegister(0b0000_1000)
	NesStatusRegisterNotUsed5       = NesControlRegister(0b0001_0000)
	NesStatusRegisterSpriteOverflow = NesControlRegister(0b0010_0000)
	NesStatusRegisterSpriteZeroHit  = NesControlRegister(0b0100_0000)
	NesStatusRegisterVBlankStarted  = NesControlRegister(0b1000_0000)
)

type NesStatusRegister uint8

func StatusRegSetSpriteOverflow(nes *Nes) {
	nes.PPU.Regs.Status = NesStatusRegister(BitflagSet(uint8(nes.PPU.Regs.Status), uint8(NesStatusRegisterSpriteOverflow)))
}

func StatusRegSetSpriteZeroHit(nes *Nes) {
	nes.PPU.Regs.Status = NesStatusRegister(BitflagSet(uint8(nes.PPU.Regs.Status), uint8(NesStatusRegisterSpriteZeroHit)))
}
func StatusRegClearSpriteZeroHit(nes *Nes) {
	nes.PPU.Regs.Status = NesStatusRegister(BitflagClear(uint8(nes.PPU.Regs.Status), uint8(NesStatusRegisterSpriteZeroHit)))
}

func StatusRegSetVBlankStart(nes *Nes) {
	nes.PPU.Regs.Status = NesStatusRegister(BitflagSet(uint8(nes.PPU.Regs.Status), uint8(NesStatusRegisterVBlankStarted)))
}
func StatusRegClearVBlankStart(nes *Nes) {
	nes.PPU.Regs.Status = NesStatusRegister(BitflagClear(uint8(nes.PPU.Regs.Status), uint8(NesStatusRegisterVBlankStarted)))
}

func StatusRegHasVBlankStart(nes *Nes) bool {
	return BitflagHas(uint8(nes.PPU.Regs.Status), uint8(NesStatusRegisterVBlankStarted))
}
