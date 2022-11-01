package main

import (
	"fmt"
)

const NesHeaderSize = 16
const NesTrainerSize = 512

type NesRomMapper uint8

func FillRom(rom *NesRom, data []byte) error {
	switch rom.Mapper {
	case 0:
		return FillRomDefault(rom, data)
	case 3:
		LogWarnLn("Using experimental mapper (v: %v)", rom.Mapper)
		return FillRomDefault(rom, data)
	}
	return fmt.Errorf("mapper %v not supported", rom.Mapper)
}

func FillRomDefault(rom *NesRom, data []byte) error {
	programSize := uint32(data[4]) * NesProgramRomPageSize
	programStart := uint32(NesHeaderSize)
	if rom.Trainer {
		programStart += NesTrainerSize
	}
	rom.Program.Full = make([]byte, programSize)
	MemoryWrite(&rom.Program, 0, data[programStart:programStart+programSize])

	graphicsSize := uint32(data[5]) * NesGraphicsRomPageSize
	graphicsStart := programStart + programSize
	rom.Graphics.Full = make([]byte, graphicsSize)
	MemoryWrite(&rom.Graphics, 0, data[graphicsStart:graphicsStart+graphicsSize])

	return nil
}
