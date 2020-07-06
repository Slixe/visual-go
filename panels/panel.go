package panels

import (
	"github.com/Slixe/visual-go/structures"
)

type Panel struct {
	Components []structures.IComponent
	structures.BasePanel
}

func CreatePanel() *Panel {
	return &Panel{}
}

func (panel Panel) Show(graphics structures.IGraphics, app structures.IApp) {}