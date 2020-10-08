package graphics

import (
	"fmt"
	"github.com/Slixe/visual-go/structures"
	rl "github.com/lachee/raylib-goplus/raylib"
	"math"
)

type Graphics struct {
	posX             float32
	posY             float32
	maxWidth         float32
	maxHeight        float32
	verticalScroll   Scroll
	horizontalScroll Scroll
	skipScrollPadding bool
}

type Scroll struct {
	allowed bool
	scrollable bool //will be used to draw scrollbar
	value     float32
	maxValue float32
	bar ScrollBar
	selected bool
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
	vec := MeasureText(*rl.GetFontDefault(), text, float32(fontSize), 0)
	g.ValidatePos(&posX, &posY, vec.X, vec.Y)
	rl.DrawText(text, int(posX), int(posY), fontSize, color)
}

func (g Graphics) DrawText(font rl.Font, text string, posX float32, posY float32, fontSize float32, spacing float32, color rl.Color) {
	vec := MeasureText(font, text, fontSize, spacing)
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

func (g *Graphics) ValidatePos(posX *float32, posY *float32, width float32, height float32) {
	if *posX < 0 {
		*posX = 0
	}

	if *posY < 0 {
		*posY = 0
	}

	if !g.skipScrollPadding {
		if !g.horizontalScroll.allowed && !g.verticalScroll.allowed {
			if *posX+width > g.maxWidth {
				*posX = g.maxWidth - width
			}

			if *posY+height > g.maxHeight {
				*posY = g.maxHeight - height
			}
		} else {
			*posX += g.horizontalScroll.value
			*posY += g.verticalScroll.value
		}
	}

	*posX += g.posX
	*posY += g.posY
}

func (g Graphics) IsInArea(pos structures.Vector4f, posX, posY float32) bool { //posX / posY are global positions
	g.ValidatePos(&pos.PosX, &pos.PosY, pos.Width, pos.Height)

	return pos.PosX <= posX && pos.PosX + pos.Width >= posX && pos.PosY <= posY && pos.PosY + pos.Height >= posY
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

func (g *Graphics) SetAllowScroll(value bool, direction structures.ScrollDirection) {
	if direction == structures.Vertical {
		g.verticalScroll.allowed = value
	} else {
		g.horizontalScroll.allowed = value
	}
}

func (g *Graphics) AllowScroll(direction structures.ScrollDirection) bool {
	value := g.verticalScroll.allowed
	if direction == structures.Horizontal {
		value = g.horizontalScroll.allowed
	}
	return value
}

func (g *Graphics) DisableScroll() {
	g.verticalScroll.allowed = false
	g.horizontalScroll.allowed = false
}

func (g *Graphics) SetScrollFromPanel(panel structures.IPanel) {
	g.verticalScroll.allowed = panel.AllowScroll(structures.Vertical)
	g.horizontalScroll.allowed = panel.AllowScroll(structures.Horizontal)
}

func (g Graphics) GetScrollValue(direction structures.ScrollDirection) float32 {
	if direction == structures.Vertical {
		return g.verticalScroll.value
	} else {
		return g.horizontalScroll.value
	}
}

func (g Graphics) GetScrollMaxValue(direction structures.ScrollDirection) float32 {
	if direction == structures.Vertical {
		return g.verticalScroll.maxValue
	} else {
		return g.horizontalScroll.maxValue
	}
}

func (g *Graphics) SetupScroll(panel structures.IPanel, app structures.IApp) {
	g.SetScrollFromPanel(panel)
	if !g.horizontalScroll.allowed && !g.verticalScroll.allowed {
		return
	}

	for _, comp := range panel.GetComponents() {
		pos := comp.GetPosition()
		if !g.ShouldRender(pos) {
			if g.verticalScroll.allowed && pos.PosY + pos.Height > g.GetHeight() && pos.PosY + pos.Height - g.GetHeight() > g.verticalScroll.maxValue { // Vertical
				g.verticalScroll.scrollable = true
				g.verticalScroll.maxValue = pos.PosY + pos.Height - (g.GetHeight() - 10)
				fmt.Println(comp, g.verticalScroll.maxValue)
			}
			if g.horizontalScroll.allowed && pos.PosX + pos.Width > g.GetWidth() && pos.PosX + pos.Width - g.GetWidth() > g.horizontalScroll.maxValue { // Horizontal
				g.horizontalScroll.scrollable = true
				g.horizontalScroll.maxValue = pos.PosX + pos.Width - (g.GetWidth() - 10) //10 is scrollbar width
			}
		}
	}

	if g.horizontalScroll.scrollable {
		g.horizontalScroll.bar = ScrollBar{
			BaseComponent: structures.BaseComponent{
				Func: func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f {
					padding := float32(0)
					if g.verticalScroll.scrollable {
						padding = 10
					}
					return Position(0, graphics.GetHeight() - 10, graphics.GetWidth() - padding, 10)
				},
			},
			Direction: structures.Horizontal,
		}
		g.horizontalScroll.bar.UpdatePosition(g, app)
	}

	if g.verticalScroll.scrollable {
		g.verticalScroll.bar = ScrollBar{
			BaseComponent: structures.BaseComponent{
				Func: func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f {
					return Position(graphics.GetWidth() - 10, 0, 10, graphics.GetHeight())
				},
			},
			Direction: structures.Vertical,
		}
		g.verticalScroll.bar.UpdatePosition(g, app)
	}
}

func (g *Graphics) UpdateScroll(value float32, direction structures.ScrollDirection) {
	scroll := &g.verticalScroll
	if direction == structures.Horizontal {
		scroll = &g.horizontalScroll
	}

	scroll.value += value

	if scroll.value > 0 {
		scroll.value = 0
	}
	if scroll.value < -scroll.maxValue {
		scroll.value = -scroll.maxValue
	}
}

func (g *Graphics) CanScroll() bool {
	return g.horizontalScroll.scrollable || g.verticalScroll.scrollable
}

func (g *Graphics) DrawScrollBar(app structures.IApp, mousePos rl.Vector2) {
	btnDown := rl.IsMouseButtonDown(rl.MouseLeftButton)
	btnReleased := rl.IsMouseButtonReleased(rl.MouseLeftButton)
	btnPressed := rl.IsMouseButtonPressed(rl.MouseLeftButton)

	if g.horizontalScroll.scrollable {
		g.horizontalScroll.bar.Show(g, app)
		if g.horizontalScroll.selected && btnReleased {
			g.horizontalScroll.selected = false
		} else if (btnDown && g.horizontalScroll.selected) || (btnPressed && g.IsInArea(g.horizontalScroll.bar.GetPosition(), mousePos.X, mousePos.Y)) {
			g.horizontalScroll.selected = true
			value := mousePos.X / g.GetWidth() * g.horizontalScroll.maxValue
			g.horizontalScroll.value = float32(math.Min(math.Max(float64(value), 0), float64(g.horizontalScroll.maxValue)))  * -1
		}
	}

	if g.verticalScroll.scrollable {
		g.verticalScroll.bar.Show(g, app)
		if g.verticalScroll.selected && btnReleased {
			g.verticalScroll.selected = false
		} else if (btnDown && g.verticalScroll.selected) || (btnPressed && g.IsInArea(g.verticalScroll.bar.GetPosition(), mousePos.X, mousePos.Y)) {
			g.verticalScroll.selected = true
			value := mousePos.Y / g.GetHeight() * g.verticalScroll.maxValue
			g.verticalScroll.value = float32(math.Min(math.Max(float64(value), 0), float64(g.verticalScroll.maxValue)))  * -1
		}
	}
}

func (g *Graphics) SkipScrollPadding(value bool) {
	g.skipScrollPadding = value
}