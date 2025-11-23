package libimobiledevice

/*
#include <plist/plist.h>
*/
import "C"

type PList struct {
	ptr C.plist_t
}

func (l *PList) Free() {
	if l != nil && l.ptr != nil {
		C.plist_free(l.ptr)
	}
}
