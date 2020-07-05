package panels

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/graphics"
	"github.com/Slixe/visual-go/structures"
	"github.com/kjk/flex"
)

type TexturedPanel struct {
	BackgroundTexture rl.Texture2D
	Components        []structures.IComponent
	Layout *flex.Node
}

func CreateTexturedPanel(texturePath string) *TexturedPanel {
	return &TexturedPanel{
		BackgroundTexture: graphics.LoadTexture(texturePath),
	}
}

func (panel TexturedPanel) Show(graphics structures.IGraphics, app structures.IApp) {
	graphics.DrawTexturePro(panel.BackgroundTexture, 0, 0, float32(app.GetWidth()), float32(app.GetHeight()), 0, rl.White)
}

func (panel TexturedPanel) GetComponents() []structures.IComponent {
	return panel.Components
}

func (panel *TexturedPanel) AddComponent(component structures.IComponent) {
	panel.Components = append(panel.Components, component)
}

func (panel *TexturedPanel) SetLayout(layout *flex.Node) {
	panel.Layout = layout
}

func (panel TexturedPanel) GetLayout() *flex.Node {
	return panel.Layout
}