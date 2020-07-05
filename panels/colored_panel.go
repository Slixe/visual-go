package panels

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
	"github.com/kjk/flex"
)

type ColoredPanel struct {
	Color      rl.Color
	Components []structures.IComponent
	Layout *flex.Node
}

func CreateColoredPanel(color rl.Color) *ColoredPanel {
	return &ColoredPanel{
		Color: color,
	}
}

func (panel ColoredPanel) Show(graphics structures.IGraphics, app structures.IApp) {
	rl.ClearBackground(panel.Color)
}

func (panel ColoredPanel) GetComponents() []structures.IComponent {
	return panel.Components
}

func (panel *ColoredPanel) AddComponent(component structures.IComponent) {
	panel.Components = append(panel.Components, component)
}

func (panel *ColoredPanel) SetLayout(layout *flex.Node) {
	panel.Layout = layout
}

func (panel ColoredPanel) GetLayout() *flex.Node {
	return panel.Layout
}