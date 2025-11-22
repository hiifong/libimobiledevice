package main

import (
	"fmt"

	"github.com/hiifong/libimobiledevice"
)

func main() {
	fmt.Println("Hello World")
	libimobiledevice.SetDebugLevel(true)
	list, err := libimobiledevice.GetDeviceList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)

	listExtended, err := libimobiledevice.GetDeviceListExtended()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(listExtended)

	device, err := libimobiledevice.NewDevice("00008140-001C0C693C08801C")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(device)

	fmt.Println("udid: ", device.UDID())

	fmt.Println("version: ", device.Version())

	device, err = libimobiledevice.NewDeviceWithOptions("00008140-001C0C693C08801C", libimobiledevice.IDEVICE_LOOKUP_USBMUX)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(device)

	connect, err := libimobiledevice.Connect(device, 80)
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
