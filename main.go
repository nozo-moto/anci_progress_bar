package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	err := exec.Command("tput", "civis").Run()
	if err != nil {
		panic(err)
	}
	length := 96
	background := ""
	text := []string{
		"▏", "▎", "▍", "▌", "▋", "▊", "▉", "█",
	}
	fmt.Printf("\033[43m")
	for i := 0; i < length/len(text); i++ {
		background += " "
	}
	fmt.Printf(background)

	fmt.Printf("\r\033[43m")
	for i := 0; i < length; i++ {
		fmt.Print("\r\033[36m")
		for j := 0; j < i/len(text); j++ {
			fmt.Print("█")
		}
		fmt.Printf("%s", text[i%len(text)])
		time.Sleep(time.Millisecond * 50)
	}
	//	for _, s := range text {
	//		fmt.Printf("\r%s", s)
	//		time.Sleep(time.Millisecond * 100)
	//	}

	fmt.Printf("\n")

	err = exec.Command("tput", "cnorm").Run()
	if err != nil {
		panic(err)
	}
}
