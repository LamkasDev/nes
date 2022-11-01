package main

type NesStackPointer uint8

const NesStackReset = NesStackPointer(0xfd)

func StackPop(nes *Nes) uint8 {
	nes.CPU.Stack = WrappingAdd8StackPtr(uint8(nes.CPU.Stack), 1)
	return BusMemoryRead(nes, NesRAMStackStart+NesPointer(nes.CPU.Stack))
}

func StackPop16(nes *Nes) uint16 {
	low := uint16(StackPop(nes))
	high := uint16(StackPop(nes))
	return high<<8 | low
}

func StackPush(nes *Nes, data uint8) {
	BusMemoryWrite(nes, NesRAMStackStart+NesPointer(nes.CPU.Stack), []byte{data})
	nes.CPU.Stack = NesStackPointer(WrappingSub8(uint8(nes.CPU.Stack), 1))
}

func StackPush16(nes *Nes, data uint16) {
	low := uint8(data & 0xff)
	high := uint8(data >> 8)
	StackPush(nes, high)
	StackPush(nes, low)
}

func StackSize(nes *Nes) uint16 {
	return uint16(NesRAMStackEnd - (NesRAMStackStart + NesPointer(nes.CPU.Stack)))
}
