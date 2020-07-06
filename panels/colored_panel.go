package panels

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type ColoredPanel struct {
	Color      rl.Color
	structures.BasePanel
}

func CreateColoredPanel(color rl.Color) *ColoredPanel {
	return &ColoredPanel{
		Color: color,
	}
}

func (panel ColoredPanel) Show(graphics structures.IGraphics, app structures.IApp) {
	graphics.DrawRectangle(0, 0, app.GetWidth(), app.GetHeight(), panel.Color)
}