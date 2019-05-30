package app

import "github.com/lishimeng/gpiod/etc"

func Exec() {
	go Listen(etc.Config.Listen.Pin, etc.Config.Listen.File, etc.Config.Listen.Interval)

	go Control(etc.Config.Control.Pin, etc.Config.Control.File, etc.Config.Control.Interval)
}