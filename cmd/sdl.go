package main

import "github.com/veandco/go-sdl2/sdl"

func CycleInput(nes *Nes) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			nes.Cycling = false
			nes.Running = false
		case *sdl.KeyboardEvent:
			for i := range nes.Joypads {
				button, ok := nes.Joypads[i].Mapping[t.Keysym.Scancode]
				if ok {
					if t.Type == sdl.KEYDOWN {
						JoypadSet(nes, uint8(i), button, true)
					} else {
						JoypadSet(nes, uint8(i), button, false)
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
