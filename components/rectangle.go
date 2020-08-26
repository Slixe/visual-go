package components

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type Rectangle struct {
	structures.BaseComponent
	Color rl.Color
}

func CreateRectangle(color rl.Color, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.ComponentPos) *Rectangle {
	return &Rectangle{
		Color: color,
		BaseComponent: structures.BaseComponent{
			Func:posFunc,
		},
	}
}

func (line Rectangle) Show(graphics structures.IGraphics, app structures.IApp) {
	graphics.DrawRectangle(int(line.GetPosition().PosX), int(line.GetPosition().PosY), int(line.GetPosition().Width), int(line.GetPosition().Height), line.Color)
}