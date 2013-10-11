package geos

/*
#cgo LDFLAGS: -lgeos_c
#include <stdlib.h>
#include <geos_c.h>

extern void initializeGEOS();
*/
import "C"

var (
	DefaultWKTReader *WKTReader
	DefaultWKTWriter *WKTWriter
	DefaultWKBReader *WKBReader
	DefaultWKBWriter *WKBWriter
)

func init() {
	C.initializeGEOS()
	DefaultWKTReader = NewWKTReader()
	DefaultWKTWriter = NewWKTWriter()
	DefaultWKBReader = NewWKBReader()
	DefaultWKBWriter = NewWKBWriter()
}
