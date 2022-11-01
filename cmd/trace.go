package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Trace(nes *Nes) string {
	op := NesInstruction(BusMemoryRead(nes, nes.CPU.Counter))
	entry := nes.CPU.Table.Table[op]

	memAddr, storedVal := NesPointer(0), byte(0)
	if entry.Mode != AddressingModeImmediate && entry.Mode != AddressingModeNoneAddressing {
		memAddr, _ = GetAbsoluteAddress(nes, entry.Mode, nes.CPU.Counter+1)
		storedVal = BusMemoryRead(nes, memAddr)
	}

	tmp := ""
	hex := []string{fmt.Sprintf("%02X", op)}
	switch entry.Length {
	case 1:
		{
			switch entry.Op {
			case 0x0a, 0x4a, 0x2a, 0x6a:
				tmp = "A "
			}
		}
	case 2:
		{
			addr := BusMemoryRead(nes, nes.CPU.Counter+1)
			hex = append(hex, fmt.Sprintf("%02X", addr))

			switch entry.Mode {
			case AddressingModeImmediate:
				tmp = fmt.Sprintf("#$%02X", addr)
			case AddressingModeZeroPage:
				tmp = fmt.Sprintf("$%02X = %02X", memAddr, storedVal)
			case AddressingModeZeroPageX:
				tmp = fmt.Sprintf("$%02X,X @ %02X = %02X", addr, memAddr, storedVal)
			case AddressingModeZeroPageY:
				tmp = fmt.Sprintf("$%02X,Y @ %02X = %02X", addr, memAddr, storedVal)
			case AddressingModeIndirectX:
				tmp = fmt.Sprintf("($%02X,X) @ %02X = %04X = %02X", addr, WrappingAdd8(addr, nes.CPU.IndexX), memAddr, storedVal)
			case AddressingModeIndirectY:
				tmp = fmt.Sprintf("($%02X),Y = %04X @ %04X = %02X", addr, WrappingSubPtr(memAddr, NesPointer(nes.CPU.IndexY)), memAddr, storedVal)
			case AddressingModeNoneAddressing:
				tmp = fmt.Sprintf("$%04X", WrappingAddPtr(nes.CPU.Counter+2, NesPointer(int8(addr))))
			}
		}
	case 3:
		{
			addrLow := BusMemoryRead(nes, nes.CPU.Counter+1)
			addrHigh := BusMemoryRead(nes, nes.CPU.Counter+2)
			addr := BusMemoryReadAddress(nes, nes.CPU.Counter+1)
			hex = append(hex, fmt.Sprintf("%02X", addrLow))
			hex = append(hex, fmt.Sprintf("%02X", addrHigh))

			switch entry.Mode {
			case AddressingModeNoneAddressing:
				if entry.Op == 0x6c {
					tmp = fmt.Sprintf("($%04X) = %04X", addr, GetJumpAddress(nes, addr))
				} else {
					tmp = fmt.Sprintf("$%04X", addr)
				}
			case AddressingModeAbsolute:
				tmp = fmt.Sprintf("$%04X = %02X", memAddr, storedVal)
			case AddressingModeAbsoluteX:
				tmp = fmt.Sprintf("$%04X,X @ %04X = %02X", addr, memAddr, storedVal)
			case AddressingModeAbsoluteY:
				tmp = fmt.Sprintf("$%04X,Y @ %04X = %02X", addr, memAddr, storedVal)
			}
		}
	}

	hexStr := strings.Join(hex[:], " ")
	asmStr := fmt.Sprintf("%04X  %v %v %v", nes.CPU.Counter, PadString(hexStr, 8), PadStringRight(entry.Name, 4), tmp)

	return fmt.Sprintf(
		"%v A:%02X X:%02X Y:%02X P:%02X SP:%02X PPU:%v,%v CYC:%v",
		PadString(asmStr, 47), nes.CPU.Accumulator, nes.CPU.IndexX, nes.CPU.IndexY, nes.CPU.Status, nes.CPU.Stack,
		PadStringRight(strconv.Itoa(int(nes.PPU.Scanline)), 3), PadStringRight(strconv.Itoa(int(nes.PPU.Cycles)), 3), nes.Bus.Cycles,
	)
}

func PadString(s string, n int) string {
	r := []rune(strings.Repeat(" ", n))
	for i, c := range s {
		r[i] = c
	}
	return string(r)
}

func PadStringRight(s string, n int) string {
	r := []rune(strings.Repeat(" ", n))
	for i, c := range s {
		r[n-len(s)+i] = c
	}
	return string(r)
}
