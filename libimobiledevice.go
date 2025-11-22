package libimobiledevice

/*
#include <stdlib.h>
#include <libimobiledevice/libimobiledevice.h>
*/
import "C"

// SetDebugLevel Set the level of debugging.  Set to false for no debug output or true to enable debug output.
func SetDebugLevel(enable bool) {
	if enable {
		C.idevice_set_debug_level(C.int(1))
	} else {
		C.idevice_set_debug_level(C.int(0))
	}
}
