package structures

import rl "github.com/DankFC/raylib-goplus/raylib"

type IApp interface {
	GetCurrentPanel() IPanel
	SetPanel(panel IPanel)
	GetWidth() int
	GetHeight() int
	AddGlobalComponent(component IComponent)
	GetGlobalComponents() []IComponent
	SetWindowTitle(title string)
	SetWindowSize(width int, height int)
	SetWindowIcon(imagePath string)
	SetGuiFont(font rl.Font)
}

type IPanel interface {
	Show(app IApp)
	GetComponents() []IComponent
	AddComponent(component IComponent)
}

type IComponent interface {
	Show(app IApp)
	GetBaseComponent() BaseComponent
}

type ISelectableComponent interface {
	IComponent
	SetSelected(value bool)
	IsSelected() bool
}

type BaseComponent struct {
	PosX   float32
	PosY   float32
	Width  float32
	Height float32
}
