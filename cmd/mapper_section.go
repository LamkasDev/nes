package main

const (
	NesMapperSectionTypeNull = NesMapperSectionType("NULL")
)

type NesMapperSectionType string
type NesMapperSectionInitialize = func(nes *Nes)
type NesMapperSectionMatches = func(nes *Nes, address NesPointer) bool
type NesMapperSectionGetAddress = func(nes *Nes, address NesPointer) NesPointer
type NesMapperSectionRead = func(nes *Nes, address NesPointer) byte
type NesMapperSectionWrite = func(nes *Nes, address NesPointer, data []byte)
type NesMapperSection struct {
	Type       NesMapperSectionType
	Start      NesPointer
	End        NesPointer
	Initialize NesMapperSectionInitialize
	GetAddress NesMapperSectionGetAddress
	Matches    NesMapperSectionMatches
	Read       NesMapperSectionRead
	Write      NesMapperSectionWrite
}

func CreateMapperSection(t NesMapperSectionType, start NesPointer, end NesPointer) NesMapperSection {
	return NesMapperSection{
		Type:       t,
		Start:      start,
		End:        end,
		GetAddress: func(nes *Nes, address NesPointer) NesPointer { return address },
		Matches:    func(nes *Nes, address NesPointer) bool { return address >= start && address <= end },
	}
}

func GetMapperSectionSize(section NesMapperSection) NesPointer {
	return (section.End + 1) - section.Start
}
