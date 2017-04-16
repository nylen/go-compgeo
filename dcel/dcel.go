package dcel

import (
	"fmt"
	"math"
)

const (
	ε = 1.0e-7
)

// A DCEL is a structure representin arbitrary plane
// divisions and 3d polytopes. Its values are relatively
// self-explanatory but constructing it is significantly
// harder.
type DCEL struct {
	Vertices  []*Vertex
	HalfEdges []*Edge
	// The first value in a face is the outside component
	// of the face, the second value is the inside component
	Faces []*Face
}

// New returns an empty DCEL with its inner
// fields initialized to empty slices, and a
// zeroth outside face.
func New() *DCEL {
	dc := new(DCEL)
	dc.Vertices = []*Vertex{}
	dc.HalfEdges = []*Edge{}
	dc.Faces = []*Face{NewFace()}
	return dc
}

// MaxX returns the Maximum of all X values
func (dc *DCEL) MaxX() float64 {
	return dc.Max(0)
}

// MaxY returns the Maximum of all Y values
func (dc *DCEL) MaxY() float64 {
	return dc.Max(1)
}

// MaxZ returns the Maximum of all Z values
func (dc *DCEL) MaxZ() float64 {
	return dc.Max(2)
}

// Max functions iterate through vertices
// to find the maximum value along a given axis
// in the DCEL
func (dc *DCEL) Max(i int) (x float64) {
	for _, p := range dc.Vertices {
		if p.Val(i) > x {
			x = p.Val(i)
		}
	}
	return x
}

// MinX returns the Minimum of all X values
func (dc *DCEL) MinX() float64 {
	return dc.Min(0)
}

// MinY returns the Minimum of all Y values
func (dc *DCEL) MinY() float64 {
	return dc.Min(1)
}

// MinZ returns the Minimum of all Z values
func (dc *DCEL) MinZ() float64 {
	return dc.Min(2)
}

// Min functions iterate through vertices
// to find the maximum value along a given axis
// in the DCEL
func (dc *DCEL) Min(i int) (x float64) {
	x = math.Inf(1)
	for _, p := range dc.Vertices {
		if p.Val(i) < x {
			x = p.Val(i)
		}
	}
	return x
}

// ScanFaces returns which index, if any, within dc matches f.
func (dc *DCEL) ScanFaces(f *Face) int {
	for i, f2 := range dc.Faces {
		if f2 == f {
			return i
		}
	}
	return -1
}

// CorrectDirectionality (rather innefficently)
// ensures that a face has the right clockwise/
// counter-clockwise orientation based on
// whether its chain is the inner or outer
// portion of a face.
func (dc *DCEL) CorrectDirectionality(f *Face) {
	// Inners need to be going CC
	// Outers need to be going Clockwise

	clock, err := f.Inner.IsClockwise()
	if err == nil && clock {
		f.Inner.Flip()
	} else {
		fmt.Println(err, clock)
	}
	clock, err = f.Outer.IsClockwise()
	if err == nil && !clock {
		f.Outer.Flip()

	}
}
