package compgeo

import (
	"errors"
	"math"
	"strconv"
)

// A DCELPoint is just a 3-dimensional point.
type DCELPoint [3]float64

// X :
// Get the value of this point on the x axis
func (dp DCELPoint) X() float64 {
	return dp[0]
}

// Y :
// Get the value of this point on the y axis
func (dp DCELPoint) Y() float64 {
	return dp[1]
}

// Z :
// Get the value of this point on the z axis
func (dp DCELPoint) Z() float64 {
	return dp[2]
}

// A DCELEdge represents an edge within a DCEL,
// specifically a half edge, which maintains
// references to it's origin vertex, the face
// it bounds, the half edge sharing its space
// bounding its adjacent face, and the previous
// and following edges which bound its face.
type DCELEdge struct {
	// Origin is the vertex this edge starts at
	Origin *DCELPoint
	// Face is the index within Faces that this
	// edge wraps around
	Face *DCELFace
	// Next and Prev are the edges following and
	// preceding this edge that also wrap around
	// Face
	Next *DCELEdge
	Prev *DCELEdge
	// Twin is the half edge who points to this
	// half-edge's origin, and respectively whose
	// origin this half-edge points to.
	Twin *DCELEdge
}

// A DCELFace points to the edges on its inner and
// outer portions. Any given face may have either
// of these values be nil, but never both.
type DCELFace struct {
	Outer, Inner *DCELEdge
}

// A DCEL is a structure representin arbitrary plane
// divisions and 3d polytopes. Its values are relatively
// self-explanatory but constructing it is significantly
// harder.
type DCEL struct {
	Vertices []DCELPoint
	// outEdges[0] is the (an) edge in HalfEdges whose
	// orgin is Vertices[0]
	OutEdges  []*DCELEdge
	HalfEdges []DCELEdge
	// The first value in a face is the outside component
	// of the face, the second value is the inside component
	Faces []DCELFace
}

// DCELEdgeTwin can obtain a given edge index's twin
// without accessing the edge itself, for index
// manipulation, or for initially setting the Twins
// in construction.
//
// Hopeful Mandate: twin edges come in pairs
// if i is even, then, i+1 is its pair,
// and otherwise i-i is its pair.
func DCELEdgeTwin(i int) int {
	if i%2 == 0 {
		return i + 1
	}
	return i - 1
}

// FullEdge returns the ith edge in the form of its
// two vertices
func (d *DCEL) FullEdge(i int) [2]*DCELPoint {
	e := d.HalfEdges[i]
	e2 := e.Twin
	return [2]*DCELPoint{
		e.Origin,
		e2.Origin}
}

// MaxX returns the Maximum of all X values
func (d *DCEL) MaxX() float64 {
	return d.Max(0)
}

// MaxY returns the Maximum of all Y values
func (d *DCEL) MaxY() float64 {
	return d.Max(1)
}

// MaxZ returns the Maximum of all Z values
func (d *DCEL) MaxZ() float64 {
	return d.Max(2)
}

// Max functions iterate through vertices
// to find the maximum value along a given axis
// in the DCEL
func (d *DCEL) Max(i int) (x float64) {
	for _, p := range d.Vertices {
		if p[i] > x {
			x = p[i]
		}
	}
	return x
}

// MinX returns the Minimum of all X values
func (d *DCEL) MinX() float64 {
	return d.Min(0)
}

// MinY returns the Minimum of all Y values
func (d *DCEL) MinY() float64 {
	return d.Min(1)
}

// MinZ returns the Minimum of all Z values
func (d *DCEL) MinZ() float64 {
	return d.Min(2)
}

// Min functions iterate through vertices
// to find the maximum value along a given axis
// in the DCEL
func (d *DCEL) Min(i int) (x float64) {
	x = math.Inf(1)
	for _, p := range d.Vertices {
		if p[i] < x {
			x = p[i]
		}
	}
	return x
}

// AllEdges iterates through the edges surrounding
// a vertex and returns them all.
func (dc *DCEL) AllEdges(vertex int) []*DCELEdge {
	e1 := dc.OutEdges[vertex]
	edges := make([]*DCELEdge, 1)
	edges[0] = e1
	edge := e1.Twin.Next
	for edge != e1 {
		edges = append(edges, edge)
		edge = edge.Twin.Next
	}
	return edges
}

// PartitionVertexEdges partitions the edges of a vertex by
// whether they connect to a vertex greater or lesser than the
// given vertex with respect to a specific dimension
func (dc *DCEL) PartitionVertexEdges(vertex int, d int) ([]*DCELEdge, []*DCELEdge, error) {
	allEdges := dc.AllEdges(vertex)
	lesser := make([]*DCELEdge, 0)
	greater := make([]*DCELEdge, 0)
	v := dc.Vertices[vertex]
	if len(v) <= d {
		return lesser, greater, errors.New("DCEL's vertex does not support " + strconv.Itoa(d) + " dimensions")
	}
	checkAgainst := v[d]
	for _, e1 := range allEdges {
		e2 := e1.Twin
		// Potential issue:
		// Will something bad happen if there are multiple
		// elements with the same value in this dimension?
		if e2.Origin[d] <= checkAgainst {
			lesser = append(lesser, e1)
		} else {
			greater = append(greater, e1)
		}
	}
	return lesser, greater, nil
}