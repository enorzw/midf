package midf

import (
	"fmt"
)

type IMiGraphic interface {
	GetGraphicType() string
	MiString() string
}

type EmptyMiGraphic struct {
}

func (this EmptyMiGraphic) GetGraphicType() string {
	return "EMPTY"
}

func (this EmptyMiGraphic) MiString() string {
	return ""
}

type MiSymbol struct {
	IsEmpty bool
	Shape   int
	Color   int
	Size    int
}

func (this MiSymbol) GetGraphicType() string {
	return "SYMBOL"
}

func (this MiSymbol) MiString() string {
	if this.IsEmpty {
		return ""
	} else {
		return fmt.Sprintf("SYMBOL(%v,%v,%v)\n", this.Shape, this.Color, this.Size)
	}
}

type MiPen struct {
	IsEmpty bool
	Width   int
	Pattern int
	Color   int
}

func (this MiPen) GetGraphicType() string {
	return "PEN"
}

func (this MiPen) MiString() string {
	if this.IsEmpty {
		return ""
	} else {
		return fmt.Sprintf("PEN(%v,%v,%v)\n", this.Width, this.Pattern, this.Color)
	}
}

type MiFont struct {
}

type MiBrush struct {
}
