package graphics

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type Graphics struct {
	posX        float32
	posY        float32
	maxWidth    float32
	maxHeight   float32
	scrollX     float32
	scrollY     float32
	allowScroll bool
	scrollable  bool
}

func (g Graphics) DrawTexture(texture rl.Texture2D, posX float32, posY float32, color rl.Color) {
	g.ValidatePos(&posX, &posY, 0, 0)
	rl.DrawTexture(texture, int(posX), int(posY), color)
}

func (g Graphics) DrawTexturePro(texture rl.Texture2D, posX float32, posY float32, width float32, height float32, rotation float32, color rl.Color) {
	g.ValidatePos(&posX, &posY, 0, 0)
	rect := rl.Rectangle{Y: 0, X: 0, Width: float32(texture.Width), Height: float32(texture.Height)}
	dst := rl.Rectangle{Y: posY, X: posX, Width: width, Height: height}
	rl.DrawTexturePro(texture, rect, dst, rl.Vector2{X: 0, Y: 0}, rotation, color)
}

func (g Graphics) DrawTextDefaultFont(text string, posX float32, posY float32, fontSize int, color rl.Color) {
	vec := g.MeasureText(*rl.GetFontDefault(), text, float32(fontSize), 0)
	g.ValidatePos(&posX, &posY, vec.X, vec.Y)
	rl.DrawText(text, int(posX), int(posY), fontSize, color)
}

func (g Graphics) DrawText(font rl.Font, text string, posX float32, posY float32, fontSize float32, spacing float32, color rl.Color) {
	vec := g.MeasureText(font, text, fontSize, spacing)
	g.ValidatePos(&posX, &posY, vec.X, vec.Y)
	rl.DrawTextEx(font, text, rl.Vector2{X: posX, Y: posY}, fontSize, spacing, color)
}

func (g Graphics) DrawLine(startX int, startY int, endX int, endY int, color rl.Color) {
	posX := float32(startX)
	posY := float32(startY)
	g.ValidatePos(&posX, &posY, float32(endX), float32(endY))
	endX += int(posX)
	endY += int(posY)

	if endX > int(g.maxWidth) {
		endX = int(g.maxWidth)
	}

	if endY > int(g.maxHeight) {
		endY = int(g.maxHeight)
	}

	rl.DrawLine(int(posX), int(posY), endX, endY, color)
}

func (g Graphics) CreateRectangle(component structures.Vector4f) rl.Rectangle {
	g.ValidatePos(&component.PosX, &component.PosY, component.Width, component.Height)
	return rl.NewRectangle(component.PosX, component.PosY, component.Width, component.Height)
}

func (g Graphics) DrawRectangle(posX int, posY int, width int, height int, color rl.Color) {
	x := float32(posX)
	y := float32(posY)

	g.ValidatePos(&x, &y, float32(width), float32(height))
	rl.DrawRectangle(int(x), int(y), width, height, color)
}

func (g Graphics) MeasureText(font rl.Font, text string, fontSize float32, spacing float32) rl.Vector2 {
	return rl.MeasureTextEx(font, text, fontSize, spacing)
}

func (g *Graphics) UpdateScroll(value float32) {
		g.scrollX += value
}

func (g *Graphics) ValidatePos(posX *float32, posY *float32, width float32, height float32) {
	if *posX < 0 {
		*posX = 0
	}

	if *posY < 0 {
		*posY = 0
	}

	if !g.allowScroll {
		if *posX+width > g.maxWidth {
			*posX = g.maxWidth - width
		}

		if *posY+height > g.maxHeight {
			*posY = g.maxHeight - height
		}
	} else {
		*posX += g.scrollX
		*posY += g.scrollY
	}

	*posX += g.posX
	*posY += g.posY
}

func (g Graphics) IsInArea(pos structures.Vector4f, posX, posY float32) bool {
	g.ValidatePos(&pos.PosX, &pos.PosY, pos.Width, pos.Height)
	return g.GetPosX() + pos.PosX <= posX && g.GetPosX() + pos.PosX + pos.Width >= posX && g.GetPosY() + pos.PosY <= posY && g.GetPosY() + pos.PosY + pos.Height >= posY
}

func (g Graphics) IsInArea2(posX, posY float32) bool {
	return g.GetPosX() <= posX && g.GetPosX() + g.GetWidth() >= posX && g.GetPosY() <= posY && g.GetPosY() + g.GetHeight() >= posY
}

func (g Graphics) ShouldRender(pos structures.Vector4f) bool {
	g.ValidatePos(&pos.PosX, &pos.PosY, pos.Width, pos.Height)
	return pos.PosX <= g.GetWidth() && pos.PosY <= g.GetHeight() && pos.PosX + pos.Width >= g.GetPosX() && pos.PosY + pos.Height >= g.GetPosY()
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

func (g *Graphics) ShouldAllowScroll(value bool) {
	g.allowScroll = value
}

func (g Graphics) AllowScroll() bool {
	return g.allowScroll
}