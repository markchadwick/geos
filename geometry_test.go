package geos

import (
	"github.com/markchadwick/spec"
)

var _ = spec.Suite("Geometry", func(c *spec.C) {
	h := NewHandle()
	defer h.Destroy()

	c.It("should create from WKT", func(c *spec.C) {
		point, err := DefaultWKTReader.Read("POINT (8 9)")
		c.Assert(err).IsNil()
		c.Assert(point).NotNil()
	})

	c.It("should have cap styles", func(c *spec.C) {
		c.Assert(int(CapRound)).Equals(1)
		c.Assert(int(CapFlat)).Equals(2)
		c.Assert(int(CapSquare)).Equals(3)
	})

	c.It("should have join styles", func(c *spec.C) {
		c.Assert(int(JoinRound)).Equals(1)
		c.Assert(int(JoinMitre)).Equals(2)
		c.Assert(int(JoinBevel)).Equals(3)
	})

	c.It("should calculate the area", func(c *spec.C) {
		square, err := DefaultWKTReader.Read("POLYGON ((0 0, 1 0, 1 1, 0 1, 0 0))")
		c.Assert(err).IsNil()
		c.Assert(square).NotNil()
		c.Assert(square.Area()).Equals(1.0000)
	})
})
