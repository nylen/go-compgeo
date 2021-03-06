package monotone

import (
	"github.com/nylen/go-compgeo/dcel"
	"github.com/nylen/go-compgeo/dcel/pointLoc/visualize"
	"github.com/nylen/go-compgeo/geom"
	"github.com/nylen/go-compgeo/search"
)

// CompEdge describes structs to satisfy search interfaces
// so edges can be put as values and keys into a binary search
// tree and compared horizontally
//
// See also: slab/compEdge.go
// these structures, or at least the Compare function portion
// should probably be in dcel

type edgeNode struct {
	v *dcel.Edge
}

func (en edgeNode) Key() search.Comparable {
	return compEdge{en.v}
}

func (en edgeNode) Val() search.Equalable {
	return valEdge{en.v}
}

type valEdge struct {
	*dcel.Edge
}

func (ve valEdge) Equals(e search.Equalable) bool {
	switch ve2 := e.(type) {
	case valEdge:
		return ve.Edge == ve2.Edge
	}
	return false
}

// We need to have our keys be CompEdges so
// they are comparable within a certain y range.
type compEdge struct {
	*dcel.Edge
}

func (ce compEdge) Compare(i interface{}) search.CompareResult {
	switch c := i.(type) {
	case compEdge:
		if visualize.VisualCh != nil {
			visualize.DrawLine(ce.Edge.Origin, ce.Edge.Twin.Origin)
			visualize.DrawLine(c.Edge.Origin, c.Edge.Twin.Origin)
		}
		if ce.Edge == c.Edge {
			return search.Equal
		}

		if geom.F64eq(ce.X(), c.X()) && geom.F64eq(ce.Y(), c.Y()) &&
			geom.F64eq(ce.Twin.X(), c.Twin.X()) && geom.F64eq(ce.Twin.Y(), c.Twin.Y()) {
			return search.Equal
		}
		y, err := ce.FindSharedPoint(c.Edge, 1)
		if err != nil {
		}
		p1, _ := ce.PointAt(1, y)
		p2, _ := c.PointAt(1, y)
		if p1[0] < p2[0] {
			return search.Less
		}
		return search.Greater
	}
	return ce.Edge.Compare(i)
}
