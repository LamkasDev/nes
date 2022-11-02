package main

import (
	"runtime/pprof"
	"syscall"
	"time"
)

var (
	winmmDLL            = syscall.NewLazyDLL("winmm.dll")
	procTimeBeginPeriod = winmmDLL.NewProc("timeBeginPeriod")
)

func CycleGlobal(nes *Nes) {
	procTimeBeginPeriod.Call(uintptr(1))
	defer pprof.StopCPUProfile()
	defer CleanRenderer(&nes.Renderer)

	go func() {
		cycles := nes.Timings.CPU / 1000
		for nes.Cycling {
			for i := uint32(0); i < cycles; i++ {
				CycleCPU(nes)
			}
			time.Sleep(time.Millisecond)
		}
	}()
	rendererElapsed := uint16(0)
	inputElapsed := uint16(0)
	for nes.Cycling {
		if rendererElapsed >= nes.Timings.Renderer {
			nes.Locks.Renderer.Lock()
			CycleRenderer(nes)
			nes.Locks.Renderer.Unlock()
			rendererElapsed = 0
		}
		if inputElapsed >= nes.Timings.Input {
			nes.Locks.Input.Lock()
			CycleInput(nes)
			nes.Locks.Input.Unlock()
			inputElapsed = 0
		}

		time.Sleep(time.Millisecond)
		rendererElapsed++
		inputElapsed++
	}
}
