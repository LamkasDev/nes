package main

import "github.com/veandco/go-sdl2/sdl"

func CycleInput(nes *Nes) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			nes.Cycling = false
			nes.Running = false
		case *sdl.KeyboardEvent:
			for _, joypad := range nes.Joypads[:1] {
				button, ok := joypad.Mapping[t.Keysym.Scancode]
				if ok {
					if t.Type == sdl.KEYDOWN {
						JoypadSet(&joypad, button, true)
					} else {
						JoypadSet(&joypad, button, false)
					}
				}
			}

			switch t.Keysym.Scancode {
			case sdl.SCANCODE_ESCAPE:
				if t.Type == sdl.KEYDOWN {
					nes.Cycling = false
					nes.Running = false
				}
			}
		}
	}
}
