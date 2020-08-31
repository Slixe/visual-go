package structures

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/kjk/flex"
)

type ScrollDirection int

const (
	Horizontal ScrollDirection = iota
	Vertical
)

type IApp interface {
	Close()
	CalculateLayout()
	SetMainLayout(node *flex.Node)
	/*GetGlobalPanel() IPanel
	SetGlobalPanel(panel IPanel)*/
	GetWidth() int
	GetHeight() int
	SetPanel(layoutName string, panel IPanel)
	RemovePanel(panel IPanel)
	GetPanels() []IPanel
	ClearPanels()
	RegisterLayout(layoutName string, node *flex.Node)
	RegisterLayoutChild(layoutName string, node *flex.Node, parent *flex.Node)
	RemoveLayout(layoutName string) *flex.Node
	NewLayout() *flex.Node
	/*AddGlobalComponent(component IComponent)
	GetGlobalComponents() []IComponent
	ClearGlobalComponents()*/
	SetWindowTitle(title string)
	SetWindowSize(width int, height int)
	SetWindowIcon(imagePath string)
	SetGuiFont(font *rl.Font)
	SetTargetFPS(targetFPS int)
	GetFPS() int
}

type IShow interface {
	Show(graphics IGraphics, app IApp)
}

type IPanel interface {
	IShow
	GetComponents() []IComponent
	AddComponent(component IComponent)
}

type IScrollable interface {
	IPanel
	SetScrollable(value bool)
	IsScrollable() bool
}

type IGraphics interface {
	DrawTexture(texture rl.Texture2D, posX float32, posY float32, color rl.Color)
	DrawTexturePro(texture rl.Texture2D, posX float32, posY float32, width float32, height float32, rotation float32, color rl.Color)
	DrawTextDefaultFont(text string, posX float32, posY float32, fontSize int, color rl.Color)
	DrawText(font rl.Font, text string, posX float32, posY float32, fontSize float32, spacing float32, color rl.Color)
	DrawLine(startX int, startY int, endX int, endY int, color rl.Color)
	CreateRectangle(component Vector4f) rl.Rectangle
	DrawRectangle(posX int, posY int, width int, height int, color rl.Color)
	MeasureText(font rl.Font, text string, fontSize float32, spacing float32) rl.Vector2
	UpdateScroll(value float32)
	ValidatePos(posX *float32, posY *float32, width float32, height float32)
	IsInArea(pos Vector4f, posX, posY float32) bool
	IsInArea2(posX, posY float32) bool
	ShouldRender(pos Vector4f) bool
	GetWidth() float32
	GetHeight() float32
	GetPosX() float32
	GetPosY() float32
	ShouldAllowScroll(value bool)
	AllowScroll() bool

}

type IComponent interface {
	IShow
	GetPosition() Vector4f
	UpdatePosition(graphics IGraphics, app IApp)
}

type ISelectable interface {
	IComponent
	SetSelected(value bool)
	IsSelected() bool
}

type IInputField interface {
	ISelectable
	SetEditable(value bool)
	IsEditable() bool
	GetText() string
	SetText(value string)
	SetMaxCharacters(value int)
	GetMaxCharacters() int
}

type IClickable interface {
	IComponent
	Callback()
}