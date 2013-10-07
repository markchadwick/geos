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
)

func init() {
	C.initializeGEOS()
	DefaultWKTReader = NewWKTReader()
	DefaultWKTWriter = NewWKTWriter()
}
