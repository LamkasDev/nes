package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

const NesProgramRomPageSize = uint32(16384)
const NesGraphicsRomPageSize = uint32(8192)

var NesTag = []byte{0x4E, 0x45, 0x53, 0x1A}

type NesRom struct {
	Program   NesMemory
	Graphics  NesMemory
	Mapper    NesRomMapper
	Version   uint8
	Mirroring NesMirroring
	Trainer   bool
	Testing   bool
}

func CreateRomAndLoad(nes *Nes, path string) {
	data, err := os.ReadFile(path)
	check(err)
	rom, err := CreateRomRaw(data)
	check(err)

	LoadRom(nes, rom)
}

func CreateRomRaw(data []byte) (NesRom, error) {
	// Process constant
	c := NesRom{}
	if !bytes.Equal(data[0:4], NesTag) {
		return c, errors.New("wrong format")
	}

	// Process flag 6
	flag6 := NesRomFlag6(data[6])
	fourScreen := uint8(flag6&NesRomFlag6IgnoreMirroring) != 0
	vertical := uint8(flag6&NesRomFlag6MirroringVertical) != 0
	if fourScreen {
		c.Mirroring = NesMirroringFourscreen
	} else if vertical {
		c.Mirroring = NesMirroringVertical
	} else {
		c.Mirroring = NesMirroringHorizontal
	}
	c.Trainer = uint8(flag6&NesRomFlag6Trainer) != 0

	// Process flag 7
	c.Mapper = NesRomMapper((data[7] & 0b1111_0000) | (data[6] >> 4))
	c.Version = (data[7] >> 2) & 0b11
	if c.Version != 0 {
		return c, fmt.Errorf("version %v not supported", c.Version)
	}

	// Map ROM
	err := FillRom(&c, data)
	if err != nil {
		return c, err
	}

	return c, nil
}

func LoadRom(nes *Nes, rom NesRom) {
	nes.Bus.ROM = rom
	LogLn("Loaded new ROM (len: %d, map: %d, mir: %v)...", len(nes.Bus.ROM.Program.Full), nes.Bus.ROM.Mapper, nes.Bus.ROM.Mirroring)
	SetupBus(nes)
	ResetCPU(nes)
}
