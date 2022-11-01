package main

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

type NesJoypadButton uint8
type NesJoypad struct {
	Strobe bool
	Index  uint8
	Status NesJoypadButton
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
