package main

import (
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

const NesRendererFps = 90

var RendererFrame NesFrame

type Renderer struct {
	Window   *sdl.Window
	Surface  *sdl.Surface
	Renderer *sdl.Renderer
	Texture  *sdl.Texture
	Pitch    int
}

func SetupRenderer(nes *Nes) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	check(err)

	w, h := int32(NesFrameWidth*3), int32(NesFrameHeight*3)
	nes.Renderer.Window, err = sdl.CreateWindow("NES", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, w, h, sdl.WINDOW_SHOWN)
	check(err)

	nes.Renderer.Surface, err = nes.Renderer.Window.GetSurface()
	check(err)

	nes.Renderer.Renderer, err = sdl.CreateRenderer(nes.Renderer.Window, -1, sdl.RENDERER_ACCELERATED)
	check(err)

	nes.Renderer.Texture, err = nes.Renderer.Renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STREAMING, NesFrameWidth, NesFrameHeight)
	check(err)

	nes.Renderer.Pitch = int(NesFrameWidth * 4)

	// ShowTileBank(nes, 0)
}

func CycleRenderer(nes *Nes) {
	nes.Renderer.Texture.Update(nil, unsafe.Pointer(&RendererFrame.Data), nes.Renderer.Pitch)
	nes.Renderer.Renderer.Clear()
	nes.Renderer.Renderer.Copy(nes.Renderer.Texture, nil, nil)
	nes.Renderer.Renderer.Present()
}

func CleanRenderer(renderer *Renderer) {
	renderer.Texture.Destroy()
	renderer.Renderer.Destroy()
	renderer.Window.Destroy()
	sdl.Quit()
}
