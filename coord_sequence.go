package geos

// #include <geos_c.h>
import "C"

type CoordSequence struct {
	cs *C.GEOSCoordSequence
}

func NewCoordSequnce(size, dims uint) *CoordSequence {
	return &CoordSequence{
		cs: C.GEOSCoordSeq_create(C.uint(size), C.uint(dims)),
	}
}

// TODO: This methods return 0 on error
func (cs *CoordSequence) SetX(idx uint, v float64) {
	C.GEOSCoordSeq_setX(cs.cs, C.uint(idx), C.double(v))
}

func (cs *CoordSequence) SetY(idx uint, v float64) {
	C.GEOSCoordSeq_setY(cs.cs, C.uint(idx), C.double(v))
}

func (cs *CoordSequence) SetZ(idx uint, v float64) {
	C.GEOSCoordSeq_setZ(cs.cs, C.uint(idx), C.double(v))
}
