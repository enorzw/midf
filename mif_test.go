package midf

import (
	"fmt"
	//"os"
	"path"
	"testing"
)

func TestConfig(t *testing.T) {

	file, _ := NewMifFile("File/aaa.mif")
	fmt.Println(file.FileName)
	fmt.Println(file.FilePath)
	file, err := NewMifFile("File/aaa.mid")
	fmt.Println(err.Error())
	fmt.Println(path.Ext("/home/use/aaa.x"))
	fmt.Println(path.Dir("/home/use/aaa.x"))
	fmt.Println(path.Split("/home/use/aaa.x"))
	head := NewMifHead()
	fmt.Println(head.String())
	head.Unique = []int{0, 1, 2}
	fmt.Println(head.String())

	fmt.Println()
	head.AddColumn("feaid", "char(10)")
	fmt.Println(head.String())
}

func TestOs(t *testing.T) {

	midf := NewMidMif("c")
	midf.Mif.Head.AddColumn("feaid", "char(13)")
	midf.Mif.Head.AddColumn("age", "char(13)")
	midf.Mif.Head.AddColumn("name", "char(13)")
	midf.Mif.Head.AddColumn("class", "char(13)")

	m := NewMiData()
	m.Add("feaid", "1")
	m.Add("age", 10)
	m.Add("name", true)
	m.Add("class", 1.1)
	// m.Row["feaid"] = "1"
	// m.Row["age"] = "10"
	// m.Row["name"] = "David"
	// m.Row["class"] = "2"
	midf.AddData(*m)
	// m.Values = make(map[string]interface{})
	// m.Values["feaid"] = "1"
	// m.Values["age"] = "10"
	// m.Values["name"] = "David"
	// m.Values["class"] = "2"
	// midf.AddData(m)

	// m1 := MiData{}
	// m1.Values = make(map[string]interface{})
	// m1.Values["feaid"] = "1"
	// m1.Values["age"] = "10"
	// m1.Values["name"] = "David"
	// m1.Values["class"] = "2"
	// midf.AddData(m1)

	midf.WriteToFile()
}
