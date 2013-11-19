package geos

// #include <geos_c.h>
import "C"

type CoordSequence struct {
	ctx C.GEOSContextHandle_t
	cs  *C.GEOSCoordSequence
}

// TODO: This methods return 0 on error
func (cs *CoordSequence) SetX(idx uint, v float64) {
	C.GEOSCoordSeq_setX_r(cs.ctx, cs.cs, C.uint(idx), C.double(v))
}

func (cs *CoordSequence) SetY(idx uint, v float64) {
	C.GEOSCoordSeq_setY_r(cs.ctx, cs.cs, C.uint(idx), C.double(v))
}

func (cs *CoordSequence) SetZ(idx uint, v float64) {
	C.GEOSCoordSeq_setZ_r(cs.ctx, cs.cs, C.uint(idx), C.double(v))
}

func (cs *CoordSequence) Point() (*Geometry, error) {
	geom := C.GEOSGeom_createPoint_r(cs.ctx, cs.cs)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(cs.ctx, geom), nil
}

func (cs *CoordSequence) LinearRing() (*Geometry, error) {
	geom := C.GEOSGeom_createLinearRing_r(cs.ctx, cs.cs)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(cs.ctx, geom), nil
}
