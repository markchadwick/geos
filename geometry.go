package geos

import (
	// "unsafe"
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

// ----------------------------------------------------------------------------
// Factory functions

func GeomFromWKT(wkt string) (*Geometry, error) {
	return DefaultWKTReader.Read(wkt)
}

func NewPoint(cs *CoordSequence) (*Geometry, error) {
	geom := C.GEOSGeom_createPoint(cs.cs)
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
}

func NewLinearRing(cs *CoordSequence) (*Geometry, error) {
	geom := C.GEOSGeom_createLinearRing(cs.cs)
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
}

func NewLineString(cs *CoordSequence) (*Geometry, error) {
	geom := C.GEOSGeom_createLineString(cs.cs)
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
}

// TODO
func GeomFromWKB(wkt string) (*Geometry, error) {
	return nil, nil
}

// TODO
func GeomFromHex(wkt string) (*Geometry, error) {
	return nil, nil
}

// ----------------------------------------------------------------------------
// Util methods

func (g *Geometry) WKT() string {
	return DefaultWKTWriter.Write(g)
}

// TODO: Not really the right way to do this
func (g *Geometry) Poly() *Geometry {
	geom := C.GEOSGeom_createPolygon(g.geom, nil, 0)
	return &Geometry{geom}
}

// ----------------------------------------------------------------------------
// Linearref methods

func (g *Geometry) Project(g1 *Geometry) float64 {
	d := C.GEOSProject(g.geom, g1.geom)
	return float64(d)
}

func (g *Geometry) Interpolate(d float64) *Geometry {
	geom := C.GEOSInterpolate(g.geom, C.double(d))
	return &Geometry{geom}
}

func (g *Geometry) ProjectNormalized(g1 *Geometry) float64 {
	d := C.GEOSProjectNormalized(g.geom, g1.geom)
	return float64(d)
}

func (g *Geometry) InterpolateNormalized(d float64) *Geometry {
	geom := C.GEOSInterpolateNormalized(g.geom, C.double(d))
	return &Geometry{geom}
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
	geom := C.GEOSBuffer(g.geom, C.double(width), C.int(quadsegs))
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
}

func (g *Geometry) BufferWithStyle(width float64, quadsegs int,
	endCapStyle CapStyle, joinStyle JoinStyle, mitreLimit float64) (*Geometry, error) {
	geom := C.GEOSBufferWithStyle(g.geom, C.double(width), C.int(quadsegs),
		C.int(endCapStyle), C.int(joinStyle), C.double(mitreLimit))
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
}

func (g *Geometry) SingleSidedBuffer(width int64, quadsegs int,
	joinStyle JoinStyle, mitreLimit float64, leftSide int) (*Geometry, error) {
	geom := C.GEOSSingleSidedBuffer(g.geom, C.double(width), C.int(quadsegs),
		C.int(joinStyle), C.double(mitreLimit), C.int(leftSide))
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
}

// ----------------------------------------------------------------------------
// Topology Operations

func (g *Geometry) Envelope() (*Geometry, error) {
	geom := C.GEOSEnvelope(g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
}

func (g *Geometry) Intersection(g1 *Geometry) (*Geometry, error) {
	geom := C.GEOSIntersection(g.geom, g1.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
}

func (g *Geometry) ConvexHull() (*Geometry, error) {
	geom := C.GEOSConvexHull(g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
}

func (g *Geometry) Difference(g1 *Geometry) (*Geometry, error) {
	geom := C.GEOSDifference(g.geom, g1.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
}

func (g *Geometry) SymDifference(g1 *Geometry) (*Geometry, error) {
	geom := C.GEOSSymDifference(g.geom, g1.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
}

func (g *Geometry) Boundary() (*Geometry, error) {
	geom := C.GEOSBoundary(g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
}

func (g *Geometry) Union(g1 *Geometry) (*Geometry, error) {
	geom := C.GEOSUnion(g.geom, g1.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
}

func (g *Geometry) UnionCascaded() (*Geometry, error) {
	geom := C.GEOSUnionCascaded(g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
}

func (g *Geometry) PointOnSurface() (*Geometry, error) {
	geom := C.GEOSPointOnSurface(g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
}

func (g *Geometry) GetCentroid() (*Geometry, error) {
	geom := C.GEOSGetCentroid(g.geom)
	if geom == nil {
		return nil, GEOSError
	}
	return &Geometry{geom}, nil
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
	return binaryPredicate(C.GEOSDisjoint(g.geom, g1.geom))
}

func (g *Geometry) Touches(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSTouches(g.geom, g1.geom))
}

func (g *Geometry) Intersects(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSIntersects(g.geom, g1.geom))
}

func (g *Geometry) Crosses(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSCrosses(g.geom, g1.geom))
}

func (g *Geometry) Within(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSWithin(g.geom, g1.geom))
}

func (g *Geometry) Contains(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSContains(g.geom, g1.geom))
}

func (g *Geometry) Overlaps(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSOverlaps(g.geom, g1.geom))
}

func (g *Geometry) Equals(g1 *Geometry) (bool, error) {
	return binaryPredicate(C.GEOSEquals(g.geom, g1.geom))
}

func (g *Geometry) EqualsExact(g1 *Geometry, tolerance float64) (bool, error) {
	return binaryPredicate(C.GEOSEqualsExact(g.geom, g1.geom,
		C.double(tolerance)))
}
