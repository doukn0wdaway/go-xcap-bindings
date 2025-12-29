package main

import (
	"fmt"
	"os/exec"
)

func main() {
	command, err := exec.Command("./xcap-cli/target/release/xcap-cli", "screenshot 35").Output()

	if err != nil {
		fmt.Printf("error:", err.Error())
	}

	fmt.Printf("string(command): %v\n", string(command))
}
