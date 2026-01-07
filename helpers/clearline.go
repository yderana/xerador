package helpers

import "fmt"

func ClearPrevLines(n int) {
	// naik n baris, lalu clear line
	for i := 0; i < n; i++ {
		fmt.Print("\033[1A") // cursor up 1 line
		fmt.Print("\r\033[2K")
	}
	// clear current line juga (jaga-jaga)
	fmt.Print("\r\033[2K")
}

func ClearPromptLine() {
	// clear current line
	fmt.Print("\r\033[2K")
	// naik 1 baris, clear juga (untuk kasus promptui bikin line terpisah)
	fmt.Print("\033[1A\r\033[2K")
	// kembali ke bawah (biar posisi normal)
	fmt.Print("\033[1B")
}

func ClearLine() {
	fmt.Print("\r\033[2K")
}
