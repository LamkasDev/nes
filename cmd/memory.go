package main

type NesMemory struct {
	Full []byte
}

func MemoryRead(memory *NesMemory, address NesPointer) byte {
	return memory.Full[address]
}

func MemoryReadLen(memory *NesMemory, address NesPointer, len uint16) []byte {
	return memory.Full[address : address+NesPointer(len)]
}

func MemoryWrite(memory *NesMemory, address NesPointer, data []byte) {
	copy(memory.Full[address:], data)
}

func MemoryReadAddress(memory *NesMemory, address NesPointer) NesPointer {
	low := MemoryRead(memory, address)
	high := MemoryRead(memory, address+1)
	return (NesPointer(high) << 8) | NesPointer(low)
}

func MemoryWriteAddress(memory *NesMemory, address NesPointer, data NesPointer) {
	low := uint8(data & 0xff)
	high := uint8(data >> 8)
	MemoryWrite(memory, address, []byte{low})
	MemoryWrite(memory, address+1, []byte{high})
}
