package main

import (
	"github.com/golang/geo/r3"
)

// Ray represents a ray in our ray tracer
type Ray struct {
	origin, direction r3.Vector
}

// PointAtParameter returns the color vector seen along the ray
func (ray Ray) PointAtParameter(t float64) r3.Vector {
	return ray.origin.Add(ray.direction.Mul(t))
}
