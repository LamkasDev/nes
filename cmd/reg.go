package main

func UpdateNegativeFlags(nes *Nes, n uint8) {
	if n>>7 == 1 {
		StatusSet(nes, NesCPUStatusNegative)
	} else {
		StatusClear(nes, NesCPUStatusNegative)
	}
}

func UpdateZeroNegativeFlags(nes *Nes, n uint8) {
	if n == 0 {
		StatusSet(nes, NesCPUStatusZero)
	} else {
		StatusClear(nes, NesCPUStatusZero)
	}

	UpdateNegativeFlags(nes, n)
}
