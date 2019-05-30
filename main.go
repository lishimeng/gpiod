package main

import (
	"github.com/lishimeng/gpiod/app"
	"time"
)

func main() {

	app.Exec()

	for {
		time.Sleep(10 * time.Second)
	}
}
