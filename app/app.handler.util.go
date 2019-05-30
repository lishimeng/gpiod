package app

import (
	"github.com/lishimeng/mmap"
	"github.com/nathan-osman/go-rpigpio"
)

const (
	LOW = iota
	HIGH
	UNKNOWN
)

func mmapFile(fileName string, fileSize int64) *[]byte {
	data, err := mmap.MapFile(fileName, fileSize)

	if err != nil {
		//fmt.Println(err)
	}

	if data == nil {
		return nil
	}
	return data
}

func closeGpios(pins []*rpi.Pin) {

	//fmt.Println("\nclose gpio ports")
	for _, pin := range pins {

		//fmt.Printf("close pin: %d\n", index)
		if pin != nil {
			err := pin.Close()
			if err != nil {
				//fmt.Printf("fail to close pin: %d\n", index)
				//fmt.Println(err)
			}
		}
	}
}

func getGpios(pins []int, readMode bool) []*rpi.Pin {

	//fmt.Println("\nopen gpio ports")
	container := make([]*rpi.Pin, len(pins))

	var mode rpi.Direction = rpi.OUT
	if readMode {
		mode = rpi.IN
	}
	for index, pinNum := range pins {
		//fmt.Printf("open gpio %d:%d\n", index, pinNum)

		p, err := rpi.OpenPin(pinNum, mode)
		if err != nil {
			//fmt.Printf("can't open gpio %d:%d\n", index, pinNum)
			//fmt.Println(err)
			continue
		}
		container[index] = p
	}
	return container
}
