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
	geom *C.GEOSGeometry
}

func geometry(geom *C.GEOSGeometry) *Geometry {
	g := &Geometry{
		geom: geom,
	}
	return g
}

// ----------------------------------------------------------------------------
// Factory functions

func GeomFromWKT(wkt string) (*Geometry, error) {
	return DefaultWKTReader.Read(wkt)
}

func GeomFromWKB(wkb []byte) (*Geometry, error) {
	return DefaultWKBReader.Read(wkb)
}

func NewPoint(cs *CoordSequence) (*Geometry, error) {
	geom := C.GEOSGeom_createPoint_r(ctx, cs.cs)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
}

func NewLinearRing(cs *CoordSequence) (*Geometry, error) {
	geom := C.GEOSGeom_createLinearRing_r(ctx, cs.cs)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
}

func NewLineString(cs *CoordSequence) (*Geometry, error) {
	geom := C.GEOSGeom_createLineString_r(ctx, cs.cs)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
}

// ----------------------------------------------------------------------------
// Util methods

func (g *Geometry) WKT() string {
	return DefaultWKTWriter.Write(g)
}

func (g *Geometry) WKB() []byte {
	return DefaultWKBWriter.Write(g)
}

func (g *Geometry) Destroy() {
	C.GEOSGeom_destroy_r(ctx, g.geom)
}

// TODO: Not really the right way to do this
func (g *Geometry) Poly() *Geometry {
	geom := C.GEOSGeom_createPolygon_r(ctx, g.geom, nil, 0)
	// Not GC'd -- handle by g
	return &Geometry{geom}
}

// ----------------------------------------------------------------------------
// Linearref methods

func (g *Geometry) Project(g1 *Geometry) float64 {
	d := C.GEOSProject_r(ctx, g.geom, g1.geom)
	return float64(d)
}

func (g *Geometry) Interpolate(d float64) *Geometry {
	geom := C.GEOSInterpolate_r(ctx, g.geom, C.double(d))
	return geometry(geom)
}

func (g *Geometry) ProjectNormalized(g1 *Geometry) float64 {
	d := C.GEOSProjectNormalized_r(ctx, g.geom, g1.geom)
	return float64(d)
}

func (g *Geometry) InterpolateNormalized(d float64) *Geometry {
	geom := C.GEOSInterpolateNormalized_r(ctx, g.geom, C.double(d))
	return geometry(geom)
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
	geom := C.GEOSBuffer_r(ctx, g.geom, C.double(width), C.int(quadsegs))
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
}

func (g *Geometry) BufferWithStyle(width float64, quadsegs int,
	endCapStyle CapStyle, joinStyle JoinStyle, mitreLimit float64) (*Geometry, error) {
	geom := C.GEOSBufferWithStyle_r(ctx, g.geom, C.double(width), C.int(quadsegs),
		C.int(endCapStyle), C.int(joinStyle), C.double(mitreLimit))
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
}

func (g *Geometry) SingleSidedBuffer(width int64, quadsegs int,
	joinStyle JoinStyle, mitreLimit float64, leftSide int) (*Geometry, error) {
	geom := C.GEOSSingleSidedBuffer_r(ctx, g.geom, C.double(width), C.int(quadsegs),
		C.int(joinStyle), C.double(mitreLimit), C.int(leftSide))
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
}

// ----------------------------------------------------------------------------
// Topology Operations

func (g *Geometry) Envelope() (*Geometry, error) {
	geom := C.GEOSEnvelope_r(ctx, g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
}

func (g *Geometry) Intersection(g1 *Geometry) (*Geometry, error) {
	geom := C.GEOSIntersection_r(ctx, g.geom, g1.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
}

func (g *Geometry) ConvexHull() (*Geometry, error) {
	geom := C.GEOSConvexHull_r(ctx, g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
}

func (g *Geometry) Difference(g1 *Geometry) (*Geometry, error) {
	geom := C.GEOSDifference_r(ctx, g.geom, g1.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
}

func (g *Geometry) SymDifference(g1 *Geometry) (*Geometry, error) {
	geom := C.GEOSSymDifference_r(ctx, g.geom, g1.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
}

func (g *Geometry) Boundary() (*Geometry, error) {
	geom := C.GEOSBoundary_r(ctx, g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
}

func (g *Geometry) Union(g1 *Geometry) (*Geometry, error) {
	geom := C.GEOSUnion_r(ctx, g.geom, g1.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
}

func (g *Geometry) UnionCascaded() (*Geometry, error) {
	geom := C.GEOSUnionCascaded_r(ctx, g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
}

func (g *Geometry) PointOnSurface() (*Geometry, error) {
	geom := C.GEOSPointOnSurface_r(ctx, g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
}

func (g *Geometry) GetCentroid() (*Geometry, error) {
	geom := C.GEOSGetCentroid_r(ctx, g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return geometry(geom), nil
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
	return binaryPredicate(C.GEOSDisjoint_r(ctx, g.geom, g1.geom))
}

func (g *Geometry) Touches(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSTouches_r(ctx, g.geom, g1.geom))
}

func (g *Geometry) Intersects(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSIntersects_r(ctx, g.geom, g1.geom))
}

func (g *Geometry) Crosses(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSCrosses_r(ctx, g.geom, g1.geom))
}

func (g *Geometry) Within(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSWithin_r(ctx, g.geom, g1.geom))
}

func (g *Geometry) Contains(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSContains_r(ctx, g.geom, g1.geom))
}

func (g *Geometry) Overlaps(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSOverlaps_r(ctx, g.geom, g1.geom))
}

func (g *Geometry) Equals(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSEquals_r(ctx, g.geom, g1.geom))
}

func (g *Geometry) EqualsExact(g1 *Geometry, tolerance float64) (bool, error) {
	return binaryPredicate(C.GEOSEqualsExact_r(ctx, g.geom, g1.geom,
		C.double(tolerance)))
}
