package main

import "github.com/veandco/go-sdl2/sdl"

const (
	NesJoypadButtonRight  = NesJoypadButton(0b1000_0000)
	NesJoypadButtonLeft   = NesJoypadButton(0b0100_0000)
	NesJoypadButtonDown   = NesJoypadButton(0b0010_0000)
	NesJoypadButtonUp     = NesJoypadButton(0b0001_0000)
	NesJoypadButtonStart  = NesJoypadButton(0b0000_1000)
	NesJoypadButtonSelect = NesJoypadButton(0b0000_0100)
	NesJoypadButtonB      = NesJoypadButton(0b0000_0010)
	NesJoypadButtonA      = NesJoypadButton(0b0000_0001)
)

var NesDefaultJoypadMapping = map[sdl.Scancode]NesJoypadButton{
	sdl.SCANCODE_RIGHT:  NesJoypadButtonRight,
	sdl.SCANCODE_LEFT:   NesJoypadButtonLeft,
	sdl.SCANCODE_DOWN:   NesJoypadButtonDown,
	sdl.SCANCODE_UP:     NesJoypadButtonUp,
	sdl.SCANCODE_RETURN: NesJoypadButtonStart,
	sdl.SCANCODE_SPACE:  NesJoypadButtonSelect,
	sdl.SCANCODE_S:      NesJoypadButtonB,
	sdl.SCANCODE_A:      NesJoypadButtonA,
}

type NesJoypadButton uint8
type NesJoypad struct {
	Strobe  bool
	Index   uint8
	Status  NesJoypadButton
	Mapping map[sdl.Scancode]NesJoypadButton
}

func JoypadRead(nes *Nes, i uint8) NesJoypadButton {
	if nes.Joypads[i].Index > 7 {
		return 1
	}
	data := (nes.Joypads[i].Status & (1 << nes.Joypads[i].Index)) >> nes.Joypads[i].Index
	if !nes.Joypads[i].Strobe && nes.Joypads[i].Index <= 7 {
		nes.Joypads[i].Index += 1
	}
	return data
}

func JoypadWrite(nes *Nes, i uint8, data NesJoypadButton) {
	nes.Joypads[i].Strobe = data&1 == 1
	if nes.Joypads[i].Strobe {
		nes.Joypads[i].Index = 0
	}
}

func JoypadSet(nes *Nes, i uint8, button NesJoypadButton, state bool) {
	if state {
		nes.Joypads[i].Status = NesJoypadButton(BitflagSet(uint8(nes.Joypads[i].Status), uint8(button)))
	} else {
		nes.Joypads[i].Status = NesJoypadButton(BitflagClear(uint8(nes.Joypads[i].Status), uint8(button)))
	}
}
