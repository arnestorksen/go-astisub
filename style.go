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

type VerticalPosition struct {
	position float64
}

func newVerticalPositionFromRows(row uint8, totalRows uint8) VerticalPosition {
	return VerticalPosition{
		position: float64(row-1) / float64(totalRows),
	}
}

func newVerticalPositionFromPercentage(percent uint8) VerticalPosition {
	return VerticalPosition{
		position: float64(percent) / 100,
	}
}

func newVerticalPosition(position float64) VerticalPosition {
	return VerticalPosition{
		position: position,
	}
}

func (i VerticalPosition) asRow(totalRows uint8) uint8 {
	return uint8(math.RoundToEven(i.position * float64(totalRows)))
}

func (i VerticalPosition) asPercent() uint8 {
	return uint8(math.RoundToEven(i.position * 100))
}

func (i VerticalPosition) asAbsolute() float64 {
	return i.position
}
