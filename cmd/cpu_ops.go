package main

// No OP
func ProcessOpNone(nes *Nes, mode AddressingMode) {}
func ProcessOpNoneRead(nes *Nes, mode AddressingMode) {
	_, cross := GetOpAddressCross(nes, mode)
	CheckCross(nes, cross)
}

// ADC (This instruction adds the contents of a memory location to the accumulator together with the carry bit. If overflow occurs the carry bit is set, this enables multiple byte addition to be performed.)
func ProcessOpADC(nes *Nes, mode AddressingMode) {
	addr, cross := GetOpAddressCross(nes, mode)
	b := BusMemoryRead(nes, addr)
	AddToAccumulator(nes, uint8(b))
	CheckCross(nes, cross)
}

// AND (A logical AND is performed, bit by bit, on the accumulator contents using the contents of a byte of memory.)
func ProcessOpAND(nes *Nes, mode AddressingMode) {
	addr, cross := GetOpAddressCross(nes, mode)
	b := BusMemoryRead(nes, addr)
	SetAccumulator(nes, b&nes.CPU.Accumulator)
	CheckCross(nes, cross)
}

// ASL (This operation shifts all the bits of the accumulator or memory contents one bit left. Bit 0 is set to 0 and bit 7 is placed in the carry flag. The effect of this operation is to multiply the memory contents by 2 (ignoring 2's complement considerations), setting the carry if the result will not fit in 8 bits.)
func ProcessOpASL(nes *Nes, mode AddressingMode) {
	addr := GetOpAddress(nes, mode)
	b := BusMemoryRead(nes, addr)
	if b>>7 == 1 {
		StatusSet(nes, NesCPUStatusCarry)
	} else {
		StatusClear(nes, NesCPUStatusCarry)
	}
	b <<= 1
	BusMemoryWrite(nes, addr, []byte{b})
	UpdateZeroNegativeFlags(nes, b)
}

func ProcessOpASLAcc(nes *Nes, mode AddressingMode) {
	if nes.CPU.Accumulator>>7 == 1 {
		StatusSet(nes, NesCPUStatusCarry)
	} else {
		StatusClear(nes, NesCPUStatusCarry)
	}
	SetAccumulator(nes, nes.CPU.Accumulator<<1)
}

// Process branch
func ProcessBranch(nes *Nes) {
	BusTick(nes, 1)

	jump := int8(BusMemoryRead(nes, nes.CPU.Counter))
	addr := WrappingAddPtr(WrappingAddPtr(nes.CPU.Counter, NesPointer(1)), NesPointer(jump))
	if DidPageCross(WrappingAddPtr(nes.CPU.Counter, 1), addr) {
		BusTick(nes, 1)
	}
	nes.CPU.Counter = addr
}

// BCC (If the carry flag is clear then add the relative displacement to the program counter to cause a branch to a new location.)
func ProcessOpBCC(nes *Nes, mode AddressingMode) {
	if !StatusHas(nes, NesCPUStatusCarry) {
		ProcessBranch(nes)
	}
}

// BCS (If the carry flag is set then add the relative displacement to the program counter to cause a branch to a new location.)
func ProcessOpBCS(nes *Nes, mode AddressingMode) {
	if StatusHas(nes, NesCPUStatusCarry) {
		ProcessBranch(nes)
	}
}

// BEQ (If the zero flag is set then add the relative displacement to the program counter to cause a branch to a new location.)
func ProcessOpBEQ(nes *Nes, mode AddressingMode) {
	if StatusHas(nes, NesCPUStatusZero) {
		ProcessBranch(nes)
	}
}

// BIT (This instructions is used to test if one or more bits are set in a target memory location. The mask pattern in A is ANDed with the value in memory to set or clear the zero flag, but the result is not kept. Bits 7 and 6 of the value from memory are copied into the N and V flags.)
func ProcessOpBIT(nes *Nes, mode AddressingMode) {
	b := BusMemoryRead(nes, GetOpAddress(nes, mode))
	if nes.CPU.Accumulator&b == 0 {
		StatusSet(nes, NesCPUStatusZero)
	} else {
		StatusClear(nes, NesCPUStatusZero)
	}

	if BitflagHas(b, uint8(NesCPUStatusNegative)) {
		StatusSet(nes, NesCPUStatusNegative)
	} else {
		StatusClear(nes, NesCPUStatusNegative)
	}
	if BitflagHas(b, uint8(NesCPUStatusOverflow)) {
		StatusSet(nes, NesCPUStatusOverflow)
	} else {
		StatusClear(nes, NesCPUStatusOverflow)
	}
}

// BMI (If the negative flag is set then add the relative displacement to the program counter to cause a branch to a new location.)
func ProcessOpBMI(nes *Nes, mode AddressingMode) {
	if StatusHas(nes, NesCPUStatusNegative) {
		ProcessBranch(nes)
	}
}

// BNE (If the zero flag is clear then add the relative displacement to the program counter to cause a branch to a new location.)
func ProcessOpBNE(nes *Nes, mode AddressingMode) {
	if !StatusHas(nes, NesCPUStatusZero) {
		ProcessBranch(nes)
	}
}

// BPL (If the negative flag is clear then add the relative displacement to the program counter to cause a branch to a new location.)
func ProcessOpBPL(nes *Nes, mode AddressingMode) {
	if !StatusHas(nes, NesCPUStatusNegative) {
		ProcessBranch(nes)
	}
}

// BRK
func ProcessOpBRK(nes *Nes, mode AddressingMode) {
	LogWarnLn("Received break.")
	nes.Cycling = false
}

// BVC (If the overflow flag is clear then add the relative displacement to the program counter to cause a branch to a new location.)
func ProcessOpBVC(nes *Nes, mode AddressingMode) {
	if !StatusHas(nes, NesCPUStatusOverflow) {
		ProcessBranch(nes)
	}
}

// BVS (If the overflow flag is set then add the relative displacement to the program counter to cause a branch to a new location.)
func ProcessOpBVS(nes *Nes, mode AddressingMode) {
	if StatusHas(nes, NesCPUStatusOverflow) {
		ProcessBranch(nes)
	}
}

// CLC (Set the carry flag to zero.)
func ProcessOpCLC(nes *Nes, mode AddressingMode) {
	StatusClear(nes, NesCPUStatusCarry)
}

// CLD (Sets the decimal mode flag to zero.)
func ProcessOpCLD(nes *Nes, mode AddressingMode) {
	StatusClear(nes, NesCPUStatusDecimal)
}

// CLI (Clears the interrupt disable flag allowing normal interrupt requests to be serviced.)
func ProcessOpCLI(nes *Nes, mode AddressingMode) {
	StatusClear(nes, NesCPUStatusInterruptDisable)
}

// CLV (Clears the overflow flag.)
func ProcessOpCLV(nes *Nes, mode AddressingMode) {
	StatusClear(nes, NesCPUStatusOverflow)
}

// Process compare
func ProcessCompare(nes *Nes, mode AddressingMode, cmp uint8) {
	addr, cross := GetOpAddressCross(nes, mode)
	b := BusMemoryRead(nes, addr)
	if b <= cmp {
		StatusSet(nes, NesCPUStatusCarry)
	} else {
		StatusClear(nes, NesCPUStatusCarry)
	}

	UpdateZeroNegativeFlags(nes, WrappingSub8(cmp, b))
	CheckCross(nes, cross)
}

// CMP (This instruction compares the contents of the accumulator with another memory held value and sets the zero and carry flags as appropriate.)
func ProcessOpCMP(nes *Nes, mode AddressingMode) {
	ProcessCompare(nes, mode, nes.CPU.Accumulator)
}

// CPX (This instruction compares the contents of the X register with another memory held value and sets the zero and carry flags as appropriate.)
func ProcessOpCPX(nes *Nes, mode AddressingMode) {
	ProcessCompare(nes, mode, nes.CPU.IndexX)
}

// CPY (This instruction compares the contents of the Y register with another memory held value and sets the zero and carry flags as appropriate.)
func ProcessOpCPY(nes *Nes, mode AddressingMode) {
	ProcessCompare(nes, mode, nes.CPU.IndexY)
}

// DCP (Subtract 1 from memory (without borrow).)
func ProcessOpDCP(nes *Nes, mode AddressingMode) {
	addr := GetOpAddress(nes, mode)
	b := BusMemoryRead(nes, addr)
	data := uint8(WrappingSub8Int(int8(b), 1))
	BusMemoryWrite(nes, addr, []byte{data})
	if data <= nes.CPU.Accumulator {
		StatusSet(nes, NesCPUStatusCarry)
	}

	UpdateZeroNegativeFlags(nes, WrappingSub8(nes.CPU.Accumulator, data))
}

// DEC (Subtracts one from the value held at a specified memory location setting the zero and negative flags as appropriate.)
func ProcessOpDEC(nes *Nes, mode AddressingMode) {
	addr := GetOpAddress(nes, mode)
	b := BusMemoryRead(nes, addr)
	b = WrappingSub8(b, 1)
	BusMemoryWrite(nes, addr, []byte{b})
	UpdateZeroNegativeFlags(nes, b)
}

// DEX (Subtracts one from the X register setting the zero and negative flags as appropriate.)
func ProcessOpDEX(nes *Nes, mode AddressingMode) {
	nes.CPU.IndexX = WrappingSub8(nes.CPU.IndexX, 1)
	UpdateZeroNegativeFlags(nes, nes.CPU.IndexX)
}

// DEY (Subtracts one from the Y register setting the zero and negative flags as appropriate.)
func ProcessOpDEY(nes *Nes, mode AddressingMode) {
	nes.CPU.IndexY = WrappingSub8(nes.CPU.IndexY, 1)
	UpdateZeroNegativeFlags(nes, nes.CPU.IndexY)
}

// EOR (An exclusive OR is performed, bit by bit, on the accumulator contents using the contents of a byte of memory.)
func ProcessOpEOR(nes *Nes, mode AddressingMode) {
	addr, cross := GetOpAddressCross(nes, mode)
	b := BusMemoryRead(nes, addr)
	SetAccumulator(nes, b^nes.CPU.Accumulator)
	CheckCross(nes, cross)
}

// Base for increase instructions
func Increase(nes *Nes, mode AddressingMode) byte {
	addr := GetOpAddress(nes, mode)
	b := BusMemoryRead(nes, addr)
	b = WrappingAdd8(b, 1)
	BusMemoryWrite(nes, addr, []byte{b})
	UpdateZeroNegativeFlags(nes, b)

	return b
}

// ISB (Increase memory by one, then subtract memory from accu-mulator (with borrow).)
func ProcessOpISB(nes *Nes, mode AddressingMode) {
	b := Increase(nes, mode)
	SubFromAccumulator(nes, int8(b))
}

// INC (Adds one to the value held at a specified memory location setting the zero and negative flags as appropriate.)
func ProcessOpINC(nes *Nes, mode AddressingMode) {
	Increase(nes, mode)
}

// INX (Adds one to the X register setting the zero and negative flags as appropriate.)
func ProcessOpINX(nes *Nes, mode AddressingMode) {
	nes.CPU.IndexX = WrappingAdd8(nes.CPU.IndexX, 1)
	UpdateZeroNegativeFlags(nes, nes.CPU.IndexX)
}

// INY (Adds one to the Y register setting the zero and negative flags as appropriate.)
func ProcessOpINY(nes *Nes, mode AddressingMode) {
	nes.CPU.IndexY = WrappingAdd8(nes.CPU.IndexY, 1)
	UpdateZeroNegativeFlags(nes, nes.CPU.IndexY)
}

// JMP (Sets the program counter to the address specified by the operand.)
func ProcessOpJMP(nes *Nes, mode AddressingMode) {
	addr := BusMemoryReadAddress(nes, nes.CPU.Counter)
	nes.CPU.Counter = addr
}

func ProcessOpJMPIndirect(nes *Nes, mode AddressingMode) {
	addr := BusMemoryReadAddress(nes, nes.CPU.Counter)
	jmpAddr := GetJumpAddress(nes, addr)
	nes.CPU.Counter = jmpAddr
}

// Bug (https://www.nesdev.org/obelisk-6502-guide/reference.html#JMP)
func GetJumpAddress(nes *Nes, addr NesPointer) NesPointer {
	if addr&0x00FF == 0x00FF {
		low := BusMemoryRead(nes, addr)
		high := BusMemoryRead(nes, addr&0xFF00)
		return NesPointer(high)<<8 | NesPointer(low)
	}

	return BusMemoryReadAddress(nes, addr)
}

// JSR (The JSR instruction pushes the address (minus one) of the return point on to the stack and then sets the program counter to the target memory address.)
func ProcessOpJSR(nes *Nes, mode AddressingMode) {
	StackPush16(nes, uint16(nes.CPU.Counter)+2-1)
	addr := BusMemoryReadAddress(nes, nes.CPU.Counter)
	nes.CPU.Counter = addr
}

// LAX (Load accumulator and X register with memory.)
func ProcessOpLAX(nes *Nes, mode AddressingMode) {
	b := BusMemoryRead(nes, GetOpAddress(nes, mode))
	SetAccumulator(nes, b)
	nes.CPU.IndexX = b
}

// LDA (Loads a byte of memory into the accumulator setting the zero and negative flags as appropriate.)
func ProcessOpLDA(nes *Nes, mode AddressingMode) {
	addr, cross := GetOpAddressCross(nes, mode)
	b := BusMemoryRead(nes, addr)
	SetAccumulator(nes, b)
	CheckCross(nes, cross)
}

// LDX (Loads a byte of memory into the X register setting the zero and negative flags as appropriate.)
func ProcessOpLDX(nes *Nes, mode AddressingMode) {
	addr, cross := GetOpAddressCross(nes, mode)
	b := BusMemoryRead(nes, addr)
	nes.CPU.IndexX = b
	UpdateZeroNegativeFlags(nes, nes.CPU.IndexX)
	CheckCross(nes, cross)
}

// LDY (Loads a byte of memory into the Y register setting the zero and negative flags as appropriate.)
func ProcessOpLDY(nes *Nes, mode AddressingMode) {
	addr, cross := GetOpAddressCross(nes, mode)
	b := BusMemoryRead(nes, addr)
	nes.CPU.IndexY = b
	UpdateZeroNegativeFlags(nes, nes.CPU.IndexY)
	CheckCross(nes, cross)
}

// Base for LSR instructions
func LSR(nes *Nes, mode AddressingMode) byte {
	addr := GetOpAddress(nes, mode)
	b := BusMemoryRead(nes, addr)
	if b&1 == 1 {
		StatusSet(nes, NesCPUStatusCarry)
	} else {
		StatusClear(nes, NesCPUStatusCarry)
	}
	b >>= 1
	BusMemoryWrite(nes, addr, []byte{b})
	UpdateZeroNegativeFlags(nes, b)
	return b
}

// LSR (Each of the bits in A or M is shift one place to the right. The bit that was in bit 0 is shifted into the carry flag. Bit 7 is set to zero.)
func ProcessOpLSR(nes *Nes, mode AddressingMode) {
	LSR(nes, mode)
}

func ProcessOpLSRAcc(nes *Nes, mode AddressingMode) {
	if nes.CPU.Accumulator&1 == 1 {
		StatusSet(nes, NesCPUStatusCarry)
	} else {
		StatusClear(nes, NesCPUStatusCarry)
	}
	SetAccumulator(nes, nes.CPU.Accumulator>>1)
}

// ORA (An inclusive OR is performed, bit by bit, on the accumulator contents using the contents of a byte of memory.)
func ProcessOpORA(nes *Nes, mode AddressingMode) {
	addr, cross := GetOpAddressCross(nes, mode)
	b := BusMemoryRead(nes, addr)
	SetAccumulator(nes, b|nes.CPU.Accumulator)
	CheckCross(nes, cross)
}

// PHA (Pushes a copy of the accumulator on to the stack.)
func ProcessOpPHA(nes *Nes, mode AddressingMode) {
	StackPush(nes, nes.CPU.Accumulator)
}

// PHP (Pushes a copy of the status flags on to the stack.)
func ProcessOpPHP(nes *Nes, mode AddressingMode) {
	flags := nes.CPU.Status
	// https://www.nesdev.org/wiki/Status_flags#The_B_flag
	flags |= NesCPUStatusBreak
	flags |= NesCPUStatusBreak2
	StackPush(nes, uint8(flags))
}

// PLA (Pulls an 8 bit value from the stack and into the accumulator. The zero and negative flags are set as appropriate.)
func ProcessOpPLA(nes *Nes, mode AddressingMode) {
	b := StackPop(nes)
	SetAccumulator(nes, b)
}

// PLP (Pulls an 8 bit value from the stack and into the processor flags. The flags will take on new states as determined by the value pulled.)
func ProcessOpPLP(nes *Nes, mode AddressingMode) {
	nes.CPU.Status = NesCPUStatus(StackPop(nes))
	// https://www.nesdev.org/wiki/Status_flags#The_B_flag
	StatusClear(nes, NesCPUStatusBreak)
	StatusSet(nes, NesCPUStatusBreak2)
}

// Base for ROL instructions
func ROL(nes *Nes, mode AddressingMode) byte {
	addr := GetOpAddress(nes, mode)
	b := BusMemoryRead(nes, addr)
	hadCarry := StatusHas(nes, NesCPUStatusCarry)

	if b>>7 == 1 {
		StatusSet(nes, NesCPUStatusCarry)
	} else {
		StatusClear(nes, NesCPUStatusCarry)
	}
	b <<= 1
	if hadCarry {
		b |= 1
	}
	BusMemoryWrite(nes, addr, []byte{b})
	UpdateZeroNegativeFlags(nes, b)
	return b
}

// ROL (Move each of the bits in either A or M one place to the left. Bit 0 is filled with the current value of the carry flag whilst the old bit 7 becomes the new carry flag value.)
func ProcessOpROL(nes *Nes, mode AddressingMode) {
	ROL(nes, mode)
}

func ProcessOpROLAcc(nes *Nes, mode AddressingMode) {
	hadCarry := StatusHas(nes, NesCPUStatusCarry)

	if nes.CPU.Accumulator>>7 == 1 {
		StatusSet(nes, NesCPUStatusCarry)
	} else {
		StatusClear(nes, NesCPUStatusCarry)
	}
	nes.CPU.Accumulator <<= 1
	if hadCarry {
		nes.CPU.Accumulator |= 1
	}
	SetAccumulator(nes, nes.CPU.Accumulator)
}

// RLA (Rotate one bit left in memory, then AND accumulator with memory.)
func ProcessOpRLA(nes *Nes, mode AddressingMode) {
	b := ROL(nes, mode)
	SetAccumulator(nes, nes.CPU.Accumulator&b)
}

// Base for ROR instructions
func ROR(nes *Nes, mode AddressingMode) uint8 {
	addr := GetOpAddress(nes, mode)
	b := BusMemoryRead(nes, addr)
	hadCarry := StatusHas(nes, NesCPUStatusCarry)

	if b&1 == 1 {
		StatusSet(nes, NesCPUStatusCarry)
	} else {
		StatusClear(nes, NesCPUStatusCarry)
	}
	b >>= 1
	if hadCarry {
		b |= 0b10000000
	}
	BusMemoryWrite(nes, addr, []byte{b})
	UpdateZeroNegativeFlags(nes, b)
	return b
}

// ROR (Move each of the bits in either A or M one place to the right. Bit 7 is filled with the current value of the carry flag whilst the old bit 0 becomes the new carry flag value.)
func ProcessOpROR(nes *Nes, mode AddressingMode) {
	ROR(nes, mode)
}

func ProcessOpRORAcc(nes *Nes, mode AddressingMode) {
	hadCarry := StatusHas(nes, NesCPUStatusCarry)

	if nes.CPU.Accumulator&1 == 1 {
		StatusSet(nes, NesCPUStatusCarry)
	} else {
		StatusClear(nes, NesCPUStatusCarry)
	}
	nes.CPU.Accumulator >>= 1
	if hadCarry {
		nes.CPU.Accumulator |= 0b10000000
	}
	SetAccumulator(nes, nes.CPU.Accumulator)
}

// RRA (Rotate one bit right in memory, then add memory to accumulator.)
func ProcessOpRRA(nes *Nes, mode AddressingMode) {
	b := ROR(nes, mode)
	AddToAccumulator(nes, b)
}

// RTI (The RTI instruction is used at the end of an interrupt processing routine. It pulls the processor flags from the stack followed by the program counter.)
func ProcessOpRTI(nes *Nes, mode AddressingMode) {
	nes.CPU.Status = NesCPUStatus(StackPop(nes))
	// https://www.nesdev.org/wiki/Status_flags#The_B_flag
	StatusClear(nes, NesCPUStatusBreak)
	StatusSet(nes, NesCPUStatusBreak2)
	nes.CPU.Counter = NesPointer(StackPop16(nes))
}

// RTS (The RTS instruction is used at the end of a subroutine to return to the calling routine. It pulls the program counter (minus one) from the stack.)
func ProcessOpRTS(nes *Nes, mode AddressingMode) {
	nes.CPU.Counter = NesPointer(StackPop16(nes) + 1)
}

// SAX (AND X register with accumulator and store result in memory. Status)
func ProcessOpSAX(nes *Nes, mode AddressingMode) {
	b := nes.CPU.Accumulator & nes.CPU.IndexX
	BusMemoryWrite(nes, GetOpAddress(nes, mode), []byte{b})
}

// SBC (This instruction subtracts the contents of a memory location to the accumulator together with the not of the carry bit. If overflow occurs the carry bit is clear, this enables multiple byte subtraction to be performed.)
func ProcessOpSBC(nes *Nes, mode AddressingMode) {
	addr, cross := GetOpAddressCross(nes, mode)
	b := BusMemoryRead(nes, addr)
	SubFromAccumulator(nes, int8(b))
	CheckCross(nes, cross)
}

// SEC (Set the carry flag to one.)
func ProcessOpSEC(nes *Nes, mode AddressingMode) {
	StatusSet(nes, NesCPUStatusCarry)
}

// SED (Set the decimal mode flag to one.)
func ProcessOpSED(nes *Nes, mode AddressingMode) {
	StatusSet(nes, NesCPUStatusDecimal)
}

// SEI (Set the interrupt disable flag to one.)
func ProcessOpSEI(nes *Nes, mode AddressingMode) {
	StatusSet(nes, NesCPUStatusInterruptDisable)
}

// STA (Stores the contents of the accumulator into memory.)
func ProcessOpSTA(nes *Nes, mode AddressingMode) {
	addr := GetOpAddress(nes, mode)
	BusMemoryWrite(nes, addr, []byte{nes.CPU.Accumulator})
}

// STX (Stores the contents of the X register into memory.)
func ProcessOpSTX(nes *Nes, mode AddressingMode) {
	addr := GetOpAddress(nes, mode)
	BusMemoryWrite(nes, addr, []byte{nes.CPU.IndexX})
}

// STY (Stores the contents of the Y register into memory.)
func ProcessOpSTY(nes *Nes, mode AddressingMode) {
	addr := GetOpAddress(nes, mode)
	BusMemoryWrite(nes, addr, []byte{nes.CPU.IndexY})
}

// SLO (Shift left one bit in memory, then OR accumulator with memory.)
func ProcessOpSLO(nes *Nes, mode AddressingMode) {
	ProcessOpASL(nes, mode)
	b := BusMemoryRead(nes, GetOpAddress(nes, mode))
	SetAccumulator(nes, nes.CPU.Accumulator|b)
}

// SRE (Shift right one bit in memory, then EOR accumulator with memory.)
func ProcessOpSRE(nes *Nes, mode AddressingMode) {
	b := LSR(nes, mode)
	SetAccumulator(nes, nes.CPU.Accumulator^b)
}

// TAX (Copies the current contents of the accumulator into the X register and sets the zero and negative flags as appropriate.)
func ProcessOpTAX(nes *Nes, mode AddressingMode) {
	nes.CPU.IndexX = nes.CPU.Accumulator
	UpdateZeroNegativeFlags(nes, nes.CPU.IndexX)
}

// TAY (Copies the current contents of the accumulator into the Y register and sets the zero and negative flags as appropriate.)
func ProcessOpTAY(nes *Nes, mode AddressingMode) {
	nes.CPU.IndexY = nes.CPU.Accumulator
	UpdateZeroNegativeFlags(nes, nes.CPU.IndexY)
}

// TSX (Copies the current contents of the stack register into the X register and sets the zero and negative flags as appropriate.)
func ProcessOpTSX(nes *Nes, mode AddressingMode) {
	nes.CPU.IndexX = uint8(nes.CPU.Stack)
	UpdateZeroNegativeFlags(nes, nes.CPU.IndexX)
}

// TXA (Copies the current contents of the X register into the accumulator and sets the zero and negative flags as appropriate.)
func ProcessOpTXA(nes *Nes, mode AddressingMode) {
	nes.CPU.Accumulator = nes.CPU.IndexX
	UpdateZeroNegativeFlags(nes, nes.CPU.Accumulator)
}

// TXS (Copies the current contents of the X register into the stack register.)
func ProcessOpTXS(nes *Nes, mode AddressingMode) {
	nes.CPU.Stack = NesStackPointer(nes.CPU.IndexX)
}

// TYA (Copies the current contents of the Y register into the accumulator and sets the zero and negative flags as appropriate.)
func ProcessOpTYA(nes *Nes, mode AddressingMode) {
	nes.CPU.Accumulator = nes.CPU.IndexY
	UpdateZeroNegativeFlags(nes, nes.CPU.Accumulator)
}
