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

func SetupJoypads(nes *Nes) {
	nes.Joypads = append(nes.Joypads, NesJoypad{Mapping: NesDefaultJoypadMapping})
	nes.Joypads = append(nes.Joypads, NesJoypad{Mapping: NesDefaultJoypadMapping})
}

func JoypadRead(joypad *NesJoypad) NesJoypadButton {
	if joypad.Index > 7 {
		return 1
	}
	data := (joypad.Status & (1 << joypad.Index)) >> joypad.Index
	if !joypad.Strobe && joypad.Index <= 7 {
		joypad.Index += 1
	}
	return data
}

func JoypadWrite(joypad *NesJoypad, data NesJoypadButton) {
	joypad.Strobe = data&1 == 1
	if joypad.Strobe {
		joypad.Index = 0
	}
}

func JoypadSet(joypad *NesJoypad, button NesJoypadButton, state bool) {
	if state {
		joypad.Status = NesJoypadButton(BitflagSet(uint8(joypad.Status), uint8(button)))
	} else {
		joypad.Status = NesJoypadButton(BitflagClear(uint8(joypad.Status), uint8(button)))
	}
}
