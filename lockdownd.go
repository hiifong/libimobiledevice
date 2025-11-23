package libimobiledevice

/*
#include <stdlib.h>
#include <libimobiledevice/lockdown.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

// LockdowndErr Error Codes
type LockdowndErr int

func (e LockdowndErr) Error() string {
	return C.GoString(C.lockdownd_strerror(C.lockdownd_error_t(e)))
}

func checkLockdowndErr(err C.lockdownd_error_t) error {
	if err != C.LOCKDOWN_E_SUCCESS {
		return LockdowndErr(err)
	}
	return nil
}

const (
	// custom
	LOCKDOWN_E_SUCCESS            LockdowndErr = C.LOCKDOWN_E_SUCCESS
	LOCKDOWN_E_INVALID_ARG        LockdowndErr = C.LOCKDOWN_E_INVALID_ARG
	LOCKDOWN_E_INVALID_CONF       LockdowndErr = C.LOCKDOWN_E_INVALID_CONF
	LOCKDOWN_E_PLIST_ERROR        LockdowndErr = C.LOCKDOWN_E_PLIST_ERROR
	LOCKDOWN_E_PAIRING_FAILED     LockdowndErr = C.LOCKDOWN_E_PAIRING_FAILED
	LOCKDOWN_E_SSL_ERROR          LockdowndErr = C.LOCKDOWN_E_SSL_ERROR
	LOCKDOWN_E_DICT_ERROR         LockdowndErr = C.LOCKDOWN_E_DICT_ERROR
	LOCKDOWN_E_RECEIVE_TIMEOUT    LockdowndErr = C.LOCKDOWN_E_RECEIVE_TIMEOUT
	LOCKDOWN_E_MUX_ERROR          LockdowndErr = C.LOCKDOWN_E_MUX_ERROR
	LOCKDOWN_E_NO_RUNNING_SESSION LockdowndErr = C.LOCKDOWN_E_NO_RUNNING_SESSION

	// native
	LOCKDOWN_E_INVALID_RESPONSE                        LockdowndErr = C.LOCKDOWN_E_INVALID_RESPONSE
	LOCKDOWN_E_MISSING_KEY                             LockdowndErr = C.LOCKDOWN_E_MISSING_KEY
	LOCKDOWN_E_MISSING_VALUE                           LockdowndErr = C.LOCKDOWN_E_MISSING_VALUE
	LOCKDOWN_E_GET_PROHIBITED                          LockdowndErr = C.LOCKDOWN_E_GET_PROHIBITED
	LOCKDOWN_E_SET_PROHIBITED                          LockdowndErr = C.LOCKDOWN_E_SET_PROHIBITED
	LOCKDOWN_E_REMOVE_PROHIBITED                       LockdowndErr = C.LOCKDOWN_E_REMOVE_PROHIBITED
	LOCKDOWN_E_IMMUTABLE_VALUE                         LockdowndErr = C.LOCKDOWN_E_IMMUTABLE_VALUE
	LOCKDOWN_E_PASSWORD_PROTECTED                      LockdowndErr = C.LOCKDOWN_E_PASSWORD_PROTECTED
	LOCKDOWN_E_USER_DENIED_PAIRING                     LockdowndErr = C.LOCKDOWN_E_USER_DENIED_PAIRING
	LOCKDOWN_E_PAIRING_DIALOG_RESPONSE_PENDING         LockdowndErr = C.LOCKDOWN_E_PAIRING_DIALOG_RESPONSE_PENDING
	LOCKDOWN_E_MISSING_HOST_ID                         LockdowndErr = C.LOCKDOWN_E_MISSING_HOST_ID
	LOCKDOWN_E_INVALID_HOST_ID                         LockdowndErr = C.LOCKDOWN_E_INVALID_HOST_ID
	LOCKDOWN_E_SESSION_ACTIVE                          LockdowndErr = C.LOCKDOWN_E_SESSION_ACTIVE
	LOCKDOWN_E_SESSION_INACTIVE                        LockdowndErr = C.LOCKDOWN_E_SESSION_INACTIVE
	LOCKDOWN_E_MISSING_SESSION_ID                      LockdowndErr = C.LOCKDOWN_E_MISSING_SESSION_ID
	LOCKDOWN_E_INVALID_SESSION_ID                      LockdowndErr = C.LOCKDOWN_E_INVALID_SESSION_ID
	LOCKDOWN_E_MISSING_SERVICE                         LockdowndErr = C.LOCKDOWN_E_MISSING_SERVICE
	LOCKDOWN_E_INVALID_SERVICE                         LockdowndErr = C.LOCKDOWN_E_INVALID_SERVICE
	LOCKDOWN_E_SERVICE_LIMIT                           LockdowndErr = C.LOCKDOWN_E_SERVICE_LIMIT
	LOCKDOWN_E_MISSING_PAIR_RECORD                     LockdowndErr = C.LOCKDOWN_E_MISSING_PAIR_RECORD
	LOCKDOWN_E_SAVE_PAIR_RECORD_FAILED                 LockdowndErr = C.LOCKDOWN_E_SAVE_PAIR_RECORD_FAILED
	LOCKDOWN_E_INVALID_PAIR_RECORD                     LockdowndErr = C.LOCKDOWN_E_INVALID_PAIR_RECORD
	LOCKDOWN_E_INVALID_ACTIVATION_RECORD               LockdowndErr = C.LOCKDOWN_E_INVALID_ACTIVATION_RECORD
	LOCKDOWN_E_MISSING_ACTIVATION_RECORD               LockdowndErr = C.LOCKDOWN_E_MISSING_ACTIVATION_RECORD
	LOCKDOWN_E_SERVICE_PROHIBITED                      LockdowndErr = C.LOCKDOWN_E_SERVICE_PROHIBITED
	LOCKDOWN_E_ESCROW_LOCKED                           LockdowndErr = C.LOCKDOWN_E_ESCROW_LOCKED
	LOCKDOWN_E_PAIRING_PROHIBITED_OVER_THIS_CONNECTION LockdowndErr = C.LOCKDOWN_E_PAIRING_PROHIBITED_OVER_THIS_CONNECTION
	LOCKDOWN_E_FMIP_PROTECTED                          LockdowndErr = C.LOCKDOWN_E_FMIP_PROTECTED
	LOCKDOWN_E_MC_PROTECTED                            LockdowndErr = C.LOCKDOWN_E_MC_PROTECTED
	LOCKDOWN_E_MC_CHALLENGE_REQUIRED                   LockdowndErr = C.LOCKDOWN_E_MC_CHALLENGE_REQUIRED
	LOCKDOWN_E_UNKNOWN_ERROR                           LockdowndErr = C.LOCKDOWN_E_UNKNOWN_ERROR
)

// LockdowndCuPairingCb Callback types used in #lockdownd_cu_pairing_cb_t
type LockdowndCuPairingCb int

const (
	// LOCKDOWN_CU_PAIRING_PIN_REQUESTED
	// PIN requested: data_ptr is a char* buffer, and data_size points to the size of
	// this buffer that must not be exceeded and has to be updated to the actual number
	// of characters filled into the buffer.
	LOCKDOWN_CU_PAIRING_PIN_REQUESTED LockdowndCuPairingCb = C.LOCKDOWN_CU_PAIRING_PIN_REQUESTED

	// LOCKDOWN_CU_PAIRING_DEVICE_INFO
	// device information available: data_ptr is a plist_t, and data_size is ignored.
	// The plist_t has to be copied if required, since it is freed when the callback
	// function returns.
	LOCKDOWN_CU_PAIRING_DEVICE_INFO LockdowndCuPairingCb = C.LOCKDOWN_CU_PAIRING_DEVICE_INFO

	// LOCKDOWN_CU_PAIRING_ERROR
	// pairing error message available: data_ptr is a NULL-terminated char* buffer containing
	// the error message, and data_size is ignored. Buffer needs to be copied if it shall
	// persist outside the callback.
	LOCKDOWN_CU_PAIRING_ERROR LockdowndCuPairingCb = C.LOCKDOWN_CU_PAIRING_ERROR
)

// LockdowndClient The ptr handle.
type LockdowndClient struct {
	ptr *C.lockdownd_client_t
}

// NewLockdowndClient Creates a new lockdownd ptr for the device.
func NewLockdowndClient(dev *Device, label string, handshake ...bool) (*LockdowndClient, error) {
	if dev == nil || dev.ptr == nil {
		return nil, errors.New("device is required")
	}

	var client C.lockdownd_client_t
	cLabel := C.CString(label)
	defer C.free(unsafe.Pointer(cLabel))

	if len(handshake) > 0 && handshake[0] {
		err := C.lockdownd_client_new_with_handshake(*dev.ptr, &client, cLabel)
		if err := checkLockdowndErr(err); err != nil {
			return nil, err
		}
	} else {
		err := C.lockdownd_client_new(*dev.ptr, &client, cLabel)
		if err := checkLockdowndErr(err); err != nil {
			return nil, err
		}
	}

	return &LockdowndClient{ptr: &client}, nil
}

// Close free lockdown ptr
func (c *LockdowndClient) Close() error {
	if c == nil || c.ptr == nil {
		return errors.New("ptr is nil")
	}
	err := C.lockdownd_client_free(*c.ptr)

	return checkLockdowndErr(err)
}

// QueryType
// Query the type of the service daemon. Depending on whether the device is
// queried in normal mode or restore mode, different types will be returned.
func (c *LockdowndClient) QueryType() (string, error) {
	if c == nil || c.ptr == nil {
		return "", LOCKDOWN_E_INVALID_ARG
	}

	var cTyp *C.char
	defer C.free(unsafe.Pointer(cTyp))

	err := C.lockdownd_query_type(*c.ptr, &cTyp)
	if err := checkLockdowndErr(err); err != nil {
		return "", err
	}

	if cTyp == nil {
		return "", nil
	}

	typ := C.GoString(cTyp)

	return typ, nil
}

// GetValue
// Retrieves a preferences plist using an optional domain and/or key name.
func (c *LockdowndClient) GetValue(domain, key string) (*PList, error) {
	if c == nil || c.ptr == nil {
		return nil, LOCKDOWN_E_INVALID_ARG
	}

	var cDomain *C.char
	if domain != "" {
		cDomain = C.CString(domain)
		defer C.free(unsafe.Pointer(cDomain))
	}
	var cKey *C.char
	if key != "" {
		cKey = C.CString(key)
		defer C.free(unsafe.Pointer(cKey))
	}
	var cPlist C.plist_t

	err := C.lockdownd_get_value(*c.ptr, cDomain, cKey, &cPlist)
	if err := checkLockdowndErr(err); err != nil {
		return nil, err
	}

	return &PList{ptr: cPlist}, nil
}

// SetValue
// Sets a preferences value using a plist and optional by domain and/or key name.
func (c *LockdowndClient) SetValue(domain, key string, value *PList) error {
	if c == nil || c.ptr == nil || value == nil {
		return LOCKDOWN_E_INVALID_ARG
	}

	cDomain := C.CString(domain)
	cKey := C.CString(key)
	cValue := C.plist_t(unsafe.Pointer(value.ptr))
	defer C.free(unsafe.Pointer(cDomain))
	defer C.free(unsafe.Pointer(cKey))

	err := C.lockdownd_set_value(*c.ptr, cDomain, cKey, cValue)
	if err := checkLockdowndErr(err); err != nil {
		return err
	}

	return nil
}

// RemoveValue Removes a preference node by domain and/or key name.
func (c *LockdowndClient) RemoveValue(domain, key string) error {
	if c == nil || c.ptr == nil {
		return LOCKDOWN_E_INVALID_ARG
	}

	var cDomain *C.char
	if domain != "" {
		cDomain = C.CString(domain)
		defer C.free(unsafe.Pointer(cDomain))
	}
	var cKey *C.char
	if key != "" {
		cKey = C.CString(key)
		defer C.free(unsafe.Pointer(cKey))
	}

	err := C.lockdownd_remove_value(*c.ptr, cDomain, cKey)
	if err := checkLockdowndErr(err); err != nil {
		return err
	}

	return nil
}

// StartService Requests to start a service and retrieve it's port on success.
func (c *LockdowndClient) StartService(identifier string) (*LockdowndServiceDescriptor, error) {
	if c == nil || c.ptr == nil {
		return nil, LOCKDOWN_E_INVALID_ARG
	}

	var cIdentifier *C.char
	if identifier != "" {
		cIdentifier = C.CString(identifier)
		defer C.free(unsafe.Pointer(cIdentifier))
	}

	var service C.lockdownd_service_descriptor_t

	err := C.lockdownd_start_service(*c.ptr, cIdentifier, &service)
	if err := checkLockdowndErr(err); err != nil {
		return nil, err
	}

	return &LockdowndServiceDescriptor{
		port:       int(service.port),
		sslEnabled: int(service.ssl_enabled),
		identifier: C.GoString(service.identifier),
	}, nil
}

// StartServiceWithEscrowBag Requests to start a service and retrieve it's port on success.
// Sends the escrow bag from the device's pair record.
func (c *LockdowndClient) StartServiceWithEscrowBag(identifier string) (*LockdowndServiceDescriptor, error) {
	if c == nil || c.ptr == nil {
		return nil, LOCKDOWN_E_INVALID_ARG
	}

	var cIdentifier *C.char
	if identifier != "" {
		cIdentifier = C.CString(identifier)
		defer C.free(unsafe.Pointer(cIdentifier))
	}

	var service C.lockdownd_service_descriptor_t

	err := C.lockdownd_start_service_with_escrow_bag(*c.ptr, cIdentifier, &service)
	if err := checkLockdowndErr(err); err != nil {
		return nil, err
	}

	return &LockdowndServiceDescriptor{
		port:       int(service.port),
		sslEnabled: int(service.ssl_enabled),
		identifier: C.GoString(service.identifier),
	}, nil
}

// StartSession Opens a session with lockdownd and switches to SSL mode if device wants it.
func (c *LockdowndClient) StartSession(hostId, sessionId string, sslEnabled bool) (string, error) {
	if c == nil || c.ptr == nil {
		return "", LOCKDOWN_E_INVALID_ARG
	}

	var cHostId *C.char
	if hostId != "" {
		cHostId = C.CString(hostId)
		defer C.free(unsafe.Pointer(cHostId))
	}

	var cSessionId *C.char
	if sessionId != "" {
		cSessionId = C.CString(sessionId)
		defer C.free(unsafe.Pointer(cSessionId))
	}

	var cSslEnabled C.int
	if sslEnabled {
		cSslEnabled = 1
	} else {
		cSslEnabled = 0
	}

	err := C.lockdownd_start_session(*c.ptr, cHostId, &cSessionId, &cSslEnabled)
	if err := checkLockdowndErr(err); err != nil {
		return "", err
	}

	return C.GoString(cSessionId), nil
}

// StopSession Closes the lockdownd session by sending the StopSession request.
func (c *LockdowndClient) StopSession(sessionId string) error {
	if c == nil || c.ptr == nil {
		return LOCKDOWN_E_INVALID_ARG
	}
	var cSessionId *C.char
	if sessionId != "" {
		cSessionId = C.CString(sessionId)
		defer C.free(unsafe.Pointer(cSessionId))
	}

	err := C.lockdownd_stop_session(*c.ptr, cSessionId)

	return checkLockdowndErr(err)
}

// Send Sends a plist to lockdownd.
func (c *LockdowndClient) Send(plist PList) error {
	if c == nil || c.ptr == nil {
		return LOCKDOWN_E_INVALID_ARG
	}

	err := C.lockdownd_send(*c.ptr, plist.ptr)

	return checkLockdowndErr(err)
}

// Receive Receives a plist from lockdownd.
func (c *LockdowndClient) Receive() (*PList, error) {
	var cPlist C.plist_t

	err := C.lockdownd_receive(*c.ptr, &cPlist)
	if err := checkLockdowndErr(err); err != nil {
		return nil, err
	}

	return &PList{ptr: cPlist}, nil
}

// Pair Pairs the device using the supplied pair record.
func (c *LockdowndClient) Pair(record *LockdowndPairRecord) error {
	var cRecord C.lockdownd_pair_record_t
	if record != nil {
		cRecord = record.ptr
	}

	err := C.lockdownd_pair(*c.ptr, cRecord)

	return checkLockdowndErr(err)
}

// PairWithOptions Pairs the device using the supplied pair record and passing the given options.
func (c *LockdowndClient) PairWithOptions(record *LockdowndPairRecord, opts *PList) (*PList, error) {
	var cRecord C.lockdownd_pair_record_t
	if record != nil {
		cRecord = record.ptr
	}

	var cOpts C.plist_t
	if opts != nil {
		cOpts = opts.ptr
	}

	var cResponse C.plist_t

	err := C.lockdownd_pair_with_options(*c.ptr, cRecord, cOpts, &cResponse)
	if err := checkLockdowndErr(err); err != nil {
		return nil, err
	}

	return &PList{ptr: cResponse}, nil
}

// ValidatePair Validates if the device is paired with the given HostID. If successful the
// specified host will become trusted host of the device indicated by the
// lockdownd preference named TrustedHostAttached. Otherwise the host must be
// paired using lockdownd_pair() first.
func (c *LockdowndClient) ValidatePair(record *LockdowndPairRecord) error {
	var cRecord C.lockdownd_pair_record_t
	if record != nil {
		cRecord = record.ptr
	}

	err := C.lockdownd_validate_pair(*c.ptr, cRecord)

	return checkLockdowndErr(err)
}

// Unpair Unpairs the device with the given HostID and removes the pairing records
// from the device and host if the internal pairing record management is used.
func (c *LockdowndClient) Unpair(record *LockdowndPairRecord) error {
	var cRecord C.lockdownd_pair_record_t
	if record != nil {
		cRecord = record.ptr
	}

	err := C.lockdownd_unpair(*c.ptr, cRecord)

	return checkLockdowndErr(err)
}

// Activate Activates the device. Only works within an open session.
// The ActivationRecord plist dictionary must be obtained using the
// activation protocol requesting from Apple's https webservice.
func (c *LockdowndClient) Activate(record *PList) error {
	var cPlist C.plist_t
	if record != nil {
		cPlist = record.ptr
	}

	err := C.lockdownd_activate(*c.ptr, cPlist)

	return checkLockdowndErr(err)
}

// Deactivate Deactivates the device, returning it to the locked “Activate with iTunes”
// screen.
func (c *LockdowndClient) Deactivate() error {
	err := C.lockdownd_deactivate(*c.ptr)

	return checkLockdowndErr(err)
}

// EnterRecovery Tells the device to immediately enter recovery mode.
func (c *LockdowndClient) EnterRecovery() error {
	err := C.lockdownd_enter_recovery(*c.ptr)

	return checkLockdowndErr(err)
}

// GoodBye Sends the Goodbye request to lockdownd signaling the end of communication.
func (c *LockdowndClient) GoodBye() error {
	err := C.lockdownd_goodbye(*c.ptr)

	return checkLockdowndErr(err)
}

// LockdowndPairRecord pair record holding device, host and root certificates along the host_id
type LockdowndPairRecord struct {
	ptr C.lockdownd_pair_record_t
}

// LockdowndServiceDescriptor service descriptor
type LockdowndServiceDescriptor struct {
	port       int    // port number the service was started on
	sslEnabled int    // an indicator if the service requires SSL
	identifier string // identifier of the service
}
