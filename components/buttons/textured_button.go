package buttons

import (
	"github.com/Slixe/visual-go/graphics"
	"github.com/Slixe/visual-go/structures"
	rl "github.com/lachee/raylib-goplus/raylib"
)

type TexturedButton struct {
	Button
	Texture  rl.Texture2D
}

func CreateTexturedButton(label string, texturePath string, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f, callback func(btn structures.IClickable)) *TexturedButton {
	return &TexturedButton{
		Button: Button{
			BaseClickable: structures.BaseClickable{
				BaseComponent: structures.BaseComponent{
					Func: posFunc,
				},
				CallbackFunc: callback,
			},
			Label: label,
		},
		Texture:  graphics.LoadTexture(texturePath),
	}
}

func (btn TexturedButton) Show(graphics structures.IGraphics, app structures.IApp) {
	rl.GuiImageButton(graphics.CreateRectangle(btn.GetPosition()), btn.Label, btn.Texture)
}