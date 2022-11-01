package main

import "math"

func SetAccumulator(nes *Nes, n uint8) {
	nes.CPU.Accumulator = n
	UpdateZeroNegativeFlags(nes, nes.CPU.Accumulator)
}

func AddToAccumulator(nes *Nes, n uint8) {
	sum := uint16(nes.CPU.Accumulator) + uint16(n)
	if StatusHas(nes, NesCPUStatusCarry) {
		sum += 1
	}
	if sum > math.MaxUint8 {
		StatusSet(nes, NesCPUStatusCarry)
	} else {
		StatusClear(nes, NesCPUStatusCarry)
	}
	res := uint8(sum)
	if (n^res)&(res^nes.CPU.Accumulator)&0x80 != 0 {
		StatusSet(nes, NesCPUStatusOverflow)
	} else {
		StatusClear(nes, NesCPUStatusOverflow)
	}

	SetAccumulator(nes, res)
}

func SubFromAccumulator(nes *Nes, n int8) {
	AddToAccumulator(nes, uint8(WrappingSub8Int(WrappingNeg8(n), 1)))
}
