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

	dev, err := libimobiledevice.NewDevice("00008140-001C0C693C08801C")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dev)

	fmt.Println("udid: ", dev.UDID())

	fmt.Println("version: ", dev.Version())

	client, err := libimobiledevice.NewLockdowndClient(dev, "test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(client)
	queryType, err := client.QueryType()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(queryType)

	value, err := client.GetValue("", "DeviceName")
	if err != nil {
		fmt.Println(err)
	}
	defer value.Free()
	fmt.Println(value)

	err = client.SetValue("", "test", nil)
	if err != nil {
		fmt.Println(err)
	}

	service, err := client.StartService("ssh")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(service)

	service, err = client.StartServiceWithEscrowBag("ssh")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(service)

	session, err := client.StartSession("", "", true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(session)

	err = client.Pair(nil)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Unpair(nil)
	if err != nil {
		fmt.Println(err)
	}
}
