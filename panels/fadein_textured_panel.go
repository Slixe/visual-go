package panels

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/graphics"
	"github.com/Slixe/visual-go/structures"
	"github.com/kjk/flex"
)

type FadeInTexturedPanel struct {
	BackgroundTexture rl.Texture2D
	Components        []structures.IComponent
	TimeBeforeFade float32
	FadeOutFactor uint8
	FadeColor rl.Color
	Callback func()
	Layout *flex.Node
	duration float32
}

func CreateFadeInTexturedPanel(texturePath string, fadeColor rl.Color, timeBeforeFade float32, fadeOutFactor uint8, onFinished func()) *FadeInTexturedPanel {
	fadeColor.A = 0
	return &FadeInTexturedPanel{
		BackgroundTexture: graphics.LoadTexture(texturePath),
		TimeBeforeFade: timeBeforeFade,
		FadeOutFactor:fadeOutFactor,
		Callback: onFinished,
		FadeColor: fadeColor,
		duration: 0,
	}
}

func (panel *FadeInTexturedPanel) Show(graphics structures.IGraphics, app structures.IApp) {
	panel.duration += rl.GetFrameTime()
	if panel.duration >= panel.TimeBeforeFade {
		if panel.FadeColor.A < 255 {
			panel.FadeColor.A += panel.FadeOutFactor
		} else {
			panel.FadeColor.A = 255
		}
	}

	graphics.DrawTexturePro(panel.BackgroundTexture, 0, 0, float32(app.GetWidth()), float32(app.GetHeight()), 0, panel.FadeColor)

	if panel.FadeColor.A == 255 {
		panel.Callback()
	}
}

func (panel FadeInTexturedPanel) GetComponents() []structures.IComponent {
	return panel.Components
}

func (panel *FadeInTexturedPanel) AddComponent(component structures.IComponent) {
	panel.Components = append(panel.Components, component)
}

func (panel *FadeInTexturedPanel) SetLayout(layout *flex.Node) {
	panel.Layout = layout
}

func (panel FadeInTexturedPanel) GetLayout() *flex.Node {
	return panel.Layout
}
