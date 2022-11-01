package main

const NesCPUTableSize = 0xFF

type NesCPUTableOp func(nes *Nes, mode AddressingMode)
type NesCPUTableEntry struct {
	Op      NesInstruction
	Name    string
	Length  uint8
	Cycles  uint8
	Mode    AddressingMode
	Process NesCPUTableOp
}
type NesCPUTable struct {
	Table map[NesInstruction]NesCPUTableEntry
}

func AddInstruction(nes *Nes, op NesInstruction, name string, process NesCPUTableOp, length uint8, cycles uint8, mode AddressingMode) {
	nes.CPU.Table.Table[op] = NesCPUTableEntry{
		Op:      op,
		Name:    name,
		Length:  length,
		Cycles:  cycles,
		Mode:    mode,
		Process: process,
	}
}

func SetupCPUTable(nes *Nes) {
	nes.CPU.Table.Table = make(map[NesInstruction]NesCPUTableEntry)

	// ADC
	AddInstruction(nes, 0x69, "ADC", ProcessOpADC, 2, 2, AddressingModeImmediate)
	AddInstruction(nes, 0x65, "ADC", ProcessOpADC, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0x75, "ADC", ProcessOpADC, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0x6d, "ADC", ProcessOpADC, 3, 4, AddressingModeAbsolute)

	AddInstruction(nes, 0x7d, "ADC", ProcessOpADC, 3, 4, AddressingModeAbsoluteX)
	AddInstruction(nes, 0x79, "ADC", ProcessOpADC, 3, 4, AddressingModeAbsoluteY)
	AddInstruction(nes, 0x61, "ADC", ProcessOpADC, 2, 6, AddressingModeIndirectX)
	AddInstruction(nes, 0x71, "ADC", ProcessOpADC, 2, 5, AddressingModeIndirectY)

	// AND
	AddInstruction(nes, 0x29, "AND", ProcessOpAND, 2, 2, AddressingModeImmediate)
	AddInstruction(nes, 0x25, "AND", ProcessOpAND, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0x35, "AND", ProcessOpAND, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0x2d, "AND", ProcessOpAND, 3, 4, AddressingModeAbsolute)
	AddInstruction(nes, 0x3d, "AND", ProcessOpAND, 3, 4, AddressingModeAbsoluteX)
	AddInstruction(nes, 0x39, "AND", ProcessOpAND, 3, 4, AddressingModeAbsoluteY)
	AddInstruction(nes, 0x21, "AND", ProcessOpAND, 2, 6, AddressingModeIndirectX)
	AddInstruction(nes, 0x31, "AND", ProcessOpAND, 2, 5, AddressingModeIndirectY)

	// ASL
	AddInstruction(nes, 0x0a, "ASL", ProcessOpASLAcc, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x06, "ASL", ProcessOpASL, 2, 5, AddressingModeZeroPage)
	AddInstruction(nes, 0x16, "ASL", ProcessOpASL, 2, 6, AddressingModeZeroPageX)
	AddInstruction(nes, 0x0e, "ASL", ProcessOpASL, 3, 6, AddressingModeAbsolute)
	AddInstruction(nes, 0x1e, "ASL", ProcessOpASL, 3, 7, AddressingModeAbsoluteX)

	// Branching
	AddInstruction(nes, 0x90, "BCC", ProcessOpBCC, 2, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0xb0, "BCS", ProcessOpBCS, 2, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0xf0, "BEQ", ProcessOpBEQ, 2, 2, AddressingModeNoneAddressing)

	// BIT
	AddInstruction(nes, 0x24, "BIT", ProcessOpBIT, 2, 2, AddressingModeZeroPage)
	AddInstruction(nes, 0x2c, "BIT", ProcessOpBIT, 3, 3, AddressingModeAbsolute)

	AddInstruction(nes, 0x30, "BMI", ProcessOpBMI, 2, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0xd0, "BNE", ProcessOpBNE, 2, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x10, "BPL", ProcessOpBPL, 2, 2, AddressingModeNoneAddressing)

	AddInstruction(nes, 0x00, "BRK", ProcessOpBRK, 1, 7, AddressingModeNoneAddressing)

	AddInstruction(nes, 0x50, "BVC", ProcessOpBVC, 2, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x70, "BVS", ProcessOpBVS, 2, 2, AddressingModeNoneAddressing)

	// Clear
	AddInstruction(nes, 0x18, "CLC", ProcessOpCLC, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0xd8, "CLD", ProcessOpCLD, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x58, "CLI", ProcessOpCLI, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0xb8, "CLV", ProcessOpCLV, 1, 2, AddressingModeNoneAddressing)

	// CMP
	AddInstruction(nes, 0xc9, "CMP", ProcessOpCMP, 2, 2, AddressingModeImmediate)
	AddInstruction(nes, 0xc5, "CMP", ProcessOpCMP, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0xd5, "CMP", ProcessOpCMP, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0xcd, "CMP", ProcessOpCMP, 3, 4, AddressingModeAbsolute)
	AddInstruction(nes, 0xdd, "CMP", ProcessOpCMP, 3, 4, AddressingModeAbsoluteX)
	AddInstruction(nes, 0xd9, "CMP", ProcessOpCMP, 3, 4, AddressingModeAbsoluteY)
	AddInstruction(nes, 0xc1, "CMP", ProcessOpCMP, 2, 6, AddressingModeIndirectX)
	AddInstruction(nes, 0xd1, "CMP", ProcessOpCMP, 2, 5, AddressingModeIndirectY)

	// CPX
	AddInstruction(nes, 0xe0, "CPX", ProcessOpCPX, 2, 2, AddressingModeImmediate)
	AddInstruction(nes, 0xe4, "CPX", ProcessOpCPX, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0xec, "CPX", ProcessOpCPX, 3, 4, AddressingModeAbsolute)

	// CPY
	AddInstruction(nes, 0xc0, "CPY", ProcessOpCPY, 2, 2, AddressingModeImmediate)
	AddInstruction(nes, 0xc4, "CPY", ProcessOpCPY, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0xcc, "CPY", ProcessOpCPY, 3, 4, AddressingModeAbsolute)

	// DCP
	AddInstruction(nes, 0xc7, "*DCP", ProcessOpDCP, 2, 5, AddressingModeZeroPage)
	AddInstruction(nes, 0xd7, "*DCP", ProcessOpDCP, 2, 6, AddressingModeZeroPageX)
	AddInstruction(nes, 0xcf, "*DCP", ProcessOpDCP, 3, 6, AddressingModeAbsolute)
	AddInstruction(nes, 0xdf, "*DCP", ProcessOpDCP, 3, 7, AddressingModeAbsoluteX)
	AddInstruction(nes, 0xdb, "*DCP", ProcessOpDCP, 3, 7, AddressingModeAbsoluteY)
	AddInstruction(nes, 0xc3, "*DCP", ProcessOpDCP, 2, 8, AddressingModeIndirectX)
	AddInstruction(nes, 0xd3, "*DCP", ProcessOpDCP, 2, 8, AddressingModeIndirectY)

	// DEC
	AddInstruction(nes, 0xc6, "DEC", ProcessOpDEC, 2, 5, AddressingModeZeroPage)
	AddInstruction(nes, 0xd6, "DEC", ProcessOpDEC, 2, 6, AddressingModeZeroPageX)
	AddInstruction(nes, 0xce, "DEC", ProcessOpDEC, 3, 6, AddressingModeAbsolute)
	AddInstruction(nes, 0xde, "DEC", ProcessOpDEC, 3, 7, AddressingModeAbsoluteX)

	// DEX, DEY
	AddInstruction(nes, 0xca, "DEX", ProcessOpDEX, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x88, "DEY", ProcessOpDEY, 1, 2, AddressingModeNoneAddressing)

	// EOR
	AddInstruction(nes, 0x49, "EOR", ProcessOpEOR, 2, 2, AddressingModeImmediate)
	AddInstruction(nes, 0x45, "EOR", ProcessOpEOR, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0x55, "EOR", ProcessOpEOR, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0x4d, "EOR", ProcessOpEOR, 3, 4, AddressingModeAbsolute)
	AddInstruction(nes, 0x5d, "EOR", ProcessOpEOR, 3, 4, AddressingModeAbsoluteX)
	AddInstruction(nes, 0x59, "EOR", ProcessOpEOR, 3, 4, AddressingModeAbsoluteY)
	AddInstruction(nes, 0x41, "EOR", ProcessOpEOR, 2, 6, AddressingModeIndirectX)
	AddInstruction(nes, 0x51, "EOR", ProcessOpEOR, 2, 5, AddressingModeIndirectY)

	// Unofficial - ISB
	AddInstruction(nes, 0xe7, "*ISB", ProcessOpISB, 2, 5, AddressingModeZeroPage)
	AddInstruction(nes, 0xf7, "*ISB", ProcessOpISB, 2, 6, AddressingModeZeroPageX)
	AddInstruction(nes, 0xef, "*ISB", ProcessOpISB, 3, 6, AddressingModeAbsolute)
	AddInstruction(nes, 0xff, "*ISB", ProcessOpISB, 3, 7, AddressingModeAbsoluteX)
	AddInstruction(nes, 0xfb, "*ISB", ProcessOpISB, 3, 7, AddressingModeAbsoluteY)
	AddInstruction(nes, 0xe3, "*ISB", ProcessOpISB, 2, 8, AddressingModeIndirectX)
	AddInstruction(nes, 0xf3, "*ISB", ProcessOpISB, 2, 8, AddressingModeIndirectY)

	// INC
	AddInstruction(nes, 0xe6, "INC", ProcessOpINC, 2, 5, AddressingModeZeroPage)
	AddInstruction(nes, 0xf6, "INC", ProcessOpINC, 2, 6, AddressingModeZeroPageX)
	AddInstruction(nes, 0xee, "INC", ProcessOpINC, 3, 6, AddressingModeAbsolute)
	AddInstruction(nes, 0xfe, "INC", ProcessOpINC, 3, 7, AddressingModeAbsoluteX)

	// INX, INY
	AddInstruction(nes, 0xE8, "INX", ProcessOpINX, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0xC8, "INY", ProcessOpINY, 1, 2, AddressingModeNoneAddressing)

	// JMP, JSR
	AddInstruction(nes, 0x4c, "JMP", ProcessOpJMP, 3, 3, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x6c, "JMP", ProcessOpJMPIndirect, 3, 5, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x20, "JSR", ProcessOpJSR, 3, 6, AddressingModeNoneAddressing)

	// Unofficial - LAX
	AddInstruction(nes, 0xa7, "*LAX", ProcessOpLAX, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0xb7, "*LAX", ProcessOpLAX, 2, 4, AddressingModeZeroPageY)
	AddInstruction(nes, 0xaf, "*LAX", ProcessOpLAX, 3, 4, AddressingModeAbsolute)
	AddInstruction(nes, 0xbf, "*LAX", ProcessOpLAX, 3, 4, AddressingModeAbsoluteY)
	AddInstruction(nes, 0xa3, "*LAX", ProcessOpLAX, 2, 6, AddressingModeIndirectX)
	AddInstruction(nes, 0xb3, "*LAX", ProcessOpLAX, 2, 5, AddressingModeIndirectY)

	// LDA
	AddInstruction(nes, 0xa9, "LDA", ProcessOpLDA, 2, 2, AddressingModeImmediate)
	AddInstruction(nes, 0xa5, "LDA", ProcessOpLDA, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0xb5, "LDA", ProcessOpLDA, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0xad, "LDA", ProcessOpLDA, 3, 4, AddressingModeAbsolute)
	AddInstruction(nes, 0xbd, "LDA", ProcessOpLDA, 3, 4, AddressingModeAbsoluteX)
	AddInstruction(nes, 0xb9, "LDA", ProcessOpLDA, 3, 4, AddressingModeAbsoluteY)
	AddInstruction(nes, 0xa1, "LDA", ProcessOpLDA, 2, 6, AddressingModeIndirectX)
	AddInstruction(nes, 0xb1, "LDA", ProcessOpLDA, 2, 5, AddressingModeIndirectY)

	// LDX
	AddInstruction(nes, 0xa2, "LDX", ProcessOpLDX, 2, 2, AddressingModeImmediate)
	AddInstruction(nes, 0xa6, "LDX", ProcessOpLDX, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0xb6, "LDX", ProcessOpLDX, 2, 4, AddressingModeZeroPageY)
	AddInstruction(nes, 0xae, "LDX", ProcessOpLDX, 3, 4, AddressingModeAbsolute)
	AddInstruction(nes, 0xbe, "LDX", ProcessOpLDX, 3, 4, AddressingModeAbsoluteY)

	// LDY
	AddInstruction(nes, 0xa0, "LDY", ProcessOpLDY, 2, 2, AddressingModeImmediate)
	AddInstruction(nes, 0xa4, "LDY", ProcessOpLDY, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0xb4, "LDY", ProcessOpLDY, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0xac, "LDY", ProcessOpLDY, 3, 4, AddressingModeAbsolute)
	AddInstruction(nes, 0xbc, "LDY", ProcessOpLDY, 3, 4, AddressingModeAbsoluteX)

	// LSR
	AddInstruction(nes, 0x4a, "LSR", ProcessOpLSRAcc, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x46, "LSR", ProcessOpLSR, 2, 5, AddressingModeZeroPage)
	AddInstruction(nes, 0x56, "LSR", ProcessOpLSR, 2, 6, AddressingModeZeroPageX)
	AddInstruction(nes, 0x4e, "LSR", ProcessOpLSR, 3, 6, AddressingModeAbsolute)
	AddInstruction(nes, 0x5e, "LSR", ProcessOpLSR, 3, 7, AddressingModeAbsoluteX)

	// NOP
	AddInstruction(nes, 0xea, "NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)

	// Unofficial NOP
	AddInstruction(nes, 0x80, "*NOP", ProcessOpNone, 2, 2, AddressingModeImmediate)
	AddInstruction(nes, 0x82, "*NOP", ProcessOpNone, 2, 2, AddressingModeImmediate)
	AddInstruction(nes, 0x89, "*NOP", ProcessOpNone, 2, 2, AddressingModeImmediate)
	AddInstruction(nes, 0xc2, "*NOP", ProcessOpNone, 2, 2, AddressingModeImmediate)
	AddInstruction(nes, 0xe2, "*NOP", ProcessOpNone, 2, 2, AddressingModeImmediate)

	AddInstruction(nes, 0x04, "*NOP", ProcessOpNoneRead, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0x44, "*NOP", ProcessOpNoneRead, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0x64, "*NOP", ProcessOpNoneRead, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0x14, "*NOP", ProcessOpNoneRead, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0x34, "*NOP", ProcessOpNoneRead, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0x54, "*NOP", ProcessOpNoneRead, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0x74, "*NOP", ProcessOpNoneRead, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0xd4, "*NOP", ProcessOpNoneRead, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0xf4, "*NOP", ProcessOpNoneRead, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0x0c, "*NOP", ProcessOpNoneRead, 3, 4, AddressingModeAbsolute)
	AddInstruction(nes, 0x1c, "*NOP", ProcessOpNoneRead, 3, 4, AddressingModeAbsoluteX)
	AddInstruction(nes, 0x3c, "*NOP", ProcessOpNoneRead, 3, 4, AddressingModeAbsoluteX)
	AddInstruction(nes, 0x5c, "*NOP", ProcessOpNoneRead, 3, 4, AddressingModeAbsoluteX)
	AddInstruction(nes, 0x7c, "*NOP", ProcessOpNoneRead, 3, 4, AddressingModeAbsoluteX)
	AddInstruction(nes, 0xdc, "*NOP", ProcessOpNoneRead, 3, 4, AddressingModeAbsoluteX)
	AddInstruction(nes, 0xfc, "*NOP", ProcessOpNoneRead, 3, 4, AddressingModeAbsoluteX)

	AddInstruction(nes, 0x02, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x12, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x22, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x32, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x42, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x52, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x62, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x72, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x92, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0xb2, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0xd2, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0xf2, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)

	AddInstruction(nes, 0x1a, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x3a, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x5a, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x7a, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0xda, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)
	// OpCode::new(0xea, "NOP", 1,2, AddressingMode::NoneAddressing)
	AddInstruction(nes, 0xfa, "*NOP", ProcessOpNone, 1, 2, AddressingModeNoneAddressing)

	// ORA
	AddInstruction(nes, 0x09, "ORA", ProcessOpORA, 2, 2, AddressingModeImmediate)
	AddInstruction(nes, 0x05, "ORA", ProcessOpORA, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0x15, "ORA", ProcessOpORA, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0x0d, "ORA", ProcessOpORA, 3, 4, AddressingModeAbsolute)
	AddInstruction(nes, 0x1d, "ORA", ProcessOpORA, 3, 4, AddressingModeAbsoluteX)
	AddInstruction(nes, 0x19, "ORA", ProcessOpORA, 3, 4, AddressingModeAbsoluteY)
	AddInstruction(nes, 0x01, "ORA", ProcessOpORA, 2, 6, AddressingModeIndirectX)
	AddInstruction(nes, 0x11, "ORA", ProcessOpORA, 2, 5, AddressingModeIndirectY)

	// Stack
	AddInstruction(nes, 0x48, "PHA", ProcessOpPHA, 1, 3, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x08, "PHP", ProcessOpPHP, 1, 3, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x68, "PLA", ProcessOpPLA, 1, 4, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x28, "PLP", ProcessOpPLP, 1, 4, AddressingModeNoneAddressing)

	// ROL
	AddInstruction(nes, 0x2a, "ROL", ProcessOpROLAcc, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x26, "ROL", ProcessOpROL, 2, 5, AddressingModeZeroPage)
	AddInstruction(nes, 0x36, "ROL", ProcessOpROL, 2, 6, AddressingModeZeroPageX)
	AddInstruction(nes, 0x2e, "ROL", ProcessOpROL, 3, 6, AddressingModeAbsolute)
	AddInstruction(nes, 0x3e, "ROL", ProcessOpROL, 3, 7, AddressingModeAbsoluteX)

	// ROR
	AddInstruction(nes, 0x6a, "ROR", ProcessOpRORAcc, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x66, "ROR", ProcessOpROR, 2, 5, AddressingModeZeroPage)
	AddInstruction(nes, 0x76, "ROR", ProcessOpROR, 2, 6, AddressingModeZeroPageX)
	AddInstruction(nes, 0x6e, "ROR", ProcessOpROR, 3, 6, AddressingModeAbsolute)
	AddInstruction(nes, 0x7e, "ROR", ProcessOpROR, 3, 7, AddressingModeAbsoluteX)

	// RTI, RTS
	AddInstruction(nes, 0x40, "RTI", ProcessOpRTI, 1, 6, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x60, "RTS", ProcessOpRTS, 1, 6, AddressingModeNoneAddressing)

	// Unofficial - RLA
	AddInstruction(nes, 0x27, "*RLA", ProcessOpRLA, 2, 5, AddressingModeZeroPage)
	AddInstruction(nes, 0x37, "*RLA", ProcessOpRLA, 2, 6, AddressingModeZeroPageX)
	AddInstruction(nes, 0x2f, "*RLA", ProcessOpRLA, 3, 6, AddressingModeAbsolute)
	AddInstruction(nes, 0x3f, "*RLA", ProcessOpRLA, 3, 7, AddressingModeAbsoluteX)
	AddInstruction(nes, 0x3b, "*RLA", ProcessOpRLA, 3, 7, AddressingModeAbsoluteY)
	AddInstruction(nes, 0x23, "*RLA", ProcessOpRLA, 2, 8, AddressingModeIndirectX)
	AddInstruction(nes, 0x33, "*RLA", ProcessOpRLA, 2, 8, AddressingModeIndirectY)

	// Unofficial - RRA
	AddInstruction(nes, 0x67, "*RRA", ProcessOpRRA, 2, 5, AddressingModeZeroPage)
	AddInstruction(nes, 0x77, "*RRA", ProcessOpRRA, 2, 6, AddressingModeZeroPageX)
	AddInstruction(nes, 0x6f, "*RRA", ProcessOpRRA, 3, 6, AddressingModeAbsolute)
	AddInstruction(nes, 0x7f, "*RRA", ProcessOpRRA, 3, 7, AddressingModeAbsoluteX)
	AddInstruction(nes, 0x7b, "*RRA", ProcessOpRRA, 3, 7, AddressingModeAbsoluteY)
	AddInstruction(nes, 0x63, "*RRA", ProcessOpRRA, 2, 8, AddressingModeIndirectX)
	AddInstruction(nes, 0x73, "*RRA", ProcessOpRRA, 2, 8, AddressingModeIndirectY)

	// Unofficial - SAX
	AddInstruction(nes, 0x87, "*SAX", ProcessOpSAX, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0x97, "*SAX", ProcessOpSAX, 2, 4, AddressingModeZeroPageY)
	AddInstruction(nes, 0x83, "*SAX", ProcessOpSAX, 2, 6, AddressingModeIndirectX)
	AddInstruction(nes, 0x8f, "*SAX", ProcessOpSAX, 3, 4, AddressingModeAbsolute)

	// SBC
	AddInstruction(nes, 0xe9, "SBC", ProcessOpSBC, 2, 2, AddressingModeImmediate)
	AddInstruction(nes, 0xe5, "SBC", ProcessOpSBC, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0xf5, "SBC", ProcessOpSBC, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0xed, "SBC", ProcessOpSBC, 3, 4, AddressingModeAbsolute)
	AddInstruction(nes, 0xfd, "SBC", ProcessOpSBC, 3, 4, AddressingModeAbsoluteX)
	AddInstruction(nes, 0xf9, "SBC", ProcessOpSBC, 3, 4, AddressingModeAbsoluteY)
	AddInstruction(nes, 0xe1, "SBC", ProcessOpSBC, 2, 6, AddressingModeIndirectX)
	AddInstruction(nes, 0xf1, "SBC", ProcessOpSBC, 2, 5, AddressingModeIndirectY)

	// Unofficial - SBC
	AddInstruction(nes, 0xeb, "*SBC", ProcessOpSBC, 2, 2, AddressingModeImmediate)

	// SEC, SEC, SEI
	AddInstruction(nes, 0x38, "SEC", ProcessOpSEC, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0xf8, "SED", ProcessOpSED, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x78, "SEI", ProcessOpSEI, 1, 2, AddressingModeNoneAddressing)

	// STA
	AddInstruction(nes, 0x85, "STA", ProcessOpSTA, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0x95, "STA", ProcessOpSTA, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0x8d, "STA", ProcessOpSTA, 3, 4, AddressingModeAbsolute)
	AddInstruction(nes, 0x9d, "STA", ProcessOpSTA, 3, 5, AddressingModeAbsoluteX)
	AddInstruction(nes, 0x99, "STA", ProcessOpSTA, 3, 5, AddressingModeAbsoluteY)
	AddInstruction(nes, 0x81, "STA", ProcessOpSTA, 2, 6, AddressingModeIndirectX)
	AddInstruction(nes, 0x91, "STA", ProcessOpSTA, 2, 6, AddressingModeIndirectY)

	// STX
	AddInstruction(nes, 0x86, "STX", ProcessOpSTX, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0x96, "STX", ProcessOpSTX, 2, 4, AddressingModeZeroPageY)
	AddInstruction(nes, 0x8e, "STX", ProcessOpSTX, 3, 4, AddressingModeAbsolute)

	// STY
	AddInstruction(nes, 0x84, "STY", ProcessOpSTY, 2, 3, AddressingModeZeroPage)
	AddInstruction(nes, 0x94, "STY", ProcessOpSTY, 2, 4, AddressingModeZeroPageX)
	AddInstruction(nes, 0x8c, "STY", ProcessOpSTY, 3, 4, AddressingModeAbsolute)

	// Unofficial - SLO
	AddInstruction(nes, 0x07, "*SLO", ProcessOpSLO, 2, 5, AddressingModeZeroPage)
	AddInstruction(nes, 0x17, "*SLO", ProcessOpSLO, 2, 6, AddressingModeZeroPageX)
	AddInstruction(nes, 0x0f, "*SLO", ProcessOpSLO, 3, 6, AddressingModeAbsolute)
	AddInstruction(nes, 0x1f, "*SLO", ProcessOpSLO, 3, 7, AddressingModeAbsoluteX)
	AddInstruction(nes, 0x1b, "*SLO", ProcessOpSLO, 3, 7, AddressingModeAbsoluteY)
	AddInstruction(nes, 0x03, "*SLO", ProcessOpSLO, 2, 8, AddressingModeIndirectX)
	AddInstruction(nes, 0x13, "*SLO", ProcessOpSLO, 2, 8, AddressingModeIndirectY)

	// Unofficial - SRE
	AddInstruction(nes, 0x47, "*SRE", ProcessOpSRE, 2, 5, AddressingModeZeroPage)
	AddInstruction(nes, 0x57, "*SRE", ProcessOpSRE, 2, 6, AddressingModeZeroPageX)
	AddInstruction(nes, 0x4f, "*SRE", ProcessOpSRE, 3, 6, AddressingModeAbsolute)
	AddInstruction(nes, 0x5f, "*SRE", ProcessOpSRE, 3, 7, AddressingModeAbsoluteX)
	AddInstruction(nes, 0x5b, "*SRE", ProcessOpSRE, 3, 7, AddressingModeAbsoluteY)
	AddInstruction(nes, 0x43, "*SRE", ProcessOpSRE, 2, 8, AddressingModeIndirectX)
	AddInstruction(nes, 0x53, "*SRE", ProcessOpSRE, 2, 8, AddressingModeIndirectY)

	// TAX, TAY, ...
	AddInstruction(nes, 0xaa, "TAX", ProcessOpTAX, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0xa8, "TAY", ProcessOpTAY, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0xba, "TSX", ProcessOpTSX, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x8a, "TXA", ProcessOpTXA, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x9a, "TXS", ProcessOpTXS, 1, 2, AddressingModeNoneAddressing)
	AddInstruction(nes, 0x98, "TYA", ProcessOpTYA, 1, 2, AddressingModeNoneAddressing)
}

func ProcessOp(nes *Nes, op NesInstruction) {
	c := nes.CPU.Counter
	entry, ok := nes.CPU.Table.Table[op]
	if ok {
		entry.Process(nes, entry.Mode)
		BusTick(nes, uint32(entry.Cycles))
		if nes.CPU.Counter == c {
			nes.CPU.Counter += NesPointer(entry.Length - 1)
		}
	} else {
		LogWarnLn("unknown instruction - %02x", op)
	}
}
