package main

// https://www.nesdev.org/wiki/Status_flags
const (
	NesCPUStatusNegative         = NesCPUStatus(0b1000_0000)
	NesCPUStatusOverflow         = NesCPUStatus(0b0100_0000)
	NesCPUStatusBreak2           = NesCPUStatus(0b0010_0000)
	NesCPUStatusBreak            = NesCPUStatus(0b0001_0000)
	NesCPUStatusDecimal          = NesCPUStatus(0b0000_1000)
	NesCPUStatusInterruptDisable = NesCPUStatus(0b0000_0100)
	NesCPUStatusZero             = NesCPUStatus(0b0000_0010)
	NesCPUStatusCarry            = NesCPUStatus(0b0000_0001)
	NesInitialStatus             = NesCPUStatusBreak2 | NesCPUStatusInterruptDisable
)

var NesCPUStatusMap = map[NesCPUStatus]string{
	NesCPUStatusCarry:            "Carry",
	NesCPUStatusZero:             "Zero",
	NesCPUStatusInterruptDisable: "Int. Dis.",
	NesCPUStatusDecimal:          "Decimal",
	NesCPUStatusBreak:            "Break",
	NesCPUStatusBreak2:           "Break 2",
	NesCPUStatusOverflow:         "Overflow",
	NesCPUStatusNegative:         "Negative",
}

type NesCPUStatus uint8

func StatusSet(nes *Nes, flag NesCPUStatus) {
	nes.CPU.Status |= flag
}
func StatusClear(nes *Nes, flag NesCPUStatus) {
	nes.CPU.Status &^= flag
}
func StatusToggle(nes *Nes, flag NesCPUStatus) {
	nes.CPU.Status ^= flag
}
func StatusHas(nes *Nes, flag NesCPUStatus) bool {
	return nes.CPU.Status&flag != 0
}
func StatusFlags(nes *Nes) []string {
	list := []string{}
	for flag, name := range NesCPUStatusMap {
		if StatusHas(nes, flag) {
			list = append(list, name)
		}
	}
	return list
}
