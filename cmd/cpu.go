package main

import "fmt"

type NesPointer uint16
type NesInstruction uint8
type NesCPU struct {
	Accumulator uint8
	IndexX      uint8
	IndexY      uint8
	Counter     NesPointer
	Status      NesCPUStatus
	Stack       NesStackPointer
	Table       NesCPUTable
}

func CycleCPU(nes *Nes) {
	// Check for interrupts
	if PPUPollNMIInterrupt(nes) {
		Interrupt(nes, InterruptCreateNMI())
	}

	// Trace
	if LogTraceEnabled {
		fmt.Println(Trace(nes))
	}

	// Process instruction
	op := NesInstruction(BusMemoryRead(nes, nes.CPU.Counter))
	nes.CPU.Counter++
	ProcessOp(nes, op)
}

func ResetCPU(nes *Nes) {
	nes.CPU.Accumulator = 0
	nes.CPU.IndexX = 0
	nes.CPU.IndexY = 0
	nes.CPU.Stack = NesStackReset
	nes.CPU.Status = NesInitialStatus

	nes.CPU.Counter = nes.Bus.Mapper.ROMStart
}

func Interrupt(nes *Nes, interrupt NesInterrupt) {
	StackPush16(nes, uint16(nes.CPU.Counter))
	flag := nes.CPU.Status
	flag |= interrupt.Mask
	StackPush(nes, uint8(flag))
	StatusSet(nes, NesCPUStatusInterruptDisable)

	BusTick(nes, uint32(interrupt.Cycles))
	nes.CPU.Counter = BusMemoryReadAddress(nes, interrupt.Vector)
}
