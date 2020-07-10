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

func CreateLabel(label string, color raylib.Color, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.ComponentPos) *Label {
	return &Label{
		BaseComponent: structures.BaseComponent{
			Func: posFunc,
		},
		Label: label,
		Color: color,
	}
}

func (label Label) Show(graphics structures.IGraphics, app structures.IApp) {
	raylib.GuiLabel(graphics.CreateRectangle(label.GetPosition()), label.Label)
}