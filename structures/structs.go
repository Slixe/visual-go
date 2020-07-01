package structures

type IApp interface {
	GetCurrentPanel() IPanel
	SetPanel(panel IPanel)
	GetWidth() int
	GetHeight() int
}

type IPanel interface {
	Show(app IApp)
	GetComponents() []IComponent
	AddComponent(component IComponent)
}

type IComponent interface {
	Show(app IApp)
}

type BaseComponent struct {
	PosX   float32
	PosY   float32
	Width  float32
	Height float32
}
