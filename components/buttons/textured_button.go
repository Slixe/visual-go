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

func CreateTexturedButton(label string, texturePath string, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f, callback func(btn structures.IClickable)) *TexturedButton {
	return &TexturedButton{
		BaseClickable: structures.BaseClickable{
			BaseComponent: structures.BaseComponent{
				Func: posFunc,
			},
			Label:        label,
			CallbackFunc: callback,
		},
		Texture:  graphics.LoadTexture(texturePath),
	}
}

func (btn TexturedButton) Show(graphics structures.IGraphics, app structures.IApp) {
	rl.GuiImageButton(graphics.CreateRectangle(btn.GetPosition()), btn.Label, btn.Texture)
}