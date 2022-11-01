package main

const (
	AddressingModeImmediate      = AddressingMode(0)
	AddressingModeZeroPage       = AddressingMode(1)
	AddressingModeZeroPageX      = AddressingMode(2)
	AddressingModeZeroPageY      = AddressingMode(3)
	AddressingModeAbsolute       = AddressingMode(4)
	AddressingModeAbsoluteX      = AddressingMode(5)
	AddressingModeAbsoluteY      = AddressingMode(6)
	AddressingModeIndirectX      = AddressingMode(7)
	AddressingModeIndirectY      = AddressingMode(8)
	AddressingModeNoneAddressing = AddressingMode(9)
)

type AddressingMode uint8

func DidPageCross(a NesPointer, b NesPointer) bool {
	return a&0xff00 != b&0xff00
}

func GetAbsoluteAddress(nes *Nes, mode AddressingMode, address NesPointer) (NesPointer, bool) {
	switch mode {
	case AddressingModeZeroPage:
		return NesPointer(BusMemoryRead(nes, address)), false
	case AddressingModeZeroPageX:
		return WrappingAdd8Ptr(BusMemoryRead(nes, address), nes.CPU.IndexX), false
	case AddressingModeZeroPageY:
		return WrappingAdd8Ptr(BusMemoryRead(nes, address), nes.CPU.IndexY), false
	case AddressingModeAbsolute:
		return BusMemoryReadAddress(nes, address), false
	case AddressingModeAbsoluteX:
		base := BusMemoryReadAddress(nes, address)
		addr := WrappingAddPtr(base, NesPointer(nes.CPU.IndexX))
		return addr, DidPageCross(base, addr)
	case AddressingModeAbsoluteY:
		base := BusMemoryReadAddress(nes, address)
		addr := WrappingAddPtr(base, NesPointer(nes.CPU.IndexY))
		return addr, DidPageCross(base, addr)
	case AddressingModeIndirectX:
		base := BusMemoryRead(nes, address)
		ptr := WrappingAdd8(base, nes.CPU.IndexX)
		low := BusMemoryRead(nes, NesPointer(ptr))
		high := BusMemoryRead(nes, NesPointer(WrappingAdd8(ptr, 1)))
		der := NesPointer(high)<<8 | NesPointer(low)

		return der, false
	case AddressingModeIndirectY:
		base := BusMemoryRead(nes, address)
		low := BusMemoryRead(nes, NesPointer(base))
		high := BusMemoryRead(nes, WrappingAdd8Ptr(base, 1))
		derBase := NesPointer(high)<<8 | NesPointer(low)
		der := WrappingAddPtr(derBase, NesPointer(nes.CPU.IndexY))

		return der, DidPageCross(der, derBase)
	default:
		panic("Adressing mode is not supported.")
	}
}

func GetOpAddress(nes *Nes, mode AddressingMode) NesPointer {
	switch mode {
	case AddressingModeImmediate:
		return nes.CPU.Counter
	default:
		addr, _ := GetAbsoluteAddress(nes, mode, nes.CPU.Counter)
		return addr
	}
}

func GetOpAddressCross(nes *Nes, mode AddressingMode) (NesPointer, bool) {
	switch mode {
	case AddressingModeImmediate:
		return nes.CPU.Counter, false
	default:
		addr, cross := GetAbsoluteAddress(nes, mode, nes.CPU.Counter)
		return addr, cross
	}
}

func CheckCross(nes *Nes, cross bool) {
	if cross {
		BusTick(nes, 1)
	}
}
