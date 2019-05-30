package app

import (
	"github.com/lishimeng/mmap"
	"github.com/nathan-osman/go-rpigpio"
	"time"
)


func Listen(pinNums []int, fileName string, interval int) {

	if pinNums == nil || len(pinNums) == 0 {
		return
	}

	data, err := mmap.MapFile(fileName, int64(len(pinNums)))

	if err != nil {
		//fmt.Println(err)
	}

	if data == nil {
		return
	}
	defer mmap.UnmapFile(data)

	//fmt.Println("write data")
	d := *data
	for i := 0; i < len(pinNums); i++ {
		d[i] = UNKNOWN
	}
	pins := getGpios(pinNums, true)
	defer closeGpios(pins)

	for {
		cacheGpios(d, pins)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func cacheGpios(datas []byte, pins []*rpi.Pin) {

	//fmt.Println("\nget gpio values")
	for index, pin := range pins {
		//fmt.Printf("get gpio value %d\n", index)
		if pin == nil {
			//fmt.Println("pin is nil")
			continue
		}
		value, err := pin.Read()
		if err != nil {
			//fmt.Println(err)
			datas[index] = UNKNOWN
		} else {
			val := int(value)
			//fmt.Printf("pin val: %d\n", val)
			if val == LOW {
				datas[index] = LOW
			} else {
				datas[index] = HIGH
			}
		}
	}
}