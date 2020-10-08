package panels

import (
	"github.com/Slixe/visual-go/structures"
	rl "github.com/lachee/raylib-goplus/raylib"
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
	graphics.DrawRectangle(0, 0, int(graphics.GetWidth()), int(graphics.GetHeight()), panel.Color)
}