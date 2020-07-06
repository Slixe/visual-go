package components

import (
	"github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type CheckBox struct {
	structures.BaseComponent
	Label    string
	Checked  bool
	Callback func(box CheckBox)
}

func CreateCheckBox(label string, checked bool, posX float32, posY float32, width float32, height float32, callback func(box CheckBox)) *CheckBox {
	return &CheckBox{
		BaseComponent: structures.BaseComponent{
			PosX:   posX,
			PosY:   posY,
			Width:  width,
			Height: height,
		},
		Label:    label,
		Checked:  checked,
		Callback: callback,
	}
}

func (checkBox *CheckBox) Show(graphics structures.IGraphics, app structures.IApp) {
	c := raylib.GuiCheckBox(graphics.CreateRectangle(checkBox.BaseComponent), checkBox.Label, checkBox.Checked)
	if c != checkBox.Checked {
		checkBox.Checked = c
		checkBox.Callback(*checkBox)
	} else {
		checkBox.Checked = c
	}
}