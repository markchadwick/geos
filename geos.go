package geos

/*
#cgo LDFLAGS: -lgeos_c
#include <stdlib.h>
#include <geos_c.h>

extern GEOSContextHandle_t initializeGEOS();
*/
import "C"

var (
	h                *Handle
	DefaultWKTReader *WKTReader
	DefaultWKTWriter *WKTWriter
	DefaultWKBReader *WKBReader
	DefaultWKBWriter *WKBWriter
)

func init() {
	h = NewHandle()
	DefaultWKTReader = h.NewWKTReader()
	DefaultWKTWriter = h.NewWKTWriter()
	DefaultWKBReader = h.NewWKBReader()
	DefaultWKBWriter = h.NewWKBWriter()
}

type Handle struct {
	ctx C.GEOSContextHandle_t
}

func NewHandle() *Handle {
	return &Handle{C.initializeGEOS()}
}

func (h *Handle) Destroy() {
	C.finishGEOS_r(h.ctx)
}

func (h *Handle) NewWKBReader() *WKBReader {
	return &WKBReader{
		ctx: h.ctx,
		r:   C.GEOSWKBReader_create_r(h.ctx),
	}
}

func (h *Handle) NewWKBWriter() *WKBWriter {
	return &WKBWriter{
		ctx: h.ctx,
		w:   C.GEOSWKBWriter_create_r(h.ctx),
	}
}

func (h *Handle) NewWKTReader() *WKTReader {
	return &WKTReader{
		ctx: h.ctx,
		r:   C.GEOSWKTReader_create_r(h.ctx),
	}
}

func (h *Handle) NewWKTWriter() *WKTWriter {
	return &WKTWriter{
		ctx: h.ctx,
		w:   C.GEOSWKTWriter_create_r(h.ctx),
	}
}

func (h *Handle) NewCoordSequnce(size, dims uint) *CoordSequence {
	return &CoordSequence{
		ctx: h.ctx,
		cs:  C.GEOSCoordSeq_create_r(h.ctx, C.uint(size), C.uint(dims)),
	}
}

func NewCoordSequnce(size, dims uint) *CoordSequence {
	return h.NewCoordSequnce(size, dims)
}

func (h *Handle) NewLinearRing(cs *CoordSequence) (*Geometry, error) {
	geom := C.GEOSGeom_createLinearRing_r(h.ctx, cs.cs)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(h.ctx, geom), nil
}

func NewLinearRing(cs *CoordSequence) (*Geometry, error) {
	return h.NewLinearRing(cs)
}
