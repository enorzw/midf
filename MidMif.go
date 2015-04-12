package midf

import (
	"fmt"
	"os"
	"path"
)

type MidMif struct {
	Mif   *MifFile
	Mid   *MidFile
	Datas []*MiData
}

func (m *MidMif) AddData(data *MiData) {
	m.Datas = append(m.Datas, data)
}

func (m *MidMif) WriteToFile() {
	miffullpath := path.Join(m.Mif.FilePath, m.Mif.FileName)
	midfullpath := path.Join(m.Mid.FilePath, m.Mid.FileName)

	mif, err := os.OpenFile(miffullpath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0660)
	if err != nil {
		panic(err)
	}
	defer mif.Close()

	mid, err := os.OpenFile(midfullpath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0660)
	if err != nil {
		panic(err)
	}
	defer mid.Close()

	mif.WriteString(m.Mif.Head.String())
	midWCount := 0

	for i := 0; i < len(m.Datas); i++ {
		data := m.Datas[i]
		wdata := make([]interface{}, m.Mif.Head.columnNumber, 10)
		for field, value := range data.Row {
			if index, ok := m.Mif.Head.columnNameMap[field]; ok {
				wdata[index] = value
			} else {
				panic("字段错误:" + field + ",mif中未定义该字段")
			}
		}

		for i := 0; i < len(wdata); i++ {
			var length int
			switch wdata[i].(type) {
			case string:
				length, _ = mid.WriteAt([]byte(fmt.Sprintf("\"%v\""+m.Mif.Head.Delimiter, wdata[i])), int64(midWCount))
			default:
				length, _ = mid.WriteAt([]byte(fmt.Sprintf("%v"+m.Mif.Head.Delimiter, wdata[i])), int64(midWCount))
			}
			midWCount += length
		}

		length, _ := mid.WriteAt([]byte("\n"), int64(midWCount))
		midWCount += length
	}

}
