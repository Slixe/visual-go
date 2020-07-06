package structures

import "github.com/kjk/flex"

func (b *BasePanel) SetLayout(layout *flex.Node) {
	b.Layout = layout
}

func (b BasePanel) GetLayout() *flex.Node {
	return b.Layout
}

func (b BasePanel) GetComponents() []IComponent {
	return b.Components
}

func (b *BasePanel) AddComponent(component IComponent) {
	b.Components = append(b.Components, component)
}

func (b BaseComponent) GetBaseComponent() BaseComponent {
	return b
}

func (b BaseSelectableComponent) IsSelected() bool {
	return b.Selected
}

func (b *BaseSelectableComponent) SetSelected(value bool) {
	b.Selected = value
}