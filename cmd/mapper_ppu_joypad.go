package main

func CreateJoypadMapper(t NesMapperSectionType, start NesPointer) NesMapperSection {
	section := CreateMapperSection(t, start, start)
	section.Read = func(nes *Nes, address NesPointer) byte { return 0 }
	section.Write = func(nes *Nes, address NesPointer, data []byte) {}

	return section
}
