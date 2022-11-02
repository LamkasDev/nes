package main

import (
	"sync"
)

type Nes struct {
	Running bool
	Cycling bool

	CPU      NesCPU
	PPU      NesPPU
	Bus      NesBus
	Joypads  []NesJoypad
	Renderer NesRenderer
	Locks    NesLocks
	Timings  NesTimings
}
type NesLocks struct {
	Input    sync.Mutex
	Renderer sync.Mutex
}
type NesTimings struct {
	CPU      uint32
	Renderer uint16
	Input    uint16
}

func SetupNES() Nes {
	speed := 1.0
	nes := Nes{
		Running: true,
		Cycling: true,
		Timings: NesTimings{
			CPU:      uint32(((21.477272 * 1000000) / 12) * speed), // NES CPU cycles * speed
			Renderer: 1000 / NesRendererFps,                        // Delay for rendering in ms (v-synced)
			Input:    0,                                            // Delay for input in ms
		},
	}
	SetupCPUTable(&nes)
	SetupPPU(&nes)
	ResetCPU(&nes)

	return nes
}
