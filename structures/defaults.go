package structures

import rl "github.com/DankFC/raylib-goplus/raylib"

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

func (b *BaseInputField) SetEditable(value bool) {
	b.Editable = value
}

func (b BaseInputField) IsEditable() bool {
	return b.Selected && b.Editable
}

func (b BaseInputField) GetText() string {
	return b.Text
}

func (b *BaseInputField) SetText(value string) {
	b.Text = value
}

func (b *BaseInputField) SetMaxCharacters(value int) {
	b.MaxCharacters = value
}

func (b BaseInputField) GetMaxCharacters() int {
	return b.MaxCharacters
}

func (b *BaseInputField) HandleKey() {
	key := rl.GetKeyPressed()
	if len(b.Text) < b.MaxCharacters {
		strKey := string(key)

		if key >= 32 && key <= 126 {
			b.Text += strKey
			b.Callback(b)
		}
	}

	b.timer += rl.GetFrameTime()
	if len(b.Text) > 0 && (rl.IsKeyPressed(259) || (b.timer > 0.1 && rl.IsKeyDown(259))) {
		b.Text = b.Text[:len(b.Text) - 1]
		b.timer = 0
		b.Callback(b)
	}
}

func (b BaseClickable) OnClicked() {}