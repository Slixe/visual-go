package structures

func (b BasePanel) GetComponents() []IComponent {
	return b.Components
}

func (b *BasePanel) AddComponent(component IComponent) {
	b.Components = append(b.Components, component)
}

func (b BaseComponent) GetPosition() ComponentPos {
	return b.cachePos
}

func (b *BaseComponent) UpdatePosition(graphics IGraphics, app IApp) {
	b.cachePos = b.Func(graphics, app)
}

func (b BaseSelectableComponent) IsSelected() bool {
	return b.Selected
}

func (b *BaseSelectableComponent) SetSelected(value bool) {
	b.Selected = value
}