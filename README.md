# midf
This is a lib to read and write mid/mif file for golang.

##Install 

    > go get github.com/enorzw/midf

###How to use
```go
package main

import(
	"fmt"
	"github.com/enorzw/midf"
)

func main(){

	file := midf.NewMidMif("c")
	file.Mif.Head.Coordsys = "Earth Projection 1,0"
	file.AddColumn("feaid", "char(13)")
	file.AddColumn("age", "char(13)")
	file.AddColumn("name", "char(13)")
	file.AddColumn("class", "char(13)")
 
	m := NewMiData()
	m.Add("feaid", "1")
	m.Add("age", 28)
	m.Add("name", "enorzw")
	m.Add("class", 1.12312)
	m.Geometry = NewMiLine(112.123, 54.232, 112.123, 54.232)
	midf.AddData(m)

	m = NewMiData()
	m.Add("feaid", "2")
	m.Add("age", 40)
	m.Add("name", "Smith")
	m.Add("class", 1.12312)
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
}

```