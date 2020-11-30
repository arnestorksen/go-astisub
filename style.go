package astisub

import "math"

type Justification struct {
	value int
}

const (
	JustificationUnchangedValue = 1
	JustificationLeftValue      = 2
	JustificationCenterValue    = 3
	JustificationRightValue     = 4
)

var (
	JustificationUnchanged = Justification{
		value: JustificationUnchangedValue,
	}
	JustificationLeft = Justification{
		value: JustificationLeftValue,
	}
	JustificationCentered = Justification{
		value: JustificationCenterValue,
	}
	JustificationRight = Justification{
		value: JustificationRightValue,
	}
)

type Origin struct {
	x float64
	y float64
}
type Extent struct {
	height float64
	width  float64
}
type Position struct {
	origin Origin
	extent Extent
}

func asPercent(i float64) uint8 {
	return uint8(math.RoundToEven(i * 100))
}
