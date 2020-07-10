package buttons

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/graphics"
	"github.com/Slixe/visual-go/structures"
)

type TexturedButton struct {
	structures.BaseComponent
	Label    string
	Texture  rl.Texture2D
	Callback func(btn TexturedButton)
}

func CreateTexturedButton(label string, texturePath string, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.ComponentPos, callback func(btn TexturedButton)) *TexturedButton {
	return &TexturedButton{
		BaseComponent: structures.BaseComponent{
			Func: posFunc,
		},
		Label:    label,
		Texture:  graphics.LoadTexture(texturePath),
		Callback: callback,
	}
}

func (btn TexturedButton) Show(graphics structures.IGraphics, app structures.IApp) {
	if rl.GuiImageButton(graphics.CreateRectangle(btn.GetPosition()), btn.Label, btn.Texture) {
		btn.Callback(btn)
	}
}