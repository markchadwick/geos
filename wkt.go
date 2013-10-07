package geos

import (
	"fmt"
	"unsafe"
)

// #include <stdlib.h>
// #include <geos_c.h>
import "C"

// ----------------------------------------------------------------------------
// WKT Reader
type WKTReader struct {
	r *C.GEOSWKTReader
}

func NewWKTReader() *WKTReader {
	return &WKTReader{
		r: C.GEOSWKTReader_create(),
	}
}

func (r *WKTReader) Read(wkt string) (*Geometry, error) {
	str := C.CString(wkt)
	defer C.free(unsafe.Pointer(str))
	geom := C.GEOSWKTReader_read(r.r, str)
	if geom == nil {
		return nil, fmt.Errorf("Malformed WKT: %s", wkt)
	}

	return &Geometry{geom}, nil
}

func (r *WKTReader) Destroy() {
	C.GEOSWKTReader_destroy(r.r)
}

// ----------------------------------------------------------------------------
// WKT Writer
type WKTWriter struct {
	w *C.GEOSWKTWriter
}

func NewWKTWriter() *WKTWriter {
	return &WKTWriter{
		w: C.GEOSWKTWriter_create(),
	}
}

func (w *WKTWriter) Write(geom *Geometry) string {
	str := C.GEOSWKTWriter_write(w.w, geom.geom)
	defer C.free(unsafe.Pointer(str))
	return C.GoString(str)
}

func (w *WKTWriter) Destroy() {
	C.GEOSWKTWriter_destroy(w.w)
}
