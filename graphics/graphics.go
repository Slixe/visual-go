package graphics

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type Graphics struct {
	posX float32
	posY float32
	maxWidth float32
	maxHeight float32
}

func (g Graphics) DrawTexture(texture rl.Texture2D, posX float32, posY float32, width float32, height float32, rotation float32) {
	g.DrawTexturePro(texture, posX, posY, width, height, rotation, rl.RayWhite)
}

func (g Graphics) DrawTexturePro(texture rl.Texture2D, posX float32, posY float32, width float32, height float32, rotation float32, color rl.Color) {
	rect := rl.Rectangle{Y: g.posY + posY, X: g.posX + posX, Width: float32(texture.Width), Height: float32(texture.Height)}
	dst := rl.Rectangle{Y: g.posY + posY, X: g.posX + posX, Width: width, Height: height}
	rl.DrawTexturePro(texture, rect, dst, rl.Vector2{X: 0, Y: 0}, rotation, color)
}

func (g Graphics) DrawTextDefaultFont(text string, posX int, posY int, fontSize int, color rl.Color) {
	rl.DrawText(text, int(g.posX) + posX, int(g.posY) + posY, fontSize, color)
}

func (g Graphics) DrawText(font rl.Font, text string, posX float32, posY float32, fontSize float32, spacing float32, color rl.Color) {
	rl.DrawTextEx(font, text, rl.Vector2{X: g.posX + posX, Y: g.posY + posY}, fontSize, spacing, color)
}

func (g Graphics) DrawLine(startX int, startY int, endX int, endY int, color rl.Color) {
	rl.DrawLine(int(g.posX) + startX, int(g.posY) + startY, endX, endY, color)
}

func (g Graphics) CreateRectangle(component structures.BaseComponent) rl.Rectangle {
	return rl.NewRectangle(g.posX + component.PosX, g.posY + component.PosY, component.Width, component.Height)
}

func (g Graphics) GetWidth() float32 {
	return g.maxWidth
}

func (g Graphics) GetHeight() float32 {
	return g.maxHeight
}