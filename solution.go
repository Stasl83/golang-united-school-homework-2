package square

import "math"

type intMyCustomType = int

func CalcSquere(sideLen float64, sidesNum intMyCustomType) float64 {
	const (
		c0 intMyCustomType = 0
		c3 intMyCustomType = 3
		c4 intMyCustomType = 4
	)
	switch sidesNum {
	case c0:
		return math.Pi * sideLen * sideLen
	case c3:
		return math.Sqrt(3) * sideLen * sideLen / 4
	case c4:
		return sideLen * sideLen
	default:
		return 0
	}
}

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
