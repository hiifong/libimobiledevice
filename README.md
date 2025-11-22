# libimobiledevice-go

[![Go Reference](https://pkg.go.dev/badge/github.com/hiifong/libimobiledevice.svg)](https://pkg.go.dev/github.com/hiifong/libimobiledevice)

> WIP: Go bindings for libimobiledevice

## Prerequisites

Ubuntu:
```go
sudo apt-get install \
    libplist-dev \
	libusbmuxd-dev \
	libimobiledevice-glue-dev \
	libtatsu-dev \
	libssl-dev \
	usbmuxd
```

[For more information, please see the official repository.](https://github.com/libimobiledevice/libimobiledevice)

## Install
```shell
go get github.com/hiifong/libimobiledevice
```

## Usage
- Enable debug
```go
libimobiledevice.SetDebugLevel(true)
```

- Get a list of UDIDs of currently available devices (USBMUX devices only).
```go
udids, err := libimobiledevice.GetDeviceList()
if err != nil {
    fmt.Println(err)
}
fmt.Println(udids)
```
- Connect to device with udid
```go
device, err := libimobiledevice.NewDevice("00008140-001C0C693C08801C")
if err != nil {
    fmt.Println(err)
}
fmt.Println(device)
```

## Contributing
Contributions are welcome! To contribute:

1. Fork the repository 
2. Create a feature branch (git checkout -b feature/amazing-feature)
3. Commit your changes (git commit -m 'Add amazing feature')
4. Push to the branch (git push origin feature/amazing-feature)
5. Open a Pull Request