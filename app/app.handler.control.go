package app

import (
	"github.com/lishimeng/mmap"
	"github.com/nathan-osman/go-rpigpio"
	"time"
)

func Control(pinNums []int, fileName string, interval int) {

	if pinNums == nil || len(pinNums) == 0 {
		return
	}

	data := mmapFile(fileName, int64(len(pinNums)))
	defer mmap.UnmapFile(data)

	//fmt.Println("write data")
	d := *data
	for i := 0; i < len(pinNums); i++ {
		d[i] = UNKNOWN
	}
	pins := getGpios(pinNums, false)
	defer closeGpios(pins)

	for {
		controlGpios(d, pins)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func controlGpios(datas []byte, pins []*rpi.Pin) {

	for i, d := range datas {
		switch d {
			case HIGH:
				refreshPin(pins, i, HIGH)
				break
			case LOW:
				refreshPin(pins, i, LOW)
				break
			case UNKNOWN:
				break
			default:
				break
		}
	}
}

func refreshPin(pins []*rpi.Pin, index int, val int) {

	if pins == nil || len(pins) <= index {
		return
	}
	pin := pins[index]
	pin.Write(rpi.Value(val))
}