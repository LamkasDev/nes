package main

import (
	"testing"
)

func TestJoypadStrobe(t *testing.T) {
	nes := SetupNES()
	JoypadWrite(&nes.Joypads[0], 1)
	JoypadSet(&nes.Joypads[0], NesJoypadButtonA, true)
	for x := 0; x < 10; x++ {
		b := JoypadRead(&nes.Joypads[0])
		if b != 1 {
			t.Fatalf("Failed joypad read of %v (v: %v)", 1, b)
		}
	}
}

func TestJoypadStrobeOnOff(t *testing.T) {
	nes := SetupNES()
	JoypadWrite(&nes.Joypads[0], 0)
	JoypadSet(&nes.Joypads[0], NesJoypadButtonRight, true)
	JoypadSet(&nes.Joypads[0], NesJoypadButtonLeft, true)
	JoypadSet(&nes.Joypads[0], NesJoypadButtonSelect, true)
	JoypadSet(&nes.Joypads[0], NesJoypadButtonB, true)
	for i := 0; i <= 1; i++ {
		b0 := JoypadRead(&nes.Joypads[0])
		b1 := JoypadRead(&nes.Joypads[0])
		b2 := JoypadRead(&nes.Joypads[0])
		b3 := JoypadRead(&nes.Joypads[0])
		b4 := JoypadRead(&nes.Joypads[0])
		b5 := JoypadRead(&nes.Joypads[0])
		b6 := JoypadRead(&nes.Joypads[0])
		b7 := JoypadRead(&nes.Joypads[0])
		if b0 != 0 || b1 != 1 || b2 != 1 || b3 != 0 || b4 != 0 || b5 != 0 || b6 != 1 || b7 != 1 {
			t.Fatalf("Failed joypad read.")
		}

		for x := 0; x < 10; x++ {
			b := JoypadRead(&nes.Joypads[0])
			if b != 1 {
				t.Fatalf("Failed joypad read of %v (v: %v)", 1, b)
			}
		}
		JoypadWrite(&nes.Joypads[0], 1)
		JoypadWrite(&nes.Joypads[0], 0)
	}
}
