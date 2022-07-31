package shared

import (
	"fmt"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

const (
	MaxGradeValue uint8 = 5

	Expert       uint8 = 5
	Candidate    uint8 = 4
	Grade1       uint8 = 3
	Grade2       uint8 = 2
	Grade3       uint8 = 1
	WithoutGrade uint8 = 0
)

var ErrInvalidGrade = fmt.Errorf("grade should be between 0 and %d", MaxGradeValue)

type Grade struct {
	value uint8
}

func NewGrade(value uint8) (Grade, error) {
	if value > MaxGradeValue {
		return Grade{}, ErrInvalidGrade
	}
	return Grade{value}, nil
}

func (g *Grade) NextGradeAchieved(endorsementCount uint) bool {
	switch g.value {
	case Candidate:
		return endorsementCount >= 40
	case Grade3:
		return endorsementCount >= 20
	case Grade2:
		return endorsementCount >= 14
	case Grade1:
		return endorsementCount >= 10
	case WithoutGrade:
		return endorsementCount >= 6
	default:
		return false
	}
}

func (g *Grade) HasNext() bool {
	return g.value < MaxGradeValue
}

func (g *Grade) Next() (Grade, error) {
	nextGrade, err := NewGrade(g.value + 1)
	if err != nil {
		return *g, err
	}
	return nextGrade, nil
}

func (g *Grade) HasPrevious() bool {
	return g.value > 0
}

func (g *Grade) Previous() (Grade, error) {
	previousGrade, err := NewGrade(g.value - 1)
	if err != nil {
		return *g, err
	}
	return previousGrade, nil
}

func (g *Grade) Export() uint8 {
	return g.value
}

func (g *Grade) Import(value uint8) {
	g.value = value
}

func (g Grade) ExportTo(ex seedwork.ExporterSetter[uint8]) {
	ex.SetState(g.value)
}
