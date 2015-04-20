package midf

import (
	"fmt"
	"os"
	//"path"
	"strings"
	"testing"
	"unicode"
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
	midf.AddColumn("feaid", "char(13)")
	midf.AddColumn("age", "char(13)")
	midf.AddColumn("name", "char(13)")
	//midf.AddColumn("class", "char(13)")
	midf.AddColumnT("class", COLUMNTYPE_CHAR, 20)

	m := NewMiData()
	m.SetValue("feaid", "1")
	m.SetValue("age", 10)
	m.SetValue("name", "shabi")
	m.SetValue("class", 1.1)
	m.Geometry = NewMiPoint(112.123, 54.232)
	midf.AddData(m)

	m = NewMiData()
	m.SetValue("feaid", "2")
	m.SetValue("age", 15)
	m.SetValue("name", "zhongwei")
	m.SetValue("class", 1.12312)
	midf.AddData(m)

	m = NewMiData()
	m.SetValue("feaid", "2")
	m.SetValue("age", 15)
	m.SetValue("name", "zhongwei")
	m.SetValue("class", 1.12312)
	m.Geometry = NewMiLine(112.123, 54.232, 112.123, 54.232)
	midf.AddData(m)

	m = NewMiData()
	m.SetValue("feaid", "2")
	m.SetValue("age", 15)
	m.SetValue("name", "zhongwei")
	m.SetValue("class", 1.12312)
	line := NewMiPolyline([]MiPoint{
		NewMiPoint(112.123, 54.232),
		NewMiPoint(112.123, 54.232),
		NewMiPoint(112.123, 54.232),
		NewMiPoint(112.123, 54.232),
	})

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

func TestOmmm(t *testing.T) {

	p, errs := ReadWkt("LINESTRING (30.23423 10.234, 10 30, 40 40)")

	if errs != nil {
		fmt.Println(errs)
	}

	fmt.Println(p.MiString())

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '.'
	}
	polygon := "  POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10),(20 30, 35 35, 30 20, 20 30))"
	fmt.Println("-----------------")
	re1 := strings.Split(polygon, "),(")
	fmt.Println(re1)
	re2 := strings.Split(polygon, "((")
	fmt.Println(re2)
	fmt.Println("-----------------")
	polygon = strings.Trim(polygon, " ")
	fmt.Println(strings.HasPrefix(polygon, "POLYGON"))
	fmt.Println("-----------------")

	index := strings.Index(polygon, "),(")
	first := polygon[:index]
	fmt.Println(index)
	fmt.Println(first)
	fmt.Println(polygon[index+3:])
	outter := strings.FieldsFunc(polygon[:index], f)
	innner := strings.FieldsFunc(polygon[index+3:], f)
	fmt.Println(outter, "\n", innner)
	fmt.Printf("Fields are: %q", strings.FieldsFunc("POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10),(20 30, 35 35, 30 20, 20 30))", f))
}
