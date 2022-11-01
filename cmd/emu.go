package main

func main() {
	nes := SetupNES()
	CreateRomAndLoad(&nes, "E:\\code\\go\\nes\\data\\pacman.nes")
	SetupRenderer(&nes)

	CycleGlobal(&nes)
}
