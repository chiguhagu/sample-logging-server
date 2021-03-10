package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for true {
		fmt.Fprintln(os.Stdout, "stdout")
		fmt.Fprintln(os.Stderr, "stderr")
		time.Sleep(time.Second * 60)
	}
}
