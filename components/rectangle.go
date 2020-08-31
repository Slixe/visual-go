package components

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type Rectangle struct {
	structures.BaseComponent
	Color rl.Color
}

func CreateRectangle(color rl.Color, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f) *Rectangle {
	return &Rectangle{
		Color: color,
		BaseComponent: structures.BaseComponent{
			Func:posFunc,
		},
	}
}

func (r Rectangle) Show(graphics structures.IGraphics, app structures.IApp) {
	graphics.DrawRectangle(int(r.GetPosition().PosX), int(r.GetPosition().PosY), int(r.GetPosition().Width), int(r.GetPosition().Height), r.Color)
}