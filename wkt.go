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
		r: C.GEOSWKTReader_create_r(ctx),
	}
}

func (r *WKTReader) Read(wkt string) (*Geometry, error) {
	str := C.CString(wkt)
	defer C.free(unsafe.Pointer(str))
	geom := C.GEOSWKTReader_read_r(ctx, r.r, str)
	if geom == nil {
		return nil, fmt.Errorf("Malformed WKT: %s", wkt)
	}

	return geometry(geom), nil
}

func (r *WKTReader) Destroy() {
	C.GEOSWKTReader_destroy_r(ctx, r.r)
}

// ----------------------------------------------------------------------------
// WKT Writer
type WKTWriter struct {
	w *C.GEOSWKTWriter
}

func NewWKTWriter() *WKTWriter {
	return &WKTWriter{
		w: C.GEOSWKTWriter_create_r(ctx),
	}
}

func (w *WKTWriter) Write(geom *Geometry) string {
	str := C.GEOSWKTWriter_write_r(ctx, w.w, geom.geom)
	defer C.free(unsafe.Pointer(str))
	return C.GoString(str)
}

func (w *WKTWriter) Destroy() {
	C.GEOSWKTWriter_destroy_r(ctx, w.w)
}
