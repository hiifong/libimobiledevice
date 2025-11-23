package libimobiledevice

/*
#include <plist/plist.h>
*/
import "C"
import "unsafe"

// PListType The enumeration of plist node types.
type PListType int

const (
	// PLIST_NONE No type
	PLIST_NONE PListType = C.PLIST_NONE

	// PLIST_BOOLEAN Boolean, scalar type
	PLIST_BOOLEAN PListType = C.PLIST_BOOLEAN

	// PLIST_INT Integer, scalar type
	PLIST_INT PListType = C.PLIST_INT

	// PLIST_REAL Real, scalar type
	PLIST_REAL PListType = C.PLIST_REAL

	// PLIST_STRING ASCII string, scalar type
	PLIST_STRING PListType = C.PLIST_STRING

	// PLIST_ARRAY Ordered array, structured type
	PLIST_ARRAY PListType = C.PLIST_ARRAY

	// PLIST_DICT Unordered dictionary (key/value pair), structured type
	PLIST_DICT PListType = C.PLIST_DICT

	// PLIST_DATE Date, scalar type
	PLIST_DATE PListType = C.PLIST_DATE

	// PLIST_DATA Binary data, scalar type
	PLIST_DATA PListType = C.PLIST_DATA

	// PLIST_KEY Key in dictionaries (ASCII String), scalar type
	PLIST_KEY PListType = C.PLIST_KEY

	// PLIST_UID Special type used for 'keyed encoding'
	PLIST_UID PListType = C.PLIST_UID

	// PLIST_NULL NULL type
	PLIST_NULL PListType = C.PLIST_NULL
)

// PListErr libplist error values
type PListErr int

const (
	// PLIST_ERR_SUCCESS operation successful
	PLIST_ERR_SUCCESS PListErr = C.PLIST_ERR_SUCCESS

	// PLIST_ERR_INVALID_ARG one or more of the parameters are invalid
	PLIST_ERR_INVALID_ARG PListErr = C.PLIST_ERR_INVALID_ARG

	// PLIST_ERR_FORMAT the plist contains nodes not compatible with the output format
	PLIST_ERR_FORMAT PListErr = C.PLIST_ERR_FORMAT

	// PLIST_ERR_PARSE parsing of the input format failed
	PLIST_ERR_PARSE PListErr = C.PLIST_ERR_PARSE

	// PLIST_ERR_NO_MEM not enough memory to handle the operation
	PLIST_ERR_NO_MEM PListErr = C.PLIST_ERR_NO_MEM

	// PLIST_ERR_IO I/O error
	PLIST_ERR_IO PListErr = C.PLIST_ERR_IO

	// PLIST_ERR_UNKNOWN an unspecified error occurred
	PLIST_ERR_UNKNOWN PListErr = C.PLIST_ERR_UNKNOWN
)

// PListFormatType libplist format types
type PListFormatType int

const (
	// PLIST_FORMAT_NONE No format
	PLIST_FORMAT_NONE PListFormatType = C.PLIST_FORMAT_NONE

	// PLIST_FORMAT_XML XML format
	PLIST_FORMAT_XML PListFormatType = C.PLIST_FORMAT_XML

	// PLIST_FORMAT_BINARY bplist00 format
	PLIST_FORMAT_BINARY PListFormatType = C.PLIST_FORMAT_BINARY

	// PLIST_FORMAT_JSON JSON format
	PLIST_FORMAT_JSON PListFormatType = C.PLIST_FORMAT_JSON

	// PLIST_FORMAT_OSTEP OpenStep "old-style" plist format
	PLIST_FORMAT_OSTEP PListFormatType = C.PLIST_FORMAT_OSTEP

	// 5-9 are reserved for possible future use

	// PLIST_FORMAT_PRINT human-readable output-only format
	PLIST_FORMAT_PRINT PListFormatType = C.PLIST_FORMAT_PRINT

	// PLIST_FORMAT_LIMD "libimobiledevice" output-only format (ideviceinfo)
	PLIST_FORMAT_LIMD PListFormatType = C.PLIST_FORMAT_LIMD

	// PLIST_FORMAT_PLUTIL plutil-style output-only format
	PLIST_FORMAT_PLUTIL PListFormatType = C.PLIST_FORMAT_PLUTIL
)

// PListWriteOptionType libplist write options
type PListWriteOptionType int

const (
	// PLIST_OPT_NONE Default value to use when none of the options is needed.
	PLIST_OPT_NONE PListWriteOptionType = C.PLIST_OPT_NONE

	// PLIST_OPT_COMPACT Use a compact representation (non-prettified). Only valid for #PLIST_FORMAT_JSON and #PLIST_FORMAT_OSTEP.
	PLIST_OPT_COMPACT PListWriteOptionType = C.PLIST_OPT_COMPACT

	// PLIST_OPT_PARTIAL_DATA Print 24 bytes maximum of #PLIST_DATA values. If the data is longer than 24 bytes,  the first 16 and last 8 bytes will be written. Only valid for #PLIST_FORMAT_PRINT.
	PLIST_OPT_PARTIAL_DATA PListWriteOptionType = C.PLIST_OPT_PARTIAL_DATA

	// PLIST_OPT_NO_NEWLINE Do not print a final newline character. Only valid for #PLIST_FORMAT_PRINT, #PLIST_FORMAT_LIMD, and #PLIST_FORMAT_PLUTIL.
	PLIST_OPT_NO_NEWLINE PListWriteOptionType = C.PLIST_OPT_NO_NEWLINE

	// PLIST_OPT_INDENT Indent each line of output. Currently only #PLIST_FORMAT_PRINT and #PLIST_FORMAT_LIMD are supported. Use #PLIST_OPT_INDENT_BY() macro to specify the level of indentation.
	PLIST_OPT_INDENT PListWriteOptionType = C.PLIST_OPT_INDENT
)

type PList struct {
	ptr C.plist_t
}

// Creation & Destruction

// NewPListDict Create a new root plist_t type #PLIST_DICT
func NewPListDict() *PList {
	return &PList{C.plist_new_dict()}
}

// NewPListArray Create a new root plist_t type #PLIST_ARRAY
func NewPListArray() *PList {
	return &PList{C.plist_new_array()}
}

// NewPListString Create a new plist_t type #PLIST_STRING
func NewPListString(val string) *PList {
	cVal := C.CString(val)
	defer C.free(unsafe.Pointer(cVal))

	return &PList{C.plist_new_string(cVal)}
}

// NewPListBool Create a new plist_t type #PLIST_BOOLEAN
func NewPListBool(val bool) *PList {
	var cVal C.uint8_t
	if val {
		cVal = C.uint8_t(1)
	} else {
		cVal = C.uint8_t(0)
	}
	return &PList{C.plist_new_bool(cVal)}
}

// NewPListUint Create a new plist_t type #PLIST_INT with an unsigned integer value
func NewPListUint(val uint) *PList {
	return &PList{C.plist_new_uint(C.uint64_t(val))}
}

// NewPListInt Create a new plist_t type #PLIST_INT with a signed integer value
func NewPListInt(val int) *PList {
	return &PList{C.plist_new_int(C.int64_t(val))}
}

// NewPListReal Create a new plist_t type #PLIST_REAL
func NewPListReal(val float64) *PList {
	return &PList{C.plist_new_real(C.double(val))}
}

// NewPListData Create a new plist_t type #PLIST_DATA
func NewPListData(val string) *PList {
	cVal := C.CString(val)
	cLen := C.uint64_t(len(val))
	defer C.free(unsafe.Pointer(cVal))
	return &PList{C.plist_new_data(cVal, cLen)}
}

// NewPListUnixDate Create a new plist_t type #PLIST_DATE
func NewPListUnixDate(sec int) *PList {
	return &PList{C.plist_new_unix_date(C.int64_t(sec))}
}

// NewPListUid Create a new plist_t type #PLIST_UID
func NewPListUid(val uint) *PList {
	return &PList{C.plist_new_uid(C.uint64_t(val))}
}

// NewPListNull Create a new plist_t type #PLIST_NULL
func NewPListNull() *PList {
	return &PList{C.plist_new_null()}
}

// Free plist
func (l *PList) Free() {
	if l != nil && l.ptr != nil {
		C.plist_free(l.ptr)
	}
}

// Array functions

// GetSize Get size of a #PLIST_ARRAY node.
func (l *PList) GetSize() uint {
	if l == nil || l.ptr == nil {
		return 0
	}
	return uint(C.plist_array_get_size(l.ptr))
}

// GetArrayItem Get the nth item in a #PLIST_ARRAY node.
func (l *PList) GetArrayItem(n uint) *PList {
	if l == nil || l.ptr == nil {
		return nil
	}

	var cPlist C.plist_t
	cPlist = C.plist_array_get_item(l.ptr, C.uint32_t(n))
	return &PList{cPlist}
}

// GetArrayItemIndex Get the index of an item. item must be a member of a #PLIST_ARRAY node.
func (l *PList) GetArrayItemIndex(item *PList) uint {
	if l == nil || l.ptr == nil || item == nil || item.ptr == nil {
		return 0
	}

	return uint(C.plist_array_get_item_index(l.ptr, item.ptr))
}

// SetArrayItem Set the nth item in a #PLIST_ARRAY node.
// The previous item at index n will be freed using #plist_free
func (l *PList) SetArrayItem(item *PList, n uint) {
	if l == nil || l.ptr == nil || item == nil || item.ptr == nil {
		return
	}

	C.plist_array_set_item(l.ptr, item.ptr, C.uint32_t(n))
}

// AppendArrayItem Append a new item at the end of a #PLIST_ARRAY node.
func (l *PList) AppendArrayItem(item *PList) {
	if l == nil || l.ptr == nil || item == nil || item.ptr == nil {
		return
	}

	C.plist_array_append_item(l.ptr, item.ptr)
}

// InsertArrayItem Insert a new item at position n in a #PLIST_ARRAY node.
func (l *PList) InsertArrayItem(item *PList, n uint) {
	if l == nil || l.ptr == nil || item == nil || item.ptr == nil {
		return
	}

	C.plist_array_insert_item(l.ptr, item.ptr, C.uint32_t(n))
}

// RemoveArrayItem Remove an existing position in a #PLIST_ARRAY node.
// Removed position will be freed using #plist_free.
func (l *PList) RemoveArrayItem(n uint) {
	if l == nil || l.ptr == nil {
		return
	}

	C.plist_array_remove_item(l.ptr, C.uint32_t(n))
}

// ArrayItemRemove Remove a node that is a child node of a #PLIST_ARRAY node.
// node will be freed using #plist_free.
func (l *PList) ArrayItemRemove() {
	if l == nil || l.ptr == nil {
		return
	}

	C.plist_array_item_remove(l.ptr)
}

// NewArrayIter Create an iterator of a #PLIST_ARRAY node.
// The allocated iterator should be freed with the standard free function.
func (l *PList) NewArrayIter() C.plist_array_iter {
	var cIter C.plist_array_iter
	C.plist_array_new_iter(l.ptr, &cIter)
	return &cIter
}

// ArrayNextItem Increment iterator of a #PLIST_ARRAY node.
func (l *PList) ArrayNextItem() {
	// TODO
	return
}
