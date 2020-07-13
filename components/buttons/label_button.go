package buttons

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type LabelButton struct {
	structures.BaseClickable
}

func CreateLabelButton(label string, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.ComponentPos, callback func(label structures.IClickable)) *LabelButton {
	return &LabelButton{
		BaseClickable: structures.BaseClickable{
			BaseComponent: structures.BaseComponent{
				Func: posFunc,
			},
			Label:    label,
			Callback: callback,
		},
	}
}

func (btn LabelButton) Show(graphics structures.IGraphics, app structures.IApp) {
	if rl.GuiLabelButton(graphics.CreateRectangle(btn.GetPosition()), btn.Label) {
		btn.Callback(&btn)
	}
}