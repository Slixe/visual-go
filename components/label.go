package components

import (
	"github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type Label struct {
	structures.BaseComponent
	Label string
	Color raylib.Color
}

func CreateLabel(label string, color raylib.Color, posX float32, posY float32, width float32, height float32) *Label {
	return &Label{
		BaseComponent: structures.BaseComponent{
			PosX:   posX,
			PosY:   posY,
			Width:  width,
			Height: height,
		},
		Label: label,
		Color: color,
	}
}

func (label Label) Show(graphics structures.IGraphics, app structures.IApp) {
	raylib.GuiLabel(graphics.CreateRectangle(label.BaseComponent), label.Label)
}

func (label Label) GetBaseComponent() structures.BaseComponent {
	return label.BaseComponent
}
