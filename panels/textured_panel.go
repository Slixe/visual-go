package panels

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/graphics"
	"github.com/Slixe/visual-go/structures"
)

type TexturedPanel struct {
	BackgroundTexture rl.Texture2D
	structures.BasePanel
}

func CreateTexturedPanel(texturePath string) *TexturedPanel {
	return &TexturedPanel{
		BackgroundTexture: graphics.LoadTexture(texturePath),
	}
}

func (panel TexturedPanel) Show(graphics structures.IGraphics, app structures.IApp) {
	graphics.DrawTexturePro(panel.BackgroundTexture, 0, 0, float32(app.GetWidth()), float32(app.GetHeight()), 0, rl.White)
}