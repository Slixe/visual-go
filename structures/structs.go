package structures

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/kjk/flex"
)

type IApp interface {
	GetGlobalPanel() IPanel
	SetGlobalPanel(panel IPanel)
	GetWidth() int
	GetHeight() int
	SetPanel(name string, panel IPanel)
	RemovePanel(name string)
	GetPanel(name string) IPanel
	GetPanels() []IPanel
	ClearPanels()
	AddGlobalComponent(component IComponent)
	GetGlobalComponents() []IComponent
	ClearGlobalComponents()
	SetWindowTitle(title string)
	SetWindowSize(width int, height int)
	SetWindowIcon(imagePath string)
	SetGuiFont(font *rl.Font)
}

type IShow interface {
	Show(graphics IGraphics, app IApp)
}

type IPanel interface {
	IShow
	SetLayout(layout *flex.Node)
	GetLayout() *flex.Node
	GetComponents() []IComponent
	AddComponent(component IComponent)
}

type IGraphics interface {
	DrawTexture(texture rl.Texture2D, posX float32, posY float32, width float32, height float32, rotation float32)
	DrawTexturePro(texture rl.Texture2D, posX float32, posY float32, width float32, height float32, rotation float32, color rl.Color)
	DrawTextDefaultFont(text string, posX int, posY int, fontSize int, color rl.Color)
	DrawText(font rl.Font, text string, posX float32, posY float32, fontSize float32, spacing float32, color rl.Color)
	DrawLine(startX int, startY int, endX int, endY int, color rl.Color)
	CreateRectangle(component BaseComponent) rl.Rectangle
	GetWidth() float32
	GetHeight() float32
}

type IComponent interface {
	IShow
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
