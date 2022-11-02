package main

const (
	NesMirroringVertical   = NesMirroring(0)
	NesMirroringHorizontal = NesMirroring(1)
	NesMirroringFourscreen = NesMirroring(2)
)

type NesMirroring uint8

func MirrorVRAMAddress(nes *Nes, address NesPointer) NesPointer {
	address &= NesPPUMirrorVRAM
	vramIndex := address - nes.PPU.Mapper.VRAMStart
	table := vramIndex / NesNametableSize
	switch nes.Bus.ROM.Mirroring {
	case NesMirroringVertical:
		switch table {
		case 2, 3:
			return vramIndex - NesNametableSizeDouble
		}
	case NesMirroringHorizontal:
		switch table {
		case 1, 2:
			return vramIndex - NesNametableSize
		case 3:
			return vramIndex - NesNametableSizeDouble
		}
	}
	return vramIndex
}
