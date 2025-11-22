//go:build linux

package libimobiledevice

/*
#cgo linux CFLAGS: -I/usr/include
#cgo linux LDFLAGS: -L/usr/lib
#cgo linux LDFLAGS: -limobiledevice-1.0

#include <stdlib.h>
#include <libimobiledevice/libimobiledevice.h>
*/
import "C"
