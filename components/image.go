package components

import (
	"github.com/Slixe/visual-go/graphics"
	"github.com/Slixe/visual-go/structures"
	rl "github.com/lachee/raylib-goplus/raylib"
)

type Image struct {
	structures.BaseComponent
	Texture rl.Texture2D
	Color rl.Color
}

func CreateImage(texturePath string, color rl.Color, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f) *Image {
	return &Image{
		Texture: graphics.LoadTexture(texturePath),
		Color: color,
		BaseComponent: structures.BaseComponent{
			Func: posFunc,
		},
	}
}

func (img Image) Show(graphics structures.IGraphics, app structures.IApp) {
	graphics.DrawTexture(img.Texture, img.GetPosition().PosX, img.GetPosition().PosY, img.Color)
}