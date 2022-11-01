package main

const NesRAMRandom = NesPointer(0xfe)
const NesRAMKeys = NesPointer(0xff)
const NesRAMStackStart = NesPointer(0x100)
const NesRAMStackEnd = NesRAMStackStart + NesPointer(NesStackReset)
