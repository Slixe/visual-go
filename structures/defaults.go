package structures

import (
	rl "github.com/lachee/raylib-goplus/raylib"
)

func (b BasePanel) GetComponents() []IComponent {
	return b.Components
}

func (b *BasePanel) AddComponent(component IComponent) {
	b.Components = append(b.Components, component)
}

func (b *BasePanel) SetAllowScroll(value bool, direction ScrollDirection) {
	if direction == Vertical {
		b.allowVerticalScroll = value
	} else {
		b.allowHorizontalScroll = value
	}
}

func (b BasePanel) AllowScroll(direction ScrollDirection) bool {
	value := b.allowVerticalScroll
	if direction == Horizontal {
		value = b.allowHorizontalScroll
	}

	return value
}

func (b BaseComponent) GetPosition() Vector4f {
	return b.cachePos
}

func (b *BaseComponent) UpdatePosition(graphics IGraphics, app IApp) {
	b.cachePos = b.Func(graphics, app)
}

func (b BaseSelectable) IsSelected() bool {
	return b.Selected
}

func (b *BaseSelectable) SetSelected(value bool) {
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
	if !b.Editable || !b.Selected {
		return
	}

	key := rl.GetKeyPressed()
	if len(b.Text) < b.MaxCharacters {
		strKey := string(key)

		if key >= 32 && key <= 126 {
			b.Text += strKey
			b.Callback(b)
		}
	}

	b.timer += rl.GetFrameTime()
	if len(b.Text) > 0 && (rl.IsKeyPressed(259) || (b.timer > 0.1 && rl.IsKeyDown(259))) { //TODO more is deleted faster is
		b.Text = b.Text[:len(b.Text) - 1]
		b.timer = 0
		b.Callback(b)
	}
}

func (b *BaseClickable) Callback() {
	b.CallbackFunc(b)
}

func (b BaseClickable) GetClickedPosition() rl.Vector2 {
	return b.clickedPos
}

func (b *BaseClickable) SetClickedPosition(position rl.Vector2) {
	b.clickedPos = position
}