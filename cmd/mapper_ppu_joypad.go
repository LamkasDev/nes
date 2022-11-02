package main

func CreateJoypadMapper(t NesMapperSectionType, start NesPointer, index uint8) NesMapperSection {
	section := CreateMapperSection(t, start, start)
	section.Read = func(nes *Nes, address NesPointer) byte {
		return byte(JoypadRead(&nes.Joypads[index]))
	}
	section.Write = func(nes *Nes, address NesPointer, data []byte) {
		JoypadWrite(&nes.Joypads[index], NesJoypadButton(data[0]))
	}

	return section
}
