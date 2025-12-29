package main

import (
	"fmt"

	"github.com/doukn0wdaway/go-xcap-bindings"
)

func main() {
	out, err := xcap.GetMonitors()
	if err != nil {
		panic(err)
	}

	fmt.Println("Monitors:")
	fmt.Println(out)

	ss, err := xcap.GetScreenshot(35)
	if err != nil {
		panic(err)
	}

	fmt.Println("Screenshot:")
	fmt.Println(ss)
}
