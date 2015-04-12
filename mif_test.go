package midf

import (
	"fmt"
	"os"
	//"path"
	"testing"
)

func TestConfig(t *testing.T) {

	// file, _ := NewMifFile("File/aaa.mif")
	// fmt.Println(file.FileName)
	// fmt.Println(file.FilePath)
	// file, err := NewMifFile("File/aaa.mid")
	// fmt.Println(err.Error())
	// fmt.Println(path.Ext("/home/use/aaa.x"))
	// fmt.Println(path.Dir("/home/use/aaa.x"))
	// fmt.Println(path.Split("/home/use/aaa.x"))
	// head := NewMifHead()
	// fmt.Println(head.String())
	// head.Unique = []int{0, 1, 2}
	// fmt.Println(head.String())

	// fmt.Println()
	// head.AddColumn("feaid", "char(10)")
	// fmt.Println(head.String())
}

func TestOs(t *testing.T) {

	midf := NewMidMif("c")
	midf.Mif.Head.Coordsys = "Earth Projection 1,0"
	midf.Mif.Head.AddColumn("feaid", "char(13)")
	midf.Mif.Head.AddColumn("age", "char(13)")
	midf.Mif.Head.AddColumn("name", "char(13)")
	midf.Mif.Head.AddColumn("class", "char(13)")

	m := NewMiData()
	m.Add("feaid", "1")
	m.Add("age", 10)
	m.Add("name", "shabi")
	m.Add("class", 1.1)
	m.Geometry = NewMiPoint(112.123, 54.232)
	midf.AddData(m)

	m = NewMiData()
	m.Add("feaid", "2")
	m.Add("age", 15)
	m.Add("name", "zhongwei")
	m.Add("class", 1.12312)
	m.Geometry = EmptyGeometry{}
	midf.AddData(m)

	m = NewMiData()
	m.Add("feaid", "2")
	m.Add("age", 15)
	m.Add("name", "zhongwei")
	m.Add("class", 1.12312)
	m.Geometry = NewMiLine(112.123, 54.232, 112.123, 54.232)
	midf.AddData(m)

	m = NewMiData()
	m.Add("feaid", "2")
	m.Add("age", 15)
	m.Add("name", "zhongwei")
	m.Add("class", 1.12312)
	line := NewMiPolyline([]MiPoint{
		NewMiPoint(112.123, 54.232),
		NewMiPoint(112.123, 54.232),
		NewMiPoint(112.123, 54.232),
		NewMiPoint(112.123, 54.232),
	})
	// line.Sections = append(line.Sections, []MiPoint{
	// 	NewMiPoint(112.123, 54.232),
	// 	NewMiPoint(112.123, 54.232),
	// 	NewMiPoint(112.123, 54.232),
	// 	NewMiPoint(112.123, 54.232),
	// })
	m.Geometry = line
	m.Graphic = MiPen{false, 1, 2, 0}
	midf.AddData(m)

	midf.WriteToFile()

	file, _ := os.OpenFile("c.mif", os.O_RDONLY, 0)
	defer file.Close()
	buf := make([]byte, 1000, 1000)
	file.Read(buf)
	fmt.Println(string(buf))

	file, _ = os.OpenFile("c.mid", os.O_RDONLY, 0)
	defer file.Close()
	buf = make([]byte, 1000, 1000)
	file.Read(buf)
	fmt.Println(string(buf))

}
