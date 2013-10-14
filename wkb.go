package geos

import (
	"errors"
	"fmt"
	"unsafe"
)

// #include <stdlib.h>
// #include <geos_c.h>
import "C"

// ----------------------------------------------------------------------------
// WKB Reader
type WKBReader struct {
	r *C.GEOSWKBReader
}

func NewWKBReader() *WKBReader {
	return &WKBReader{
		r: C.GEOSWKBReader_create(),
	}
}

func (r *WKBReader) Read(wkb []byte) (*Geometry, error) {
	if len(wkb) < 1 {
		return nil, errors.New("Tried to read empty WKB")
	}

	d := (*C.uchar)(unsafe.Pointer(&wkb[0]))
	length := C.size_t(len(wkb))

	geom := C.GEOSWKBReader_read(r.r, d, length)
	if geom == nil {
		return nil, fmt.Errorf("Malformed WKB: %s", wkb)
	}

	return geometry(geom), nil
}

func (r *WKBReader) ReadHex(wkb []byte) (*Geometry, error) {
	d := (*C.uchar)(unsafe.Pointer(&wkb[0]))
	length := C.size_t(len(wkb))

	geom := C.GEOSWKBReader_readHEX(r.r, d, length)
	if geom == nil {
		return nil, fmt.Errorf("Malformed WKB Hex: %s", wkb)
	}

	return geometry(geom), nil
}

func (r *WKBReader) Destroy() {
	C.GEOSWKBReader_destroy(r.r)
}

// ----------------------------------------------------------------------------
// WKB Writer
type WKBWriter struct {
	w *C.GEOSWKBWriter
}

func NewWKBWriter() *WKBWriter {
	return &WKBWriter{
		w: C.GEOSWKBWriter_create(),
	}
}

func (w *WKBWriter) Write(geom *Geometry) []byte {
	size := C.size_t(1)
	cs := C.GEOSWKBWriter_write(w.w, geom.geom, &size)
	return C.GoBytes(unsafe.Pointer(cs), C.int(size))
}

func (w *WKBWriter) WriteHex(geom *Geometry) []byte {
	size := C.size_t(1)
	cs := C.GEOSWKBWriter_writeHEX(w.w, geom.geom, &size)
	return C.GoBytes(unsafe.Pointer(cs), C.int(size))
}

func (w *WKBWriter) Destroy() {
	C.GEOSWKBWriter_destroy(w.w)
}
