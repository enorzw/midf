package midf

type Row map[string]interface{}

type MiData struct {
	Row
	Geometry IMiGeometry
}

func NewMiData() *MiData {
	data := new(MiData)
	data.Row = make(map[string]interface{})
	data.Geometry = MiGeometry{}
	return data
}

func (m *MiData) Add(fname string, fvalue interface{}) {
	m.Row[fname] = fvalue
}
