package components

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type Line struct {
structures.BaseComponent
Color rl.Color
}

func CreateLine(color rl.Color, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f) *Line {
	return &Line{
		Color: color,
		BaseComponent: structures.BaseComponent{
			Func:posFunc,
		},
	}
}

func (line Line) Show(graphics structures.IGraphics, app structures.IApp) {
	graphics.DrawLine(int(line.GetPosition().PosX), int(line.GetPosition().PosY), int(line.GetPosition().Width), int(line.GetPosition().Height), line.Color)
}