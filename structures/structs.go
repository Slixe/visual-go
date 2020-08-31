package structures

type BasePanel struct {
	Components []IComponent
}

type BaseComponent struct {
	IComponent
	Func     func(graphics IGraphics, app IApp) Vector4f
	cachePos Vector4f
}

type BaseInputField struct {
	BaseSelectableComponent
	Editable      bool
	Text          string
	MaxCharacters int
	Callback      func(field IInputField)
	timer float32
}

type BaseClickable struct {
	BaseComponent
	Label        string
	CallbackFunc func(label IClickable)
}

type Vector4f struct {
	PosX   float32
	PosY   float32
	Width  float32
	Height float32
}

type BaseSelectableComponent struct {
	BaseComponent
	Selected bool
}