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

func (g Graphics) DrawTexture(texture rl.Texture2D, posX float32, posY float32, color rl.Color) {
	rl.DrawTexture(texture, int(g.posX + posX), int(g.posY + posY), color)
}

func (g Graphics) DrawTexturePro(texture rl.Texture2D, posX float32, posY float32, width float32, height float32, rotation float32, color rl.Color) {
	if posX + width > g.maxWidth {
		posX = g.maxWidth - width
	}

	if posY + height > g.maxHeight {
		posY = g.maxHeight - height
	}

	rect := rl.Rectangle{Y: 0, X: 0, Width: float32(texture.Width), Height: float32(texture.Height)}
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
	if endX > int(g.maxWidth) {
		endX = int(g.maxWidth)
	}

	if endY > int(g.maxHeight) {
		endY = int(g.maxHeight)
	}

	if startX > int(g.maxWidth) {
		startX = int(g.maxWidth)
	}

	if startY > int(g.maxHeight) {
		startY = int(g.maxHeight)
	}

	rl.DrawLine(int(g.posX) + startX, int(g.posY) + startY, int(g.posX) + endX, int(g.posY) + endY, color)
}

func (g Graphics) CreateRectangle(component structures.ComponentPos) rl.Rectangle {
	if component.PosX < 0 {
		component.PosX += g.maxWidth
	}

	if component.PosY < 0 {
		component.PosY += g.maxHeight
	}

	if component.PosX + component.Width > g.maxWidth {
		component.PosX = g.maxWidth - component.Width
	}

	if component.PosY + component.Height > g.maxHeight {
		component.PosY = g.maxHeight - component.Height
	}

	return rl.NewRectangle(g.posX + component.PosX, g.posY + component.PosY, component.Width, component.Height)
}

func (g Graphics) DrawRectangle(posX int, posY int, width int, height int, color rl.Color) {
	if posX + width > int(g.maxWidth) {
		width = int(g.maxWidth) - posX
	}

	if posY + height > int(g.maxHeight) {
		height = int(g.maxHeight) - posY
	}

	rl.DrawRectangle(int(g.posX) + posX, int(g.posY) + posY, width, height, color)
}

func (g Graphics) MeasureText(font rl.Font, text string, fontSize float32, spacing float32) rl.Vector2 {
	return rl.MeasureTextEx(font, text, fontSize, spacing)
}

func (g Graphics) GetWidth() float32 {
	return g.maxWidth
}

func (g Graphics) GetHeight() float32 {
	return g.maxHeight
}

func (g Graphics) GetPosX() float32 {
	return g.posX
}

func (g Graphics) GetPosY() float32 {
	return g.posY
}