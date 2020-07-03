package panels

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/graphics"
	"github.com/Slixe/visual-go/structures"
)

type FadeOutTexturedPanel struct {
	BackgroundTexture rl.Texture2D
	Components        []structures.IComponent
	TimeBeforeFade float32
	FadeOutFactor uint8
	FadeColor rl.Color
	Callback func()
	duration float32
}

func CreateFadeOutTexturedPanel(texturePath string, fadeColor rl.Color, timeBeforeFade float32, fadeOutFactor uint8, onFinished func()) *FadeOutTexturedPanel {
	return &FadeOutTexturedPanel{
		BackgroundTexture: rl.LoadTexture(texturePath),
		TimeBeforeFade: timeBeforeFade,
		FadeOutFactor:fadeOutFactor,
		Callback: onFinished,
		FadeColor: fadeColor,
		duration: 0,
	}
}

func (panel *FadeOutTexturedPanel) Show(app structures.IApp) {
	panel.duration += rl.GetFrameTime()
	if panel.duration >= panel.TimeBeforeFade {
		if panel.FadeColor.A > panel.FadeOutFactor {
			panel.FadeColor.A -= panel.FadeOutFactor
		} else {
			panel.FadeColor.A = 0
		}
	}
	graphics.DrawTexturePro(panel.BackgroundTexture, 0, 0, float32(app.GetWidth()), float32(app.GetHeight()), 0, panel.FadeColor)

	if panel.FadeColor.A == 0 {
		panel.Callback()
	}
}

func (panel FadeOutTexturedPanel) GetComponents() []structures.IComponent {
	return panel.Components
}

func (panel *FadeOutTexturedPanel) AddComponent(component structures.IComponent) {
	panel.Components = append(panel.Components, component)
}
