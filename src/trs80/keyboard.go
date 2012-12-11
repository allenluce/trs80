package main

// http://www.trs-80.com/trs80-zaps-internals.htm#keyboard13
const keyboardBegin = 0x3800
const keyboardEnd = keyboardBegin + 256

func (cpu *cpu) readKeyboard(addr word) byte {
	addr -= keyboardBegin

	var b byte

	if addr == 0x40 {
		b = 0x04 // break
	} else {
		// No keys pressed.
		b = 0
	}

	return b
}