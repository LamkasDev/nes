package main

const (
	NesInterruptTypeNMI = NesInterruptType(1)
)

type NesInterruptType uint8
type NesInterrupt struct {
	Type   NesInterruptType
	Vector NesPointer
	Mask   NesCPUStatus
	Cycles uint8
}

func InterruptCreateNMI() NesInterrupt {
	return NesInterrupt{
		Type:   NesInterruptTypeNMI,
		Vector: 0xfffa,
		Mask:   NesCPUStatusBreak2,
		Cycles: 2,
	}
}
