package libimobiledevice

/*
#include <stdlib.h>
#include <libimobiledevice/libimobiledevice.h>
*/
import "C"
import (
	"unsafe"
)

// DeviceErr Error Codes
type DeviceErr int

func (e DeviceErr) Error() string {
	return C.GoString(C.idevice_strerror(C.idevice_error_t(e)))
}

func checkErr(err C.idevice_error_t) error {
	if err != C.IDEVICE_E_SUCCESS {
		return DeviceErr(err)
	}
	return nil
}

const (
	IDEVICE_E_SUCCESS         DeviceErr = C.IDEVICE_E_SUCCESS
	IDEVICE_E_INVALID_ARG     DeviceErr = C.IDEVICE_E_INVALID_ARG
	IDEVICE_E_UNKNOWN_ERROR   DeviceErr = C.IDEVICE_E_UNKNOWN_ERROR
	IDEVICE_E_NO_DEVICE       DeviceErr = C.IDEVICE_E_NO_DEVICE
	IDEVICE_E_NOT_ENOUGH_DATA DeviceErr = C.IDEVICE_E_NOT_ENOUGH_DATA
	IDEVICE_E_CONNREFUSED     DeviceErr = C.IDEVICE_E_CONNREFUSED
	IDEVICE_E_SSL_ERROR       DeviceErr = C.IDEVICE_E_SSL_ERROR
	IDEVICE_E_TIMEOUT         DeviceErr = C.IDEVICE_E_TIMEOUT
)

// DeviceOpt Options for idevice_new_with_options()
type DeviceOpt int

const (
	IDEVICE_LOOKUP_USBMUX         DeviceOpt = C.IDEVICE_LOOKUP_USBMUX         // include USBMUX devices during lookup
	IDEVICE_LOOKUP_NETWORK        DeviceOpt = C.IDEVICE_LOOKUP_NETWORK        // include network devices during lookup
	IDEVICE_LOOKUP_PREFER_NETWORK DeviceOpt = C.IDEVICE_LOOKUP_PREFER_NETWORK // prefer network connection if device is available via USBMUX *and* network
)

// DeviceConnType Type of connection a device is available on
type DeviceConnType int

const (
	CONNECTION_USBMUXD DeviceConnType = C.CONNECTION_USBMUXD // device is available via USBMUX
	CONNECTION_NETWORK DeviceConnType = C.CONNECTION_NETWORK // device is available via network
)

// DeviceEventType
// discovery (events/asynchronous)
// The event type for device add or removal
type DeviceEventType int

const (
	IDEVICE_DEVICE_ADD    DeviceEventType = C.IDEVICE_DEVICE_ADD    // device was added
	IDEVICE_DEVICE_REMOVE DeviceEventType = C.IDEVICE_DEVICE_REMOVE // device was removed
	IDEVICE_DEVICE_PAIRED DeviceEventType = C.IDEVICE_DEVICE_PAIRED // device completed pairing process
)

type Device struct {
	dev *C.idevice_t
}

// SetDebugLevel Set the level of debugging.  Set to false for no debug output or true to enable debug output.
func SetDebugLevel(enable bool) {
	if enable {
		C.idevice_set_debug_level(C.int(1))
	} else {
		C.idevice_set_debug_level(C.int(0))
	}
}

// GetDeviceList Get a list of UDIDs of currently available devices (USBMUX devices only).
func GetDeviceList() ([]string, error) {
	var cDevices **C.char
	var cCount C.int

	err := C.idevice_get_device_list(&cDevices, &cCount)
	if err := checkErr(err); err != nil {
		return nil, err
	}
	defer C.idevice_device_list_free(cDevices)

	// see https://go.dev/wiki/cgo#turning-c-arrays-into-go-slices
	cSlice := unsafe.Slice((*C.char)(*cDevices), int(cCount))
	devices := make([]string, cCount)

	for i, cStr := range cSlice {
		devices[i] = C.GoString(&cStr)
	}

	return devices, nil
}

// GetDeviceListExtended Get a list of currently available devices
func GetDeviceListExtended() ([]DeviceInfo, error) {
	var cDevices *C.idevice_info_t
	var cCount C.int

	err := C.idevice_get_device_list_extended(&cDevices, &cCount)
	if err := checkErr(err); err != nil {
		return nil, err
	}
	defer C.idevice_device_list_extended_free(cDevices)

	cSlice := unsafe.Slice((*C.idevice_info_t)(cDevices), int(cCount))
	infos := make([]DeviceInfo, int(cCount))

	for i, info := range cSlice {
		cinfo := (C.idevice_info_t)(unsafe.Pointer(info))
		infos[i] = DeviceInfo{
			UDID:     C.GoString(cinfo.udid),
			ConnType: DeviceConnType(cinfo.conn_type),
			ConnData: unsafe.Pointer(cinfo.conn_data),
		}
	}

	return infos, nil
}

// NewDevice  Creates an idevice_t structure for the device specified by UDID,
// if the device is available (USBMUX devices only).
func NewDevice(udid string) (*Device, error) {
	var cDevice C.idevice_t
	cUDID := C.CString(udid)
	defer C.free(unsafe.Pointer(cUDID))

	err := C.idevice_new(&cDevice, cUDID)
	if err := checkErr(err); err != nil {
		return nil, err
	}

	return &Device{&cDevice}, nil
}

// NewDeviceWithOptions Creates an idevice_t structure for the device specified by UDID,
// if the device is available, with the given lookup options.
func NewDeviceWithOptions(udid string, opt DeviceOpt) (*Device, error) {
	var cDevice C.idevice_t
	cUDID := C.CString(udid)
	defer C.free(unsafe.Pointer(cUDID))

	err := C.idevice_new_with_options(&cDevice, cUDID, C.enum_idevice_options(opt))
	if err := checkErr(err); err != nil {
		return nil, err
	}

	return &Device{&cDevice}, nil
}

// DeviceFree free device
func DeviceFree(device *Device) error {
	err := C.idevice_free(*device.dev)
	return checkErr(err)
}

// UDID Gets the Unique Device ID for the device.
func (d *Device) UDID() string {
	var cUdid *C.char
	defer C.free(unsafe.Pointer(cUdid))

	err := checkErr(C.idevice_get_udid(*d.dev, &cUdid))
	if err != nil {
		return ""
	}

	return C.GoString(cUdid)
}

func (d *Device) Version() int {
	return int(C.idevice_get_device_version(*d.dev))
}

type DeviceInfo struct {
	UDID     string
	ConnType DeviceConnType
	ConnData unsafe.Pointer
}

type DeviceEvent struct {
	EventType DeviceEventType
	UDID      string
	ConnType  DeviceConnType
}

type DeviceConnection struct {
	conn *C.idevice_connection_t
}

// Connect Set up a connection to the given device.
func Connect(device *Device, port int) (*DeviceConnection, error) {
	var cConn C.idevice_connection_t

	err := C.idevice_connect(*device.dev, C.uint16_t(port), &cConn)
	if err := checkErr(err); err != nil {
		return nil, err
	}

	return &DeviceConnection{&cConn}, nil
}

// Send send data to a device via the given connection.
func (dc *DeviceConnection) Send(data []byte) (int, error) {
	var cSent C.uint32_t
	cData := (*C.char)(unsafe.Pointer(&data[0]))
	cLen := C.uint32_t(len(data))

	err := C.idevice_connection_send(C.idevice_connection_t(*dc.conn), cData, cLen, &cSent)
	if err := checkErr(err); err != nil {
		return 0, err
	}

	return int(cSent), nil
}

// Receive receive data from a device via the given connection.
// This function is like idevice_connection_receive_timeout, but with a
// predefined reasonable timeout.
func (dc *DeviceConnection) Receive(buf []byte) (int, error) {
	var cReceive C.uint32_t
	cBuf := (*C.char)(unsafe.Pointer(&buf[0]))
	cLen := C.uint32_t(len(buf))

	err := C.idevice_connection_receive(C.idevice_connection_t(*dc.conn), cBuf, cLen, &cReceive)
	if err := checkErr(err); err != nil {
		return 0, err
	}

	return int(cReceive), nil
}

func (dc *DeviceConnection) ReceiveWithTimeout(buf []byte, timeout int) (int, error) {
	var cReceive C.uint32_t
	cBuf := (*C.char)(unsafe.Pointer(&buf[0]))
	cLen := C.uint32_t(len(buf))

	err := C.idevice_connection_receive_timeout(C.idevice_connection_t(*dc.conn), cBuf, cLen, &cReceive, C.uint32_t(timeout))
	if err := checkErr(err); err != nil {
		return 0, err
	}

	return int(cReceive), nil
}

// EnableSSL Enables SSL for the given connection.
func (dc *DeviceConnection) EnableSSL() error {
	err := C.idevice_connection_enable_ssl(*dc.conn)
	return checkErr(err)
}

// DisableSSL Disable SSL for the given connection.
func (dc *DeviceConnection) DisableSSL() error {
	err := C.idevice_connection_disable_ssl(*dc.conn)
	return checkErr(err)
}

// DisableBypassSSL Disable bypass SSL for the given connection without sending out terminate messages.
func (dc *DeviceConnection) DisableBypassSSL(bypass int) error {
	cSSLBypass := C.uint8_t(bypass)

	err := C.idevice_connection_disable_bypass_ssl(*dc.conn, cSSLBypass)

	return checkErr(err)
}

// Close Disconnect from the device and clean up the connection structure.
func (dc *DeviceConnection) Close() error {
	if dc == nil || dc.conn == nil {
		return nil
	}

	err := C.idevice_disconnect(*dc.conn)
	return checkErr(err)
}

type DeviceSubscriptionContext C.idevice_subscription_context_t

// LibimobiledeviceVersion Returns a static string of the libimobiledevice version.
func LibimobiledeviceVersion() string {
	return C.GoString(C.libimobiledevice_version())
}
