package midf

type MiGeometry struct {
}

type MiFont struct {
}

type MiBrush struct {
}

type MiPen struct {
}

type MiSymbol struct {
}

type IMiGeometry interface {
}

type MiPoint struct {
	X      float64
	Y      float64
	Symbol MiSymbol
}

func (this *MiPoint) MiString() string {
	buffer := ""
	buffer += "POINT "
	return buffer
}

type MiLine struct {
	X1  float64
	X2  float64
	Y1  float64
	Y2  float64
	Pen MiPen
}

type MiPolyline struct {
	Sections []MiPoint
	Pen      MiPen
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
