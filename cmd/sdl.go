package main

import "github.com/veandco/go-sdl2/sdl"

func CycleSDL(nes *Nes) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			nes.Cycling = false
			nes.Running = false
		case *sdl.KeyboardEvent:
			switch t.Keysym.Scancode {
			case sdl.SCANCODE_ESCAPE:
				if t.Type == sdl.KEYDOWN {
					nes.Cycling = false
					nes.Running = false
				}
			case sdl.SCANCODE_W:
				if t.Type == sdl.KEYDOWN {
					BusMemoryWrite(nes, NesRAMKeys, []byte{0x77})
				}
			case sdl.SCANCODE_S:
				if t.Type == sdl.KEYDOWN {
					BusMemoryWrite(nes, NesRAMKeys, []byte{0x73})
				}
			case sdl.SCANCODE_A:
				if t.Type == sdl.KEYDOWN {
					BusMemoryWrite(nes, NesRAMKeys, []byte{0x61})
				}
			case sdl.SCANCODE_D:
				if t.Type == sdl.KEYDOWN {
					BusMemoryWrite(nes, NesRAMKeys, []byte{0x64})
				}
			}
		}
	}
}
