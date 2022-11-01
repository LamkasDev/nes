package main

import "math"

func BitflagSet(b uint8, flag uint8) uint8    { return b | flag }
func BitflagClear(b uint8, flag uint8) uint8  { return b &^ flag }
func BitflagToggle(b uint8, flag uint8) uint8 { return b ^ flag }
func BitflagHas(b uint8, flag uint8) bool     { return b&flag != 0 }

// Conversion uint16 extensions
func WrappingDown16(a uint16) uint8 {
	return uint8(a % (math.MaxUint8 + 1))
}

// Addition uint8 extensions
func WrappingAdd8(a uint8, b uint8) uint8 {
	return uint8((uint16(a) + uint16(b)) % (math.MaxUint8 + 1))
}

func WrappingAdd8Ptr(a uint8, b uint8) NesPointer {
	return NesPointer(WrappingAdd8(a, b))
}

func WrappingAdd8StackPtr(a uint8, b uint8) NesStackPointer {
	return NesStackPointer(WrappingAdd8(a, b))
}

// Addition uint16 extensions
func WrappingAdd16(a uint16, b uint16) uint16 {
	return uint16((uint32(a) + uint32(b)) % (math.MaxUint16 + 1))
}

func WrappingAdd16Ptr(a uint16, b uint16) NesPointer {
	return NesPointer(WrappingAdd16(a, b))
}

func WrappingAddPtr(a NesPointer, b NesPointer) NesPointer {
	return NesPointer((uint32(a) + uint32(b)) % (math.MaxUint16 + 1))
}

// Substraction uint8 extensions
func WrappingSub8(a uint8, b uint8) uint8 {
	res := int16(a) - int16(b)
	if res < 0 {
		return uint8(res + (math.MaxUint8 + 1))
	}
	return uint8(res)
}

func WrappingSub8Int(a int8, b int8) int8 {
	res := int16(a) - int16(b)
	if res < math.MinInt8 {
		return int8(res + (math.MaxInt8 + 1))
	}
	return int8(res)
}

// Substraction uint16 extensions
func WrappingSub16(a uint16, b uint16) uint16 {
	res := int32(a) - int32(b)
	if res < 0 {
		return uint16(res + (math.MaxUint16 + 1))
	}
	return uint16(res)
}

func WrappingSub16Ptr(a uint16, b uint16) NesPointer {
	return NesPointer(WrappingSub16(a, b))
}

func WrappingSubPtr(a NesPointer, b NesPointer) NesPointer {
	return NesPointer(WrappingSub16(uint16(a), uint16(b)))
}

// Negative int8 extensions
func WrappingNeg8(i int8) int8 {
	if i == math.MinInt8 {
		return i
	}
	return -i
}

func WrappingNeg16(i int16) int16 {
	if i == math.MinInt16 {
		return i
	}
	return -i
}
