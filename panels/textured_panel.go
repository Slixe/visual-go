package panels

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/graphics"
	"github.com/Slixe/visual-go/structures"
)

type TexturedPanel struct {
	BackgroundTexture rl.Texture2D
	Components        []structures.IComponent
}

func CreateTexturedPanel(texturePath string) *TexturedPanel {
	return &TexturedPanel{
		BackgroundTexture: rl.LoadTexture(texturePath),
	}
}

func (panel TexturedPanel) Show(app structures.IApp) {
	graphics.DrawTexture(panel.BackgroundTexture, 0, 0, float32(app.GetWidth()), float32(app.GetHeight()), 0)
}

func (panel TexturedPanel) GetComponents() []structures.IComponent {
	return panel.Components
}

func (panel *TexturedPanel) AddComponent(component structures.IComponent) {
	panel.Components = append(panel.Components, component)
}
