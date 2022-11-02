package main

func CreateJoypadMapper(t NesMapperSectionType, start NesPointer, index uint8) NesMapperSection {
	section := CreateMapperSection(t, start, start)
	section.Initialize = func(nes *Nes) {
		nes.Joypads = append(nes.Joypads, NesJoypad{Mapping: NesDefaultJoypadMapping})
	}
	section.Read = func(nes *Nes, address NesPointer) byte {
		// nes.Locks.Input.Lock()
		// defer nes.Locks.Input.Unlock()
		return byte(JoypadRead(nes, index))
	}
	section.Write = func(nes *Nes, address NesPointer, data []byte) {
		JoypadWrite(nes, index, NesJoypadButton(data[0]))
	}

	return section
}
