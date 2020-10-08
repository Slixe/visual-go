package components

import (
	"github.com/Slixe/visual-go/structures"
	"github.com/lachee/raylib-goplus/raylib"
)

type Label struct {
	structures.BaseComponent
	Label string
}

func CreateLabel(label string, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f) *Label {
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