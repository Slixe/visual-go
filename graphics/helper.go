package graphics

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

var OffsetX = float32(0)
var OffsetY = float32(0)

func DrawTexture(texture rl.Texture2D, posX float32, posY float32, width float32, height float32, rotation float32) {
	DrawTexturePro(texture, posX, posY, width, height, rotation, rl.RayWhite)
}

func DrawTexturePro(texture rl.Texture2D, posX float32, posY float32, width float32, height float32, rotation float32, color rl.Color) {
	rect := rl.Rectangle{Y: OffsetY + posY, X: OffsetX + posX, Width: float32(texture.Width), Height: float32(texture.Height)}
	dst := rl.Rectangle{Y: OffsetY + posY, X: OffsetX + posX, Width: width, Height: height}
	rl.DrawTexturePro(texture, rect, dst, rl.Vector2{X: 0, Y: 0}, rotation, color)
}

func CreateRectangle(component structures.BaseComponent) rl.Rectangle {
	return rl.NewRectangle(OffsetX + component.PosX, OffsetY + component.PosY, component.Width, component.Height)
}

func DrawTextDefaultFont(text string, posX int, posY int, fontSize int, color rl.Color) {
	rl.DrawText(text, int(OffsetX) + posX, int(OffsetY) + posY, fontSize, color)
}

func DrawText(font rl.Font, text string, posX float32, posY float32, fontSize float32, spacing float32, color rl.Color) {
	rl.DrawTextEx(font, text, rl.Vector2{X: OffsetX + posX, Y: OffsetY + posY}, fontSize, spacing, color)
}

func DrawLine(startX int, startY int, endX int, endY int, color rl.Color) {
	rl.DrawLine(int(OffsetX) + startX, int(OffsetY) + startY, endX, endY, color)
}

func LoadFont(fontPath string, fontSize int) *rl.Font {
	return rl.LoadFontEx(fontPath, fontSize, nil, 256)
}