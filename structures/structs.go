package structures

import rl "github.com/lachee/raylib-goplus/raylib"

type ScrollDirection int

const (
	Horizontal ScrollDirection = iota
	Vertical
)

type BasePanel struct {
	Components []IComponent
	allowVerticalScroll bool
	allowHorizontalScroll bool
}

type BaseComponent struct {
	IComponent
	Func     func(graphics IGraphics, app IApp) Vector4f
	cachePos Vector4f
}

type BaseInputField struct {
	BaseSelectable
	Editable      bool
	Text          string
	MaxCharacters int
	Callback      func(field IInputField)
	timer float32
}

type BaseClickable struct {
	BaseComponent
	CallbackFunc func(clickable IClickable)
	clickedPos rl.Vector2
}

type BaseSelectable struct {
	BaseComponent
	Selected bool
}

type Vector4f struct {
	PosX   float32
	PosY   float32
	Width  float32
	Height float32
}