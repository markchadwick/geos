package geos

import (
	"encoding/hex"
	"github.com/markchadwick/spec"
)

// hex(AsBinary(geometry))                     AsText(geometry)
// ------------------------------------------  -----------------------
// 010100000001DE02098A1B52C090A0F831E62E4640  POINT(-72.4303 44.3664)
// 0101000000569FABADD82B52C067D5E76A2B264640  POINT(-72.6851 44.2982)
// 0101000000075F984C153052C06B2BF697DD2B4640  POINT(-72.7513 44.3427)
// 010100000087A757CA323052C0FD87F4DBD7294640  POINT(-72.7531 44.3269)
// 0101000000AEB6627FD92D52C00E4FAF9465304640  POINT(-72.7164 44.3781)
// 01010000007B14AE47E11A52C0D7A3703D0A374640  POINT(-72.42 44.43)
// 0101000000B5A679C7292652C0B98D06F016204640  POINT(-72.5963 44.2507)
// 01010000005B423EE8D91C52C0B1BFEC9E3C144640  POINT(-72.4508 44.1581)
// 0101000000287E8CB96B3552C0B84082E2C7084640  POINT(-72.8347 44.0686)
// 0101000000865AD3BCE32452C07B14AE47E11A4640  POINT(-72.5764 44.21)
var _ = spec.Suite("WKB Reader", func(c *spec.C) {
	r := NewWKBReader()
	defer r.Destroy()

	c.It("should read a simple WKB to a geometry", func(c *spec.C) {
		// POINT(-72.4303 44.3664)
		b, err := hex.DecodeString("010100000001DE02098A1B52C090A0F831E62E4640")
		c.Assert(err).IsNil()

		p, err := r.Read(b)
		c.Assert(err).IsNil()
		c.Assert(p).NotNil()
	})
})

var _ = spec.Suite("WKB Writer", func(c *spec.C) {
	r := NewWKBReader()
	defer r.Destroy()

	w := NewWKBWriter()
	defer w.Destroy()

	c.It("should write a geometry to WKB", func(c *spec.C) {
		point, _ := GeomFromWKT("POINT(1.2345 2.3456)")
		b := w.Write(point)

		p, err := r.Read(b)
		c.Assert(err).IsNil()
		c.Assert(p).NotNil()
		c.Assert(p.geom).NotNil()
	})
})
