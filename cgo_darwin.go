//go:build darwin

package libimobiledevice

/*
#cgo darwin CFLAGS: -I/opt/homebrew/include
#cgo darwin LDFLAGS: -L/opt/homebrew/lib
#cgo darwin LDFLAGS: -limobiledevice-1.0

#include <stdlib.h>
#include <libimobiledevice/libimobiledevice.h>
*/
import "C"
