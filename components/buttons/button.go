package buttons

import (
	"github.com/Slixe/visual-go/structures"
	rl "github.com/lachee/raylib-goplus/raylib"
)

type Button struct {
	structures.BaseClickable
	Label string
}

func CreateButton(label string, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f, callback func(label structures.IClickable)) *Button {
	return &Button{
		BaseClickable: structures.BaseClickable{
			BaseComponent: structures.BaseComponent{
				Func: posFunc,
			},
			CallbackFunc: callback,
		},
		Label: label,
	}
}

func (btn Button) Show(graphics structures.IGraphics, app structures.IApp) {
	rl.GuiButton(graphics.CreateRectangle(btn.GetPosition()), btn.Label)
}