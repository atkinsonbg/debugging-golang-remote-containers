package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	args := []string{"gobook.pdf"}
	cmd := exec.Command("exiftool", args...)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err := cmd.Run()
	if err != nil {
		fmt.Println(&errb)
	}
	fmt.Println(&outb)
}