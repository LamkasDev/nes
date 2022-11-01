package main

type Nes struct {
	Running       bool
	Cycling       bool
	CycleDelay    float32
	RendererDelay float32

	CPU      NesCPU
	PPU      NesPPU
	Bus      NesBus
	Renderer Renderer
}

func SetupNES() Nes {
	nes := Nes{
		Running:       true,
		Cycling:       true,
		CycleDelay:    100, // 0.1 ms
		RendererDelay: (1000 / NesRendererFps) * 1000,
	}
	SetupCPUTable(&nes)
	SetupPPU(&nes)
	ResetCPU(&nes)

	return nes
}
