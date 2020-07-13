package structures

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/kjk/flex"
)

type IApp interface {
	Close()
	CalculateLayout()
	SetMainLayout(node *flex.Node)
	GetGlobalPanel() IPanel
	SetGlobalPanel(panel IPanel)
	GetWidth() int
	GetHeight() int
	SetPanel(layoutName string, panel IPanel) bool
	RemovePanel(panel IPanel)
	GetPanels() []IPanel
	ClearPanels()
	RegisterLayout(layoutName string, node *flex.Node)
	RemoveLayout(layoutName string) *flex.Node
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
	GetComponents() []IComponent
	AddComponent(component IComponent)
}

type IGraphics interface {
	DrawTexture(texture rl.Texture2D, posX float32, posY float32, color rl.Color)
	DrawTexturePro(texture rl.Texture2D, posX float32, posY float32, width float32, height float32, rotation float32, color rl.Color)
	DrawTextDefaultFont(text string, posX int, posY int, fontSize int, color rl.Color)
	DrawText(font rl.Font, text string, posX float32, posY float32, fontSize float32, spacing float32, color rl.Color)
	DrawLine(startX int, startY int, endX int, endY int, color rl.Color)
	CreateRectangle(component ComponentPos) rl.Rectangle
	DrawRectangle(posX int, posY int, width int, height int, color rl.Color)
	GetWidth() float32
	GetHeight() float32
	GetPosX() float32
	GetPosY() float32
}

type IComponent interface {
	IShow
	GetPosition() ComponentPos
	UpdatePosition(graphics IGraphics, app IApp)
}

type ISelectableComponent interface {
	IComponent
	SetSelected(value bool)
	IsSelected() bool
}

type IInputField interface {
	SetEditable(value bool)
	IsEditable() bool
	GetText() string
	SetText(value string)
	SetMaxCharacters(value int)
	GetMaxCharacters() int
}

type IClickable interface {
	IComponent
	OnClicked()
}

type BasePanel struct {
	Components []IComponent
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
	Label string
	Callback func(label IClickable)
}

type BaseComponent struct {
	Func func(graphics IGraphics, app IApp) ComponentPos
	cachePos ComponentPos
}

type ComponentPos struct {
	PosX   float32
	PosY   float32
	Width  float32
	Height float32
}

type BaseSelectableComponent struct {
	BaseComponent
	Selected bool
}