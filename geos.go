package geos

/*
#cgo LDFLAGS: -lgeos_c
#include <stdlib.h>
#include <geos_c.h>

extern GEOSContextHandle_t initializeGEOS();
*/
import "C"

var (
	ctx              C.GEOSContextHandle_t
	DefaultWKTReader *WKTReader
	DefaultWKTWriter *WKTWriter
	DefaultWKBReader *WKBReader
	DefaultWKBWriter *WKBWriter
)

func init() {
	ctx = C.initializeGEOS()
	DefaultWKTReader = NewWKTReader()
	DefaultWKTWriter = NewWKTWriter()
	DefaultWKBReader = NewWKBReader()
	DefaultWKBWriter = NewWKBWriter()
}
