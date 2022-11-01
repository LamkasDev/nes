package main

type NesScrollRegister struct {
	X     uint8
	Y     uint8
	Latch bool
}

func ScrollRegWrite(nes *Nes, data uint8) {
	if !nes.PPU.Regs.Scroll.Latch {
		nes.PPU.Regs.Scroll.X = data
	} else {
		nes.PPU.Regs.Scroll.Y = data
	}
	nes.PPU.Regs.Scroll.Latch = !nes.PPU.Regs.Scroll.Latch
}

func ScrollRegReset(nes *Nes) {
	nes.PPU.Regs.Scroll.Latch = false
}
