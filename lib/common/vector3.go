package common

import "fmt"

type Vector3 struct {
	X, Y, Z float64
}

func NewVector3(x float64, y float64, z float64) Vector3 {
	return Vector3{x, y, z}
}

func (v *Vector3) String() string {
	return fmt.Sprintf("<Vector3 x=%f y=%f z=%f>", v.X, v.Y, v.Z)
}

func (v *Vector3) Add(o *Vector3) Vector3 {
	return Vector3{X: v.X + o.X, Y: v.Y + o.Y, Z: v.Z + o.Z}
}

func (v *Vector3) Sub(o *Vector3) Vector3 {
	return Vector3{X: v.X - o.X, Y: v.Y - o.Y, Z: v.Z - o.Z}
}

func (v *Vector3) Multi(o *Vector3) Vector3 {
	return Vector3{X: v.X * o.X, Y: v.Y * o.Y, Z: v.Z * o.Z}
}

func (v *Vector3) MultiFloat64(m float64) Vector3 {
	return Vector3{X: v.X * m, Y: v.Y * m, Z: v.Z * m}
}

func (v *Vector3) Lerp(o *Vector3, t float64) Vector3 {
	return Vector3{
		X: Lerp(v.X, o.X, t),
		Y: Lerp(v.Y, o.Y, t),
		Z: Lerp(v.Z, o.Z, t),
	}
}
