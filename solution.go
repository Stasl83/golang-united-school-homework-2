package square

import "math"

type intMyCustomType = int

const (
	SidesCircle   intMyCustomType = 0
	SidesTriangle intMyCustomType = 3
	SidesSquare   intMyCustomType = 4
)

func CalcSquere(sideLen float64, sidesNum intMyCustomType) float64 {

	switch sidesNum {
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	case SidesTriangle:
		return math.Sqrt(3) * sideLen * sideLen / 4
	case SidesSquare:
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
