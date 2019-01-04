// pointLoc holds interfaces for point location.

package pointLoc

import "github.com/nylen/go-compgeo/dcel"

// LocatesPoints is an interface to represent point location
// queries.
type LocatesPoints interface {
	PointLocate(vs ...float64) (*dcel.Face, error)
}
