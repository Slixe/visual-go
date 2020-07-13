package buttons

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/graphics"
	"github.com/Slixe/visual-go/structures"
)

type TexturedButton struct {
	structures.BaseClickable
	Texture  rl.Texture2D
}

func CreateTexturedButton(label string, texturePath string, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.ComponentPos, callback func(btn structures.IClickable)) *TexturedButton {
	return &TexturedButton{
		BaseClickable: structures.BaseClickable{
			BaseComponent: structures.BaseComponent{
				Func: posFunc,
			},
			Label:    label,
			Callback: callback,
		},
		Texture:  graphics.LoadTexture(texturePath),
	}
}

func (btn TexturedButton) Show(graphics structures.IGraphics, app structures.IApp) {
	if rl.GuiImageButton(graphics.CreateRectangle(btn.GetPosition()), btn.Label, btn.Texture) {
		btn.Callback(&btn)
	}
}