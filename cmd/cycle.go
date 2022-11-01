package main

import (
	"github.com/loov/hrtime"
)

func CycleGlobal(nes *Nes) {
	defer CleanRenderer(&nes.Renderer)

	lastCycle := hrtime.Now()
	lastRendererCycle := hrtime.Now()
	for nes.Cycling {
		currentCycle := hrtime.Now()

		// CPU
		diffCycle := (currentCycle - lastCycle).Microseconds()
		if diffCycle > int64(nes.CycleDelay) {
			lastCycle = currentCycle
			CycleCPU(nes)
		}
		// Renderer
		diffRendererCycle := (currentCycle - lastRendererCycle).Microseconds()
		if diffRendererCycle > int64(nes.RendererDelay) {
			lastRendererCycle = currentCycle
			CycleRenderer(nes)
		}

		CycleSDL(nes)
	}
	for nes.Running {
		CycleSDL(nes)
	}
}
