package panels

import (
	"github.com/Slixe/visual-go/structures"
	"github.com/kjk/flex"
)

type Panel struct {
	Components []structures.IComponent
	Layout *flex.Node
}

func CreatePanel() *Panel {
	return &Panel{}
}

func (panel Panel) Show(graphics structures.IGraphics, app structures.IApp) {}

func (panel Panel) GetComponents() []structures.IComponent {
	return panel.Components
}

func (panel *Panel) AddComponent(component structures.IComponent) {
	panel.Components = append(panel.Components, component)
}

func (panel *Panel) SetLayout(layout *flex.Node) {
	panel.Layout = layout
}

func (panel Panel) GetLayout() *flex.Node {
	return panel.Layout
}