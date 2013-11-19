package geos

import (
	"errors"
)

// #include <stdlib.h>
// #include <geos_c.h>
import "C"

var (
	GEOSError = errors.New("GEOS Error unknown")
)

type Geometry struct {
	ctx  C.GEOSContextHandle_t
	geom *C.GEOSGeometry
}

func geometry(ctx C.GEOSContextHandle_t, geom *C.GEOSGeometry) *Geometry {
	g := &Geometry{
		ctx:  ctx,
		geom: geom,
	}
	return g
}

func (g *Geometry) Destroy() {
	C.GEOSGeom_destroy_r(g.ctx, g.geom)
}

// TODO: Not really the right way to do this
func (g *Geometry) Poly() *Geometry {
	geom := C.GEOSGeom_createPolygon_r(g.ctx, g.geom, nil, 0)
	// Not GC'd -- handle by g
	return &Geometry{g.ctx, geom}
}

// ----------------------------------------------------------------------------
// Linearref methods

func (g *Geometry) Project(g1 *Geometry) float64 {
	d := C.GEOSProject_r(g.ctx, g.geom, g1.geom)
	return float64(d)
}

func (g *Geometry) Interpolate(d float64) *Geometry {
	geom := C.GEOSInterpolate_r(g.ctx, g.geom, C.double(d))
	return geometry(g.ctx, geom)
}

func (g *Geometry) ProjectNormalized(g1 *Geometry) float64 {
	d := C.GEOSProjectNormalized_r(g.ctx, g.geom, g1.geom)
	return float64(d)
}

func (g *Geometry) InterpolateNormalized(d float64) *Geometry {
	geom := C.GEOSInterpolateNormalized_r(g.ctx, g.geom, C.double(d))
	return geometry(g.ctx, geom)
}

// ----------------------------------------------------------------------------
// Buffer methods

type CapStyle int
type JoinStyle int

const (
	CapRound CapStyle = iota + 1
	CapFlat
	CapSquare
)

const (
	JoinRound JoinStyle = iota + 1
	JoinMitre
	JoinBevel
)

func (g *Geometry) Buffer(width float64, quadsegs int) (*Geometry, error) {
	geom := C.GEOSBuffer_r(g.ctx, g.geom, C.double(width), C.int(quadsegs))
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(g.ctx, geom), nil
}

func (g *Geometry) BufferWithStyle(width float64, quadsegs int,
	endCapStyle CapStyle, joinStyle JoinStyle, mitreLimit float64) (*Geometry, error) {
	geom := C.GEOSBufferWithStyle_r(g.ctx, g.geom, C.double(width), C.int(quadsegs),
		C.int(endCapStyle), C.int(joinStyle), C.double(mitreLimit))
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(g.ctx, geom), nil
}

func (g *Geometry) SingleSidedBuffer(width int64, quadsegs int,
	joinStyle JoinStyle, mitreLimit float64, leftSide int) (*Geometry, error) {
	geom := C.GEOSSingleSidedBuffer_r(g.ctx, g.geom, C.double(width), C.int(quadsegs),
		C.int(joinStyle), C.double(mitreLimit), C.int(leftSide))
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(g.ctx, geom), nil
}

// ----------------------------------------------------------------------------
// Topology Operations

func (g *Geometry) Envelope() (*Geometry, error) {
	geom := C.GEOSEnvelope_r(g.ctx, g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(g.ctx, geom), nil
}

func (g *Geometry) Intersection(g1 *Geometry) (*Geometry, error) {
	geom := C.GEOSIntersection_r(g.ctx, g.geom, g1.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(g.ctx, geom), nil
}

func (g *Geometry) ConvexHull() (*Geometry, error) {
	geom := C.GEOSConvexHull_r(g.ctx, g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(g.ctx, geom), nil
}

func (g *Geometry) Difference(g1 *Geometry) (*Geometry, error) {
	geom := C.GEOSDifference_r(g.ctx, g.geom, g1.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(g.ctx, geom), nil
}

func (g *Geometry) SymDifference(g1 *Geometry) (*Geometry, error) {
	geom := C.GEOSSymDifference_r(g.ctx, g.geom, g1.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(g.ctx, geom), nil
}

func (g *Geometry) Boundary() (*Geometry, error) {
	geom := C.GEOSBoundary_r(g.ctx, g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(g.ctx, geom), nil
}

func (g *Geometry) Union(g1 *Geometry) (*Geometry, error) {
	geom := C.GEOSUnion_r(g.ctx, g.geom, g1.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(g.ctx, geom), nil
}

func (g *Geometry) UnionCascaded() (*Geometry, error) {
	geom := C.GEOSUnionCascaded_r(g.ctx, g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(g.ctx, geom), nil
}

func (g *Geometry) PointOnSurface() (*Geometry, error) {
	geom := C.GEOSPointOnSurface_r(g.ctx, g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(g.ctx, geom), nil
}

func (g *Geometry) GetCentroid() (*Geometry, error) {
	geom := C.GEOSGetCentroid_r(g.ctx, g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(g.ctx, geom), nil
}

// ----------------------------------------------------------------------------
// Binary predicates

const (
	binaryPredExc   = C.char(2)
	binaryPredTrue  = C.char(1)
	binaryPredFalse = C.char(0)
)

func binaryPredicate(c C.char) (bool, error) {
	switch c {
	default:
		return false, GEOSError
	case binaryPredTrue:
		return true, nil
	case binaryPredFalse:
		return false, nil
	}
}

// TODO Relate Pattern

func (g *Geometry) Disjoint(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSDisjoint_r(g.ctx, g.geom, g1.geom))
}

func (g *Geometry) Touches(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSTouches_r(g.ctx, g.geom, g1.geom))
}

func (g *Geometry) Intersects(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSIntersects_r(g.ctx, g.geom, g1.geom))
}

func (g *Geometry) Crosses(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSCrosses_r(g.ctx, g.geom, g1.geom))
}

func (g *Geometry) Within(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSWithin_r(g.ctx, g.geom, g1.geom))
}

func (g *Geometry) Contains(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSContains_r(g.ctx, g.geom, g1.geom))
}

func (g *Geometry) Overlaps(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSOverlaps_r(g.ctx, g.geom, g1.geom))
}

func (g *Geometry) Equals(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSEquals_r(g.ctx, g.geom, g1.geom))
}

func (g *Geometry) EqualsExact(g1 *Geometry, tolerance float64) (bool, error) {
	return binaryPredicate(C.GEOSEqualsExact_r(g.ctx, g.geom, g1.geom,
		C.double(tolerance)))
}
