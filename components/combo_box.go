package components

import (
	"github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
	"strings"
)

type ComboBox struct {
	structures.BaseComponent
	Values []string
	Active int
	Callback func(box ComboBox)
}

func CreateComboBox(values []string, posX float32, posY float32, width float32, height float32, callback func(box ComboBox)) *ComboBox {
	return &ComboBox{
		BaseComponent: structures.BaseComponent{
			PosX:   posX,
			PosY:   posY,
			Width:  width,
			Height: height,
		},
		Values: values,
		Active: 0,
		Callback: callback,
	}
}

func (box *ComboBox) Show(graphics structures.IGraphics, app structures.IApp) {
	selected := raylib.GuiComboBox(graphics.CreateRectangle(box.BaseComponent), strings.Join(box.Values, ";"), box.Active)
	if box.Active != selected {
		box.Active = selected
		box.Callback(*box)
	} else {
		box.Active = selected
	}
}

func (box ComboBox) GetBaseComponent() structures.BaseComponent {
	return box.BaseComponent
}