package geos

import (
	"github.com/markchadwick/spec"
)

var _ = spec.Suite("Geometry", func(c *spec.C) {
	c.It("should create from WKT", func(c *spec.C) {
		point, err := GeomFromWKT("POINT (8 9)")
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
})
