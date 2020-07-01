package panels

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type Panel struct {
	Color      rl.Color
	Components []structures.IComponent
}

func CreatePanel(color rl.Color) *Panel {
	return &Panel{
		Color: color,
	}
}

func (panel Panel) Show(app structures.IApp) {
	rl.ClearBackground(panel.Color)
}

func (panel Panel) GetComponents() []structures.IComponent {
	return panel.Components
}

func (panel *Panel) AddComponent(component structures.IComponent) {
	panel.Components = append(panel.Components, component)
}
