package graphics

import (
	"github.com/Slixe/visual-go/structures"
	rl "github.com/lachee/raylib-goplus/raylib"
)

type ScrollBar struct {
	structures.BaseComponent
	Direction structures.ScrollDirection
}

func (s ScrollBar) Show(graphics structures.IGraphics, app structures.IApp) {
	graphics.DrawRectangle(int(s.GetPosition().PosX), int(s.GetPosition().PosY), int(s.GetPosition().Width), int(s.GetPosition().Height), rl.DarkGray)
	if s.Direction == structures.Vertical {
		height := s.GetPosition().Height / 4
		posY := int(graphics.GetScrollValue(structures.Vertical) / graphics.GetScrollMaxValue(structures.Vertical) * (s.GetPosition().Height - height) * -1)
		graphics.DrawRectangle(int(s.GetPosition().PosX), posY, 10, int(height), rl.LightGray)
	} else {
		width := s.GetPosition().Width / 4
		posX := int(graphics.GetScrollValue(structures.Horizontal) / graphics.GetScrollMaxValue(structures.Horizontal) * (s.GetPosition().Width - width) * -1)
		graphics.DrawRectangle(posX, int(s.GetPosition().PosY), int(width), 10, rl.LightGray)
	}
}
