package buttons

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type Button struct {
	structures.BaseClickable
}

func CreateButton(label string, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.ComponentPos, callback func(label structures.IClickable)) *Button {
	return &Button{
		BaseClickable: structures.BaseClickable{
			BaseComponent: structures.BaseComponent{
				Func: posFunc,
			},
			Label:    label,
			Callback: callback,
		},
	}
}

func (btn Button) Show(graphics structures.IGraphics, app structures.IApp) {
	if rl.GuiButton(graphics.CreateRectangle(btn.GetPosition()), btn.Label) {
		btn.Callback(&btn)
	}
}