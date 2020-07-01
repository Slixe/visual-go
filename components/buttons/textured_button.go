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

func LoadAndCreateTexturedButton(label string, texturePath string, posX float32, posY float32, width float32, height float32, callback func(btn TexturedButton)) *TexturedButton {
	return CreateTexturedButton(label, rl.LoadTexture(texturePath), posX, posY, width, height, callback)
}

func CreateTexturedButton(label string, texture rl.Texture2D, posX float32, posY float32, width float32, height float32, callback func(btn TexturedButton)) *TexturedButton {
	return &TexturedButton{
		BaseComponent: structures.BaseComponent{
			PosX:   posX,
			PosY:   posY,
			Width:  width,
			Height: height,
		},
		Label:    label,
		Texture:  texture,
		Callback: callback,
	}
}

func (btn TexturedButton) Show(app structures.IApp) {
	if rl.GuiImageButton(graphics.CreateRectangle(btn.BaseComponent), btn.Label, btn.Texture) {
		btn.Callback(btn)
	}
}
