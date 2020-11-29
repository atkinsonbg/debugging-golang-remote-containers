package main

import (
	//"bytes"
	"fmt"
	"os/exec"
)

func main() {
	args := []string{"gobook2.pdf"}
	cmd := exec.Command("exiftool", args...)

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}