package geos

import (
	"github.com/markchadwick/spec"
)

var _ = spec.Suite("WKT Reader", func(c *spec.C) {
	r := NewWKTReader()
	defer r.Destroy()

	c.It("should read a simple WKT to a geometry", func(c *spec.C) {
		p, err := r.Read("POINT (100 200)")
		c.Assert(err).IsNil()
		c.Assert(p).NotNil()
		c.Assert(p.geom).NotNil()
	})
})

var _ = spec.Suite("WKT Writer", func(c *spec.C) {
	r := NewWKTReader()
	defer r.Destroy()

	w := NewWKTWriter()
	defer w.Destroy()

	c.It("should read a geometry to WKT", func(c *spec.C) {
		point, _ := r.Read("POINT (100 200)")

		wkt := w.Write(point)
		c.Assert(wkt).Equals("POINT (100.0000000000000000 200.0000000000000000)")
	})
})
