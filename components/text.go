package components

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type Text struct {
	structures.BaseComponent
	Text string
	Font rl.Font
	FontSize float32
	Spacing float32
	Color rl.Color
}

func CreateText(font rl.Font, text string, color rl.Color, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f, fontSize float32, spacing float32) *Text {
	return &Text{
		Text: text,
		Font: font,
		Color: color,
		FontSize: fontSize,
		Spacing: spacing,
		BaseComponent: structures.BaseComponent{
			Func: posFunc,
		},
	}
}

func (text Text) Show(graphics structures.IGraphics, app structures.IApp) {
	graphics.DrawText(text.Font, text.Text, text.GetPosition().PosX, text.GetPosition().PosY, text.FontSize, text.Spacing, text.Color)
}