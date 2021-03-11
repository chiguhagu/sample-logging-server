package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for true {
		fmt.Fprintln(os.Stdout, "{\"msg\":\"stdout\"}")
		fmt.Fprintln(os.Stderr, "{\"msg\":\"stderr\"}")
		time.Sleep(time.Second * 60)
	}
}
