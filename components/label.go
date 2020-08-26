package components

import (
	"github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type Label struct {
	structures.BaseComponent
	Label string
}

func CreateLabel(label string, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.ComponentPos) *Label {
	return &Label{
		BaseComponent: structures.BaseComponent{
			Func: posFunc,
		},
		Label: label,
	}
}

func (label Label) Show(graphics structures.IGraphics, app structures.IApp) {
	raylib.GuiLabel(graphics.CreateRectangle(label.GetPosition()), label.Label)
}