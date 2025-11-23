package libimobiledevice

/*
#include <plist/plist.h>
*/
import "C"

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

func (l *PList) Free() {
	if l != nil && l.ptr != nil {
		C.plist_free(l.ptr)
	}
}
