package helper

import (
	"contact-program/model"
	"fmt"
	"os"
	"os/exec"
	"strings"
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

func PhoneToString(phoneDatas []model.PhoneData) string {
	var stringPhone []string
	for _, v := range phoneDatas {
		stringPhone = append(stringPhone, *v.GetPhone())
	}

	return strings.Join(stringPhone, ", ")
}

func PhonesToPhoneDatas(phones []string) []model.PhoneData {
	phoneDatas := []model.PhoneData{}
	for _, v := range phones {
		phoneData := model.PhoneData{}
		phoneData.SetPhone(&v)
		phoneDatas = append(phoneDatas, phoneData)
	}

	return phoneDatas
}
