package midf

import (
	"fmt"
)

type IMiGeometry interface {
	GetGeoType() string
	MiString() string
}

type EmptyGeometry struct {
}

func (e EmptyGeometry) GetGeoType() string {
	return "EMPTY"
}
func (e EmptyGeometry) MiString() string {
	return "NONE" + STR_NEWLINE
}

type MiPoint struct {
	X float64
	Y float64
}

func NewMiPoint(x float64, y float64) MiPoint {
	return MiPoint{x, y}
}

func (this MiPoint) GetGeoType() string {
	return "POINT"
}

func (this MiPoint) MiString() string {
	return fmt.Sprintf("POINT %v %v", this.X, this.Y)
}

type MiLine struct {
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
}

func NewMiLine(x1, y1, x2, y2 float64) MiLine {
	return MiLine{x1, y1, x2, y2}
}

func (this MiLine) GetGeoType() string {
	return "LINE"
}

func (this MiLine) MiString() string {
	return fmt.Sprintf("LINE %v %v %v %v\n", this.X1, this.Y1, this.X2, this.Y2)
}

type MiPolyline struct {
	Sections [][]MiPoint
}

func NewMiPolyline(points []MiPoint) MiPolyline {
	polyline := MiPolyline{}
	polyline.Sections = make([][]MiPoint, 1, 10)
	polyline.Sections[0] = points
	return polyline
}

func (this MiPolyline) GetGeoType() string {
	return "PLINE"
}

func (this MiPolyline) MiString() string {
	result := ""
	if this.Sections == nil {
		return "NONE"
	}
	secCount := len(this.Sections)
	if secCount == 0 {
		return "NONE"
	}
	if secCount == 1 {
		result = "PLINE" + STR_NEWLINE
	} else {
		result = fmt.Sprintf("PLINE"+STR_SPACE+"MULTIPLE"+STR_SPACE+"%v"+STR_NEWLINE, secCount)
	}
	for i := 0; i < secCount; i++ {
		secLen := len(this.Sections[i])
		result += fmt.Sprintf("%v"+STR_NEWLINE, secLen)
		for j := 0; j < secLen; j++ {
			result += fmt.Sprintf("%v"+STR_SPACE+"%v"+STR_NEWLINE, this.Sections[i][j].X, this.Sections[i][j].Y)
		}
	}
	return result
}

type MiRegion struct {
}

type MiArc struct {
}

type MiText struct {
}

type MiRectangle struct {
}

type MiRoundedRectangle struct {
}

type MiEllipse struct {
}
