package midf

type Row map[string]interface{}

type MiData struct {
	Row
	Geometry IMiGeometry
	Graphic  IMiGraphic
}

func NewMiData() *MiData {
	data := new(MiData)
	data.Row = make(map[string]interface{})
	data.Geometry = EmptyGeometry{}
	data.Graphic = EmptyMiGraphic{}
	return data
}

func (m *MiData) SetValue(fname string, fvalue interface{}) {
	m.Row[fname] = fvalue
}

func (m *MiData) SetGeometry(geo IMiGeometry) {
	m.Geometry = geo
}

func (m *MiData) SetGraphic(graphic IMiGraphic) {
	m.Graphic = graphic
}
