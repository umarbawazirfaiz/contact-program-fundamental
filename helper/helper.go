package helper

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Fungsi untuk kembali ke menu
func BackHandler() {
	fmt.Print("Tekan enter untuk kembali ke menu")
	var back int
	fmt.Scanln(&back)
}
