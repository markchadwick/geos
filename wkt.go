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
	ctx C.GEOSContextHandle_t
	r   *C.GEOSWKTReader
}

func (r *WKTReader) Read(wkt string) (*Geometry, error) {
	str := C.CString(wkt)
	defer C.free(unsafe.Pointer(str))
	geom := C.GEOSWKTReader_read_r(r.ctx, r.r, str)
	if geom == nil {
		return nil, fmt.Errorf("Malformed WKT: %s", wkt)
	}

	return geometry(r.ctx, geom), nil
}

func (r *WKTReader) Destroy() {
	C.GEOSWKTReader_destroy_r(r.ctx, r.r)
}

// ----------------------------------------------------------------------------
// WKT Writer
type WKTWriter struct {
	ctx C.GEOSContextHandle_t
	w   *C.GEOSWKTWriter
}

func (w *WKTWriter) Write(geom *Geometry) string {
	str := C.GEOSWKTWriter_write_r(w.ctx, w.w, geom.geom)
	defer C.free(unsafe.Pointer(str))
	return C.GoString(str)
}

func (w *WKTWriter) Destroy() {
	C.GEOSWKTWriter_destroy_r(w.ctx, w.w)
}
