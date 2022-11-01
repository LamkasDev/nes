package main

import (
	"bytes"
	"testing"
)

type InstructionTestCheck func(nes *Nes) bool
type InstructionTest struct {
	ROM            []byte
	Memory         []byte
	Stack          []byte
	Flags          []NesCPUStatus
	InitializeRegs func(nes *Nes)

	CheckResults map[string]InstructionTestCheck
	CheckFlags   map[NesCPUStatus]bool
	CheckMemory  []byte
	CheckStack   []byte
}

func RunInstructionTest(suite *testing.T, test InstructionTest) Nes {
	// Disable logging
	LogInfoEnabled = false

	// Setup NES
	nes := SetupNES()
	LoadRom(&nes, NesRom{Program: NesMemory{Full: test.ROM}, Testing: true})
	if test.Memory != nil {
		copy(nes.Bus.RAM.Full[:], test.Memory)
	}
	if test.Stack != nil {
		off := NesPointer(len(test.Stack)) - 1
		copy(nes.Bus.RAM.Full[NesRAMStackEnd-off:], test.Stack)
		nes.CPU.Stack = NesStackReset - NesStackPointer(len(test.Stack))
	}
	for _, flag := range test.Flags {
		StatusSet(&nes, flag)
	}
	if test.InitializeRegs != nil {
		test.InitializeRegs(&nes)
	}

	// Process ROM
	for nes.Cycling {
		CycleCPU(&nes)
	}

	// Check results
	for name, check := range test.CheckResults {
		if !check(&nes) {
			suite.Fatalf(name)
		}
	}
	for flag, value := range test.CheckFlags {
		if value == true && !StatusHas(&nes, flag) {
			suite.Fatalf("Expected %v flag", NesCPUStatusMap[flag])
		} else if value == false && StatusHas(&nes, flag) {
			suite.Fatalf("Unexpected %v flag", NesCPUStatusMap[flag])
		}
	}
	if test.CheckMemory != nil {
		if !bytes.Equal(MemoryReadLen(&nes.Bus.RAM, 0, uint16(len(test.CheckMemory))), test.CheckMemory) {
			suite.Fatalf("Expected %v in memory", test.CheckMemory)
		}
	}
	if test.CheckStack != nil {
		addr := NesRAMStackStart + NesPointer(nes.CPU.Stack) + 1
		if !bytes.Equal(MemoryReadLen(&nes.Bus.RAM, addr, StackSize(&nes)+1), test.CheckStack) {
			suite.Fatalf("Expected %v in stack", test.CheckStack)
		}
	}

	return nes
}

func AddressingModeCheck(t *testing.T, nes *Nes, mode AddressingMode, v NesPointer) {
	addr := GetOpAddress(nes, mode)
	if addr != v {
		t.Fatalf("%v failed. (v: 0x%04x, expected: 0x%04x)", mode, addr, v)
	}
}

// Adressing modes
func TestAddressingModes(t *testing.T) {
	nes := SetupNES()
	LoadRom(&nes, NesRom{Program: NesMemory{Full: []byte{0xfe, 0xef}}, Testing: true})
	nes.CPU.IndexX = 0x0f
	nes.CPU.IndexY = 0x0f

	AddressingModeCheck(t, &nes, AddressingModeImmediate, nes.Bus.Mapper.Sections[NesMapperSectionTypeROM].Start)
	AddressingModeCheck(t, &nes, AddressingModeZeroPage, 0x00fe)
	AddressingModeCheck(t, &nes, AddressingModeZeroPageX, WrappingAdd8Ptr(0x00fe, nes.CPU.IndexX))
	AddressingModeCheck(t, &nes, AddressingModeZeroPageY, WrappingAdd8Ptr(0x00fe, nes.CPU.IndexY))
	AddressingModeCheck(t, &nes, AddressingModeAbsolute, 0xeffe)
	AddressingModeCheck(t, &nes, AddressingModeAbsoluteX, WrappingAddPtr(0xeffe, NesPointer(nes.CPU.IndexX)))
	AddressingModeCheck(t, &nes, AddressingModeAbsoluteY, WrappingAddPtr(0xeffe, NesPointer(nes.CPU.IndexY)))

	indX := WrappingAdd8Ptr(0x00fe, nes.CPU.IndexX)
	copy(MemoryReadLen(&nes.Bus.RAM, indX, 1), []byte{0xef, 0xfe})
	AddressingModeCheck(t, &nes, AddressingModeIndirectX, 0xfeef)

	indY := NesPointer(0x00fe)
	copy(MemoryReadLen(&nes.Bus.RAM, indY, 1), []byte{0xef, 0xfe})
	AddressingModeCheck(t, &nes, AddressingModeIndirectY, WrappingAddPtr(0xfeef, NesPointer(nes.CPU.IndexY)))
}

// Trace test
func TestTrace(t *testing.T) {
	nes := SetupNES()
	BusMemoryWrite(&nes, 100, []byte{0xa2, 0x01, 0xca, 0x88, 0x00})

	nes.CPU.Counter = 0x64
	nes.CPU.Accumulator = 1
	nes.CPU.IndexX = 2
	nes.CPU.IndexY = 3
	r1 := Trace(&nes)
	CycleCPU(&nes)
	r2 := Trace(&nes)
	CycleCPU(&nes)
	r3 := Trace(&nes)
	CycleCPU(&nes)

	if r1 != "0064  A2 01     LDX #$01                        A:01 X:02 Y:03 P:24 SP:FD" {
		t.Fatalf(r1)
	}
	if r2 != "0066  CA        DEX                             A:01 X:01 Y:03 P:24 SP:FD" {
		t.Fatalf(r2)
	}
	if r3 != "0067  88        DEY                             A:01 X:00 Y:03 P:26 SP:FD" {
		t.Fatalf(r3)
	}
}

// ADC instructions
func TestOpADC(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x65, 0x00, 0x00}, Memory: []byte{1},
		InitializeRegs: func(nes *Nes) { nes.CPU.Accumulator = 1 },
		CheckResults: map[string]InstructionTestCheck{
			"Accumulator failed add": func(nes *Nes) bool { return nes.CPU.Accumulator == 2 },
		},
		CheckFlags: map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: false, NesCPUStatusCarry: false},
	}
	RunInstructionTest(t, test)

	// Test zero and carry flag
	test.Memory = []byte{255}
	test.CheckResults = map[string]InstructionTestCheck{
		"Accumulator failed add (carry)": func(nes *Nes) bool { return nes.CPU.Accumulator == 0 },
	}
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: true, NesCPUStatusNegative: false, NesCPUStatusCarry: true}
	RunInstructionTest(t, test)
}

// AND instructions (tests true and false cases)
func TestOpAND(t *testing.T) {
	// Test true case
	test := InstructionTest{
		ROM: []byte{0x25, 0x00, 0x00}, Memory: []byte{1},
		InitializeRegs: func(nes *Nes) { nes.CPU.Accumulator = 1 },
		CheckResults: map[string]InstructionTestCheck{
			"Accumulator is not 1": func(nes *Nes) bool { return nes.CPU.Accumulator == 1 },
		},
		CheckFlags: map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: false},
	}
	RunInstructionTest(t, test)

	// Test false case
	test.Memory = []byte{2}
	test.CheckResults = map[string]InstructionTestCheck{
		"Accumulator is not 0": func(nes *Nes) bool { return nes.CPU.Accumulator == 0 },
	}
	test.CheckFlags[NesCPUStatusZero] = true
	RunInstructionTest(t, test)
}

// ASL instructions
func TestOpASL(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x06, 0x00, 0x00}, Memory: []byte{0b0000_0001},
		CheckFlags:  map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: false, NesCPUStatusCarry: false},
		CheckMemory: []byte{0b0000_0010},
	}
	RunInstructionTest(t, test)

	// Test negative flag
	test.Memory = []byte{0b0100_0000}
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: true, NesCPUStatusCarry: false}
	test.CheckMemory = []byte{0b1000_0000}
	RunInstructionTest(t, test)

	// Test carry and zero flag
	test.Memory = []byte{0b1000_0000}
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: true, NesCPUStatusNegative: false, NesCPUStatusCarry: true}
	test.CheckMemory = []byte{0b0000_0000}
	RunInstructionTest(t, test)
}

// Base for branching tests
func Branch(t *testing.T, code uint8, jumpFlags []NesCPUStatus, continueFlags []NesCPUStatus) {
	test := InstructionTest{
		ROM:   []byte{code, 2, 0x00 /* on set */, 0x00 /* on clear */, 0x00},
		Flags: jumpFlags,
		CheckResults: map[string]InstructionTestCheck{
			"Expected branch jump": func(nes *Nes) bool {
				return nes.CPU.Counter == nes.Bus.Mapper.Sections[NesMapperSectionTypeROM].Start+5
			},
		},
	}
	RunInstructionTest(t, test)

	test.Flags = continueFlags
	test.CheckResults = map[string]InstructionTestCheck{
		"Unexpected branch jump": func(nes *Nes) bool {
			return nes.CPU.Counter == nes.Bus.Mapper.Sections[NesMapperSectionTypeROM].Start+3
		},
	}
	RunInstructionTest(t, test)
}

// BCC instructions
func TestOpBCC(t *testing.T) {
	Branch(t, 0x90, []NesCPUStatus{}, []NesCPUStatus{NesCPUStatusCarry})
}

// BCS instructions
func TestOpBCS(t *testing.T) {
	Branch(t, 0xb0, []NesCPUStatus{NesCPUStatusCarry}, []NesCPUStatus{})
}

// BEQ instructions
func TestOpBEQ(t *testing.T) {
	Branch(t, 0xf0, []NesCPUStatus{NesCPUStatusZero}, []NesCPUStatus{})
}

// BIT instructions
func TestOpBIT(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x24, 0x00, 0x00}, Memory: []byte{0b1100_0000},
		InitializeRegs: func(nes *Nes) { nes.CPU.Accumulator = 0b1100_0000 },
		CheckFlags:     map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusOverflow: true, NesCPUStatusNegative: true},
	}
	RunInstructionTest(t, test)

	test.Memory = []byte{0b0011_0000}
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: true, NesCPUStatusOverflow: false, NesCPUStatusNegative: false}
	RunInstructionTest(t, test)
}

// BMI instructions
func TestOpBMI(t *testing.T) {
	Branch(t, 0x30, []NesCPUStatus{NesCPUStatusNegative}, []NesCPUStatus{})
}

// BNE instructions
func TestOpBNE(t *testing.T) {
	Branch(t, 0xd0, []NesCPUStatus{}, []NesCPUStatus{NesCPUStatusZero})
}

// BPL instructions
func TestOpBPL(t *testing.T) {
	Branch(t, 0x10, []NesCPUStatus{}, []NesCPUStatus{NesCPUStatusNegative})
}

// BVC instructions
func TestOpBVC(t *testing.T) {
	Branch(t, 0x50, []NesCPUStatus{}, []NesCPUStatus{NesCPUStatusOverflow})
}

// BVS instructions
func TestOpBVS(t *testing.T) {
	Branch(t, 0x70, []NesCPUStatus{NesCPUStatusOverflow}, []NesCPUStatus{})
}

// CLC instructions
func TestOpCLC(t *testing.T) {
	test := InstructionTest{
		ROM:        []byte{0x18, 0x00},
		Flags:      []NesCPUStatus{NesCPUStatusCarry},
		CheckFlags: map[NesCPUStatus]bool{NesCPUStatusCarry: false},
	}
	RunInstructionTest(t, test)
}

// CLD instructions
func TestOpCLD(t *testing.T) {
	test := InstructionTest{
		ROM:        []byte{0xd8, 0x00},
		Flags:      []NesCPUStatus{NesCPUStatusDecimal},
		CheckFlags: map[NesCPUStatus]bool{NesCPUStatusDecimal: false},
	}
	RunInstructionTest(t, test)
}

// CLI instructions
func TestOpCLI(t *testing.T) {
	test := InstructionTest{
		ROM:        []byte{0x58, 0x00},
		Flags:      []NesCPUStatus{NesCPUStatusInterruptDisable},
		CheckFlags: map[NesCPUStatus]bool{NesCPUStatusInterruptDisable: false},
	}
	RunInstructionTest(t, test)
}

// CLV instructions
func TestOpCLV(t *testing.T) {
	test := InstructionTest{
		ROM:        []byte{0xb8, 0x00},
		Flags:      []NesCPUStatus{NesCPUStatusOverflow},
		CheckFlags: map[NesCPUStatus]bool{NesCPUStatusOverflow: false},
	}
	RunInstructionTest(t, test)
}

// Base for compare tests
func Compare(t *testing.T, code byte, InitializeRegs func(nes *Nes)) {
	// Test true case
	test := InstructionTest{
		ROM: []byte{code, 0x00, 0x00}, Memory: []byte{1},
		InitializeRegs: InitializeRegs,
		CheckFlags:     map[NesCPUStatus]bool{NesCPUStatusCarry: true, NesCPUStatusZero: true},
	}
	RunInstructionTest(t, test)

	// Test false case (and negative flag)
	test.Memory = []byte{2}
	test.CheckFlags[NesCPUStatusCarry] = false
	test.CheckFlags[NesCPUStatusZero] = false
	test.CheckFlags[NesCPUStatusNegative] = true
	RunInstructionTest(t, test)
}

// CMP instructions
func TestOpCMP(t *testing.T) {
	Compare(t, 0xc5, func(nes *Nes) { nes.CPU.Accumulator = 1 })
}

// CPX instructions
func TestOpCPX(t *testing.T) {
	Compare(t, 0xe4, func(nes *Nes) { nes.CPU.IndexX = 1 })
}

// CPY instructions
func TestOpCPY(t *testing.T) {
	Compare(t, 0xc4, func(nes *Nes) { nes.CPU.IndexY = 1 })
}

// DEC instructions
func TestOpDEC(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0xc6, 0x00, 0x00}, Memory: []byte{1},
		CheckFlags:  map[NesCPUStatus]bool{NesCPUStatusZero: true, NesCPUStatusNegative: false},
		CheckMemory: []byte{0},
	}
	RunInstructionTest(t, test)

	test.Memory = []byte{0}
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: true}
	test.CheckMemory = []byte{255}
	RunInstructionTest(t, test)
}

// Base for decrement tests
func Decrement(t *testing.T, code byte, InitializeRegs func(nes *Nes), CheckResult func(nes *Nes) bool, InitializeRegsZero func(nes *Nes), CheckResultZero func(nes *Nes) bool) {
	test := InstructionTest{
		ROM:            []byte{code, 0x00},
		InitializeRegs: InitializeRegs,
		CheckFlags:     map[NesCPUStatus]bool{NesCPUStatusZero: true, NesCPUStatusNegative: false},
		CheckResults: map[string]InstructionTestCheck{
			"Register failed decrement to zero": CheckResult,
		},
	}
	RunInstructionTest(t, test)

	test.InitializeRegs = InitializeRegsZero
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: true}
	test.CheckResults = map[string]InstructionTestCheck{
		"Register failed decrement to negative": CheckResultZero,
	}
	RunInstructionTest(t, test)
}

// DEX instructions
func TestOpDEX(t *testing.T) {
	Decrement(t, 0xca,
		func(nes *Nes) { nes.CPU.IndexX = 1 }, func(nes *Nes) bool { return nes.CPU.IndexX == 0 },
		func(nes *Nes) { nes.CPU.IndexX = 0 }, func(nes *Nes) bool { return nes.CPU.IndexX == 255 },
	)
}

// DEY instructions
func TestOpDEY(t *testing.T) {
	Decrement(t, 0x88,
		func(nes *Nes) { nes.CPU.IndexY = 1 }, func(nes *Nes) bool { return nes.CPU.IndexY == 0 },
		func(nes *Nes) { nes.CPU.IndexY = 0 }, func(nes *Nes) bool { return nes.CPU.IndexY == 255 },
	)
}

// EOR instructions
func TestOpEOR(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x45, 0x00, 0x00}, Memory: []byte{0b0101_0101},
		InitializeRegs: func(nes *Nes) { nes.CPU.Accumulator = 0b0000_1111 },
		CheckFlags:     map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: false},
		CheckResults: map[string]InstructionTestCheck{
			"Accumulator failed XOR": func(nes *Nes) bool { return nes.CPU.Accumulator == 0b0101_1010 },
		},
	}
	RunInstructionTest(t, test)

	// Test negative flags
	test.Memory = []byte{0b0000_0000}
	test.InitializeRegs = func(nes *Nes) { nes.CPU.Accumulator = 0b1000_0000 }
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: true}
	test.CheckResults = map[string]InstructionTestCheck{
		"Accumulator failed XOR": func(nes *Nes) bool { return nes.CPU.Accumulator == 0b1000_0000 },
	}
	RunInstructionTest(t, test)

	// Test zero flag
	test.Memory = []byte{0b1111_1111}
	test.InitializeRegs = func(nes *Nes) { nes.CPU.Accumulator = 0b1111_1111 }
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: true, NesCPUStatusNegative: false}
	test.CheckResults = map[string]InstructionTestCheck{
		"Accumulator failed XOR": func(nes *Nes) bool { return nes.CPU.Accumulator == 0b0000_0000 },
	}
	RunInstructionTest(t, test)
}

// INC instructions
func TestOpINC(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0xe6, 0x00, 0x00}, Memory: []byte{1},
		CheckFlags:  map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: false},
		CheckMemory: []byte{2},
	}
	RunInstructionTest(t, test)
}

// Base for increment tests
func Increment(t *testing.T, code byte, InitializeRegs func(nes *Nes), CheckResult func(nes *Nes) bool) {
	test := InstructionTest{
		ROM:            []byte{code, 0x00},
		InitializeRegs: InitializeRegs,
		CheckFlags:     map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: false},
		CheckResults: map[string]InstructionTestCheck{
			"Register failed increment": CheckResult,
		},
	}
	RunInstructionTest(t, test)
}

// INX instructions
func TestOpINX(t *testing.T) {
	Increment(t, 0xe8,
		func(nes *Nes) { nes.CPU.IndexX = 1 }, func(nes *Nes) bool { return nes.CPU.IndexX == 2 },
	)
}

// INY instructions
func TestOpINY(t *testing.T) {
	Increment(t, 0xc8,
		func(nes *Nes) { nes.CPU.IndexY = 1 }, func(nes *Nes) bool { return nes.CPU.IndexY == 2 },
	)
}

// Base for load tests
func Load(t *testing.T, code byte, GetValue func(nes *Nes) uint8) {
	test := InstructionTest{
		ROM: []byte{code, 0x00, 0x00}, Memory: []byte{1},
		CheckFlags: map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: false},
		CheckResults: map[string]InstructionTestCheck{
			"Register failed load": func(nes *Nes) bool { return GetValue(nes) == MemoryRead(&nes.Bus.RAM, 0) },
		},
	}
	RunInstructionTest(t, test)

	test.Memory = []byte{0}
	test.CheckFlags[NesCPUStatusZero] = true
	RunInstructionTest(t, test)
}

// JMP instructions
func TestOpJMP(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x4c, 0x0d, 0x0d, 0x00},
		CheckResults: map[string]InstructionTestCheck{
			"Expected jump": func(nes *Nes) bool { return nes.CPU.Counter == (0x0d0d)+1 },
		},
	}
	RunInstructionTest(t, test)
}

// JSR instructions
func TestOpJSR(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x20, 0x0d, 0x0d, 0x00},
		CheckResults: map[string]InstructionTestCheck{
			"Expected jump": func(nes *Nes) bool { return nes.CPU.Counter == (0x0d0d)+1 },
		},
		CheckStack: []byte{0x02, 0xc0},
	}
	RunInstructionTest(t, test)
}

// LDA instructions
func TestOpLDA(t *testing.T) {
	Load(t, 0xa5, func(nes *Nes) uint8 { return nes.CPU.Accumulator })
}

// LDX instructions
func TestOpLDX(t *testing.T) {
	Load(t, 0xa6, func(nes *Nes) uint8 { return nes.CPU.IndexX })
}

// LDY instructions
func TestOpLDY(t *testing.T) {
	Load(t, 0xa4, func(nes *Nes) uint8 { return nes.CPU.IndexY })
}

// LSR instructions
func TestOpLSR(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x46, 0x00, 0x00}, Memory: []byte{0b1000_0000},
		CheckFlags:  map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: false, NesCPUStatusCarry: false},
		CheckMemory: []byte{0b0100_0000},
	}
	RunInstructionTest(t, test)

	test.Memory = []byte{0b0000_0001}
	test.CheckFlags[NesCPUStatusZero] = true
	test.CheckFlags[NesCPUStatusCarry] = true
	test.CheckMemory = []byte{0b0000_0000}
	RunInstructionTest(t, test)
}

// ORA instructions
func TestOpORA(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x05, 0x00, 0x00}, Memory: []byte{0b1111_0000},
		InitializeRegs: func(nes *Nes) { nes.CPU.Accumulator = 0b0000_1111 },
		CheckFlags:     map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: true},
		CheckResults: map[string]InstructionTestCheck{
			"Accumulator failed OR": func(nes *Nes) bool { return nes.CPU.Accumulator == 0b1111_1111 },
		},
	}
	RunInstructionTest(t, test)

	// Test zero flag
	test.Memory = []byte{0b0000_0000}
	test.InitializeRegs = func(nes *Nes) { nes.CPU.Accumulator = 0b0000_0000 }
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: true, NesCPUStatusNegative: false}
	test.CheckResults = map[string]InstructionTestCheck{
		"Accumulator failed OR": func(nes *Nes) bool { return nes.CPU.Accumulator == 0b0000_0000 },
	}
	RunInstructionTest(t, test)
}

// PHA instructions
func TestOpPHA(t *testing.T) {
	test := InstructionTest{
		ROM:            []byte{0x48, 0x00},
		InitializeRegs: func(nes *Nes) { nes.CPU.Accumulator = 1 },
		CheckStack:     []byte{1},
	}
	RunInstructionTest(t, test)
}

// PHP instructions
func TestOpPHP(t *testing.T) {
	test := InstructionTest{
		ROM:        []byte{0x08, 0x00},
		Flags:      []NesCPUStatus{NesCPUStatusZero},
		CheckStack: []byte{byte(NesInitialStatus | NesCPUStatusBreak | NesCPUStatusBreak2 | NesCPUStatusZero)},
	}
	RunInstructionTest(t, test)
}

// PLA instructions
func TestOpPLA(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x68, 0x00}, Stack: []byte{1},
		CheckResults: map[string]InstructionTestCheck{
			"Accumulator failed pull": func(nes *Nes) bool { return nes.CPU.Accumulator == 1 },
		},
	}
	RunInstructionTest(t, test)
}

// PLP instructions
func TestOpPLP(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x28, 0x00}, Stack: []byte{byte(NesCPUStatusNegative | NesCPUStatusZero)},
		CheckFlags: map[NesCPUStatus]bool{NesCPUStatusNegative: true, NesCPUStatusZero: true},
	}
	RunInstructionTest(t, test)
}

// ROL instructions
func TestOpROL(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x26, 0x00, 0x00}, Memory: []byte{0b0000_0001},
		CheckFlags:  map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: false, NesCPUStatusCarry: false},
		CheckMemory: []byte{0b0000_0010},
	}
	RunInstructionTest(t, test)

	// Test new carry and zero flag
	test.Memory = []byte{0b1000_0000}
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: true, NesCPUStatusNegative: false, NesCPUStatusCarry: true}
	test.CheckMemory = []byte{0b0000_0000}
	RunInstructionTest(t, test)

	// Test previous carry flag
	test.Flags = []NesCPUStatus{NesCPUStatusCarry}
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: false, NesCPUStatusCarry: true}
	test.CheckMemory = []byte{0b0000_0001}
	RunInstructionTest(t, test)
}

// ROR instructions
func TestOpROR(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x66, 0x00, 0x00}, Memory: []byte{0b1000_0000},
		CheckFlags:  map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: false, NesCPUStatusCarry: false},
		CheckMemory: []byte{0b0100_0000},
	}
	RunInstructionTest(t, test)

	// Test new carry and zero flag
	test.Memory = []byte{0b0000_0001}
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: true, NesCPUStatusNegative: false, NesCPUStatusCarry: true}
	test.CheckMemory = []byte{0b0000_0000}
	RunInstructionTest(t, test)

	// Test previous carry flag
	test.Flags = []NesCPUStatus{NesCPUStatusCarry}
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: true, NesCPUStatusCarry: true}
	test.CheckMemory = []byte{0b1000_0000}
	RunInstructionTest(t, test)
}

// RTI instructions
func TestOpRTI(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x40, 0x00}, Stack: []byte{byte(NesCPUStatusZero | NesCPUStatusNegative), 0x01, 0x80},
		CheckFlags: map[NesCPUStatus]bool{NesCPUStatusZero: true, NesCPUStatusNegative: true},
		CheckResults: map[string]InstructionTestCheck{
			"Program counter failed pull": func(nes *Nes) bool { return nes.CPU.Counter == 0x8001+1 },
		},
	}
	RunInstructionTest(t, test)
}

// RTS instructions
func TestOpRTS(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x60, 0x00}, Stack: []byte{0x00, 0x80},
		CheckResults: map[string]InstructionTestCheck{
			"Program counter failed pull": func(nes *Nes) bool { return nes.CPU.Counter == 0x8001+1 },
		},
	}
	RunInstructionTest(t, test)
}

// SBC instructions
func TestOpSBC(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0xe5, 0x00, 0x00}, Memory: []byte{4},
		Flags:          []NesCPUStatus{NesCPUStatusCarry},
		InitializeRegs: func(nes *Nes) { nes.CPU.Accumulator = 5 },
		CheckFlags:     map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: false, NesCPUStatusCarry: true},
		CheckResults: map[string]InstructionTestCheck{
			"Accumulator failed sub": func(nes *Nes) bool { return nes.CPU.Accumulator == 1 },
		},
	}
	RunInstructionTest(t, test)

	// Test zero flag
	test.Flags = []NesCPUStatus{}
	test.CheckResults = map[string]InstructionTestCheck{
		"Accumulator failed sub (carry)": func(nes *Nes) bool { return nes.CPU.Accumulator == 0 },
	}
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: true, NesCPUStatusNegative: false, NesCPUStatusCarry: true}
	RunInstructionTest(t, test)

	// Test negative flag
	test.Flags = []NesCPUStatus{NesCPUStatusCarry}
	test.Memory = []byte{6}
	test.CheckResults = map[string]InstructionTestCheck{
		"Accumulator failed sub (negative)": func(nes *Nes) bool { return nes.CPU.Accumulator == 255 },
	}
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: true, NesCPUStatusCarry: false}
	RunInstructionTest(t, test)
}

// SEC instructions
func TestOpSEC(t *testing.T) {
	test := InstructionTest{
		ROM:        []byte{0x38, 0x00},
		CheckFlags: map[NesCPUStatus]bool{NesCPUStatusCarry: true},
	}
	RunInstructionTest(t, test)
}

// SED instructions
func TestOpSED(t *testing.T) {
	test := InstructionTest{
		ROM:        []byte{0xf8, 0x00},
		CheckFlags: map[NesCPUStatus]bool{NesCPUStatusDecimal: true},
	}
	RunInstructionTest(t, test)
}

// SEI instructions
func TestOpSEI(t *testing.T) {
	test := InstructionTest{
		ROM:        []byte{0x78, 0x00},
		CheckFlags: map[NesCPUStatus]bool{NesCPUStatusInterruptDisable: true},
	}
	RunInstructionTest(t, test)
}

// STA instructions
func Store(t *testing.T, code byte, InitializeRegs func(nes *Nes), value uint8) {
	test := InstructionTest{
		ROM:            []byte{code, 0x00, 0x00},
		InitializeRegs: InitializeRegs,
		CheckMemory:    []byte{value},
	}
	RunInstructionTest(t, test)
}

// STA instructions
func TestOpSTA(t *testing.T) {
	Store(t, 0x85, func(nes *Nes) { nes.CPU.Accumulator = 1 }, 1)
}

// STX instructions
func TestOpSTX(t *testing.T) {
	Store(t, 0x86, func(nes *Nes) { nes.CPU.IndexX = 1 }, 1)
}

// STY instructions
func TestOpSTY(t *testing.T) {
	Store(t, 0x84, func(nes *Nes) { nes.CPU.IndexY = 1 }, 1)
}

// Base for transfer tests
func Transfer(t *testing.T, code byte, SetSourceValue func(nes *Nes, value uint8), GetSourceValue func(nes *Nes) uint8, GetTargetValue func(nes *Nes) uint8) {
	test := InstructionTest{
		ROM:            []byte{code, 0x00},
		InitializeRegs: func(nes *Nes) { SetSourceValue(nes, 1) },
		CheckResults: map[string]InstructionTestCheck{
			"Register failed transfer": func(nes *Nes) bool { return GetSourceValue(nes) == GetTargetValue(nes) },
		},
		CheckFlags: map[NesCPUStatus]bool{NesCPUStatusZero: false, NesCPUStatusNegative: false},
	}
	RunInstructionTest(t, test)

	test.InitializeRegs = func(nes *Nes) { SetSourceValue(nes, 0) }
	test.CheckFlags = map[NesCPUStatus]bool{NesCPUStatusZero: true, NesCPUStatusNegative: false}
	RunInstructionTest(t, test)
}

// TAX instructions
func TestOpTAX(t *testing.T) {
	Transfer(t, 0xaa, func(nes *Nes, value uint8) { nes.CPU.Accumulator = value }, func(nes *Nes) uint8 { return nes.CPU.Accumulator }, func(nes *Nes) uint8 { return nes.CPU.IndexX })
}

// TAY instructions
func TestOpTAY(t *testing.T) {
	Transfer(t, 0xa8, func(nes *Nes, value uint8) { nes.CPU.Accumulator = value }, func(nes *Nes) uint8 { return nes.CPU.Accumulator }, func(nes *Nes) uint8 { return nes.CPU.IndexY })
}

// TSX instructions
func TestOpTSX(t *testing.T) {
	Transfer(t, 0xba, func(nes *Nes, value uint8) { nes.CPU.Stack = NesStackPointer(value) }, func(nes *Nes) uint8 { return uint8(nes.CPU.Stack) }, func(nes *Nes) uint8 { return nes.CPU.IndexX })
}

// TXA instructions
func TestOpTXA(t *testing.T) {
	Transfer(t, 0x8a, func(nes *Nes, value uint8) { nes.CPU.IndexX = value }, func(nes *Nes) uint8 { return nes.CPU.IndexX }, func(nes *Nes) uint8 { return nes.CPU.Accumulator })
}

// TXS instructions
func TestOpTXS(t *testing.T) {
	test := InstructionTest{
		ROM:            []byte{0x9a, 0x00},
		InitializeRegs: func(nes *Nes) { nes.CPU.IndexX = 1 },
		CheckResults: map[string]InstructionTestCheck{
			"Register failed transfer": func(nes *Nes) bool { return uint8(nes.CPU.Stack) == nes.CPU.IndexX },
		},
	}
	RunInstructionTest(t, test)

	test.InitializeRegs = func(nes *Nes) { nes.CPU.IndexX = 0 }
	test.CheckResults = map[string]InstructionTestCheck{
		"Register failed transfer of zero": func(nes *Nes) bool { return uint8(nes.CPU.Stack) == 0 },
	}
	RunInstructionTest(t, test)
}

// TYA instructions
func TestOpTYA(t *testing.T) {
	Transfer(t, 0x98, func(nes *Nes, value uint8) { nes.CPU.IndexY = value }, func(nes *Nes) uint8 { return nes.CPU.IndexY }, func(nes *Nes) uint8 { return nes.CPU.Accumulator })
}
