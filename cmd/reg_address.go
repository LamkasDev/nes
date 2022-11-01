package main

const NesAddressRegisterMirror = 0x3fff

type NesAddressRegister struct {
	value [2]uint8
	high  bool
}

func AddressRegCreate() NesAddressRegister {
	return NesAddressRegister{
		high: true,
	}
}

func AddressRegGet(reg *NesAddressRegister) NesPointer {
	return (NesPointer(reg.value[0]) << 8) | NesPointer(reg.value[1])
}

func AddressRegSet(reg *NesAddressRegister, data NesPointer) {
	if data > NesAddressRegisterMirror {
		AddressRegSetRaw(reg, data&NesAddressRegisterMirror)
	}
}

func AddressRegSetRaw(reg *NesAddressRegister, data NesPointer) {
	reg.value[0] = uint8(data >> 8)
	reg.value[1] = uint8(data & 0xff)
}

func AddressRegUpdate(reg *NesAddressRegister, data uint8) {
	if reg.high {
		reg.value[0] = data
	} else {
		reg.value[1] = data
	}
	AddressRegSet(reg, AddressRegGet(reg))
	reg.high = !reg.high
}

func AddressRegIncrement(reg *NesAddressRegister, n uint8) {
	low := reg.value[1]
	reg.value[1] = WrappingAdd8(reg.value[1], n)
	if low > reg.value[1] {
		reg.value[0] = WrappingAdd8(reg.value[0], 1)
	}
	AddressRegSet(reg, AddressRegGet(reg))
}

func AddressRegReset(reg *NesAddressRegister) {
	reg.high = true
}
