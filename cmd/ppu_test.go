package main

import (
	"testing"
)

func TestPPUWrite(t *testing.T) {
	nes := SetupNES()
	PPUWriteAddress(&nes, 0x23)
	PPUWriteAddress(&nes, 0x05)
	PPUWrite(&nes, []byte{0x66})

	b := MemoryRead(&nes.PPU.VRAM, 0x0305)
	if b != 0x66 {
		t.Fatalf("Failed write of %v (v: %v)", 0x66, b)
	}
}

func TestPPURead(t *testing.T) {
	nes := SetupNES()
	PPUWriteControl(&nes, 0)
	MemoryWrite(&nes.PPU.VRAM, 0x0305, []byte{0x66})
	PPUWriteAddress(&nes, 0x23)
	PPUWriteAddress(&nes, 0x05)
	PPURead(&nes)

	addr := AddressRegGet(&nes.PPU.Regs.Address)
	if addr != 0x2306 {
		t.Fatalf("Failed read of %v (v: %v)", 0x2306, addr)
	}
	data := PPURead(&nes)
	if data != 0x66 {
		t.Fatalf("Failed read of %v (v: %v)", 0x66, data)
	}
}

func TestPPUReadCrossPage(t *testing.T) {
	nes := SetupNES()
	PPUWriteControl(&nes, 0)
	MemoryWrite(&nes.PPU.VRAM, 0x01ff, []byte{0x66})
	MemoryWrite(&nes.PPU.VRAM, 0x0200, []byte{0x77})
	PPUWriteAddress(&nes, 0x21)
	PPUWriteAddress(&nes, 0xff)
	PPURead(&nes)

	a := PPURead(&nes)
	if a != 0x66 {
		t.Fatalf("Failed read of %v (v: %v)", 0x66, a)
	}
	b := PPURead(&nes)
	if b != 0x77 {
		t.Fatalf("Failed read of %v (v: %v)", 0x77, b)
	}
}

func TestPPUReadStep32(t *testing.T) {
	nes := SetupNES()
	PPUWriteControl(&nes, 0b100)
	MemoryWrite(&nes.PPU.VRAM, 0x01ff, []byte{0x66})
	MemoryWrite(&nes.PPU.VRAM, 0x01ff+32, []byte{0x77})
	MemoryWrite(&nes.PPU.VRAM, 0x01ff+64, []byte{0x88})
	PPUWriteAddress(&nes, 0x21)
	PPUWriteAddress(&nes, 0xff)
	PPURead(&nes)

	a := PPURead(&nes)
	if a != 0x66 {
		t.Fatalf("Failed read of %v (v: %v)", 0x66, a)
	}
	b := PPURead(&nes)
	if b != 0x77 {
		t.Fatalf("Failed read of %v (v: %v)", 0x77, b)
	}
	c := PPURead(&nes)
	if c != 0x88 {
		t.Fatalf("Failed read of %v (v: %v)", 0x88, c)
	}
}

func TestPPUHorizonal(t *testing.T) {
	nes := SetupNES()
	nes.Bus.ROM.Mirroring = NesMirroringHorizontal
	PPUWriteAddress(&nes, 0x24)
	PPUWriteAddress(&nes, 0x05)
	PPUWrite(&nes, []byte{0x66})

	PPUWriteAddress(&nes, 0x28)
	PPUWriteAddress(&nes, 0x05)
	PPUWrite(&nes, []byte{0x77})

	PPUWriteAddress(&nes, 0x20)
	PPUWriteAddress(&nes, 0x05)
	PPURead(&nes)
	a := PPURead(&nes)
	if a != 0x66 {
		t.Fatalf("Failed read of %v (v: %v)", 0x66, a)
	}

	PPUWriteAddress(&nes, 0x2c)
	PPUWriteAddress(&nes, 0x05)
	PPURead(&nes)
	b := PPURead(&nes)
	if b != 0x77 {
		t.Fatalf("Failed read of %v (v: %v)", 0x77, b)
	}
}

func TestPPUVertical(t *testing.T) {
	nes := SetupNES()
	nes.Bus.ROM.Mirroring = NesMirroringVertical
	PPUWriteAddress(&nes, 0x20)
	PPUWriteAddress(&nes, 0x05)
	PPUWrite(&nes, []byte{0x66})

	PPUWriteAddress(&nes, 0x2c)
	PPUWriteAddress(&nes, 0x05)
	PPUWrite(&nes, []byte{0x77})

	PPUWriteAddress(&nes, 0x28)
	PPUWriteAddress(&nes, 0x05)
	PPURead(&nes)
	a := PPURead(&nes)
	if a != 0x66 {
		t.Fatalf("Failed read of %v (v: %v)", 0x66, a)
	}

	PPUWriteAddress(&nes, 0x24)
	PPUWriteAddress(&nes, 0x05)
	PPURead(&nes)
	b := PPURead(&nes)
	if b != 0x77 {
		t.Fatalf("Failed read of %v (v: %v)", 0x77, b)
	}
}

func TestPPUReset(t *testing.T) {
	nes := SetupNES()
	MemoryWrite(&nes.PPU.VRAM, 0x0305, []byte{0x66})

	PPUWriteAddress(&nes, 0x21)
	PPUWriteAddress(&nes, 0x23)
	PPUWriteAddress(&nes, 0x05)

	PPURead(&nes)
	a := PPURead(&nes)
	if a != 0x66 {
		t.Fatalf("Failed read of %v (v: %v)", 0x66, a)
	}

	PPUReadStatus(&nes)
	PPUWriteAddress(&nes, 0x23)
	PPUWriteAddress(&nes, 0x05)

	PPURead(&nes)
	b := PPURead(&nes)
	if b != 0x66 {
		t.Fatalf("Failed read of %v (v: %v)", 0x66, b)
	}
}

func TestPPUVramMirroring(t *testing.T) {
	nes := SetupNES()
	PPUWriteControl(&nes, 0)
	nes.PPU.VRAM.Full[0x0305] = 0x66

	PPUWriteAddress(&nes, 0x63)
	PPUWriteAddress(&nes, 0x05)

	PPURead(&nes)
	b := PPURead(&nes)
	if b != 0x66 {
		t.Fatalf("failed read of %v (v: %v)", 0x66, b)
	}
}

func TestPPUResetVBlank(t *testing.T) {
	nes := SetupNES()
	StatusRegSetVBlankStart(&nes)

	status := PPUReadStatus(&nes)
	if status>>7 != 1 {
		t.Fatalf("failed status reset")
	}
	if nes.PPU.Regs.Status>>7 != 0 {
		t.Fatalf("failed status reset")
	}
}

func TestPPUOam(t *testing.T) {
	nes := SetupNES()
	SetupBus(&nes)
	PPUWriteOAMAddress(&nes, 0x10)
	PPUWriteOamData(&nes, []byte{0x66, 0x77})

	PPUWriteOAMAddress(&nes, 0x10)
	a := PPUReadOamData(&nes)
	if a != 0x66 {
		t.Fatalf("Failed read of %v (v: %v)", 0x66, a)
	}

	PPUWriteOAMAddress(&nes, 0x11)
	b := PPUReadOamData(&nes)
	if b != 0x77 {
		t.Fatalf("Failed read of %v (v: %v)", 0x77, b)
	}
}
