package main

import "github.com/veandco/go-sdl2/sdl"

// https://www.nesdev.org/wiki/PPU_registers
const (
	NesMaskRegisterGrayscale      = NesControlRegister(0b0000_0001)
	NesMaskRegisterLeftBackground = NesControlRegister(0b0000_0010)
	NesMaskRegisterLeftSprite     = NesControlRegister(0b0000_0100)
	NesMaskRegisterShowBackground = NesControlRegister(0b0000_1000)
	NesMaskRegisterShowSprites    = NesControlRegister(0b0001_0000)
	NesMaskRegisterEmphasiseRed   = NesControlRegister(0b0010_0000)
	NesMaskRegisterEmphasiseGreen = NesControlRegister(0b0100_0000)
	NesMaskRegisterEmphasiseBlue  = NesControlRegister(0b1000_0000)
)

type NesMaskRegister uint8

func MaskRegHasGrayscale(nes *Nes) bool {
	return BitflagHas(uint8(nes.PPU.Regs.Mask), uint8(NesMaskRegisterGrayscale))
}

func MaskRegHasLeftBackground(nes *Nes) bool {
	return BitflagHas(uint8(nes.PPU.Regs.Mask), uint8(NesMaskRegisterLeftBackground))
}

func MaskRegHasLeftSprite(nes *Nes) bool {
	return BitflagHas(uint8(nes.PPU.Regs.Mask), uint8(NesMaskRegisterLeftSprite))
}

func MaskRegHasShowBackground(nes *Nes) bool {
	return BitflagHas(uint8(nes.PPU.Regs.Mask), uint8(NesMaskRegisterShowBackground))
}

func MaskRegHasShowSprites(nes *Nes) bool {
	return BitflagHas(uint8(nes.PPU.Regs.Mask), uint8(NesMaskRegisterShowSprites))
}

func MaskRegEmphasise(nes *Nes) []sdl.Color {
	c := []sdl.Color{}
	if BitflagHas(uint8(nes.PPU.Regs.Mask), uint8(NesMaskRegisterEmphasiseRed)) {
		c = append(c, ColorRed)
	}
	if BitflagHas(uint8(nes.PPU.Regs.Mask), uint8(NesMaskRegisterEmphasiseGreen)) {
		c = append(c, ColorGreen)
	}
	if BitflagHas(uint8(nes.PPU.Regs.Mask), uint8(NesMaskRegisterEmphasiseBlue)) {
		c = append(c, ColorBlue)
	}

	return c
}
