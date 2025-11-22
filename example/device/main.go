package main

import (
	"fmt"

	"github.com/hiifong/libimobiledevice"
	"github.com/hiifong/libimobiledevice/device"
)

func main() {
	fmt.Println("Hello World")
	libimobiledevice.SetDebugLevel(true)
	list, err := device.GetDeviceList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)

	listExtended, err := device.GetDeviceListExtended()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(listExtended)

	dev, err := device.NewDevice("00008140-001C0C693C08801C")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dev)

	fmt.Println("udid: ", dev.UDID())

	fmt.Println("version: ", dev.Version())

	dev, err = device.NewDeviceWithOptions("00008140-001C0C693C08801C", device.IDEVICE_LOOKUP_USBMUX)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dev)

	connect, err := device.Connect(dev, 80)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		err := connect.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	n, err := connect.Send([]byte("Hello World"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	buf := make([]byte, 1024)
	n, err = connect.Receive(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf[:n]))

	n, err = connect.ReceiveWithTimeout(buf, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf[:n]))
}
