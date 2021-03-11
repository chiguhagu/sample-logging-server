package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	l := zap.NewExample()
	l.WithOptions()
	for true {
		l.Error("error log!")
		l.Info("information log!")
		time.Sleep(time.Second * 60)
	}
}
