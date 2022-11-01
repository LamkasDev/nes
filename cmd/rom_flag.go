package main

type NesRomFlag6 uint8

const (
	NesRomFlag6MirroringHorizonal = NesRomFlag6(0b0000_0001)
	NesRomFlag6MirroringVertical  = NesRomFlag6(0b0000_0010)
	NesRomFlag6BatteryBacked      = NesRomFlag6(0b0000_0100)
	NesRomFlag6Trainer            = NesRomFlag6(0b0000_1000)
	NesRomFlag6IgnoreMirroring    = NesRomFlag6(0b0001_0000)
	NesRomFlag6Mapper             = NesRomFlag6(0b1110_0000)
)
