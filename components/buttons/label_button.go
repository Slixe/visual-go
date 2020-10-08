package buttons

import (
	"github.com/Slixe/visual-go/structures"
	rl "github.com/lachee/raylib-goplus/raylib"
)

type LabelButton struct {
	Button
}

func CreateLabelButton(label string, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f, callback func(label structures.IClickable)) *LabelButton {
	return &LabelButton{
		Button: Button{
			BaseClickable: structures.BaseClickable{
				BaseComponent: structures.BaseComponent{
					Func: posFunc,
				},
				CallbackFunc: callback,
			},
			Label: label,
		},
	}
}

func (btn LabelButton) Show(graphics structures.IGraphics, app structures.IApp) {
	rl.GuiLabelButton(graphics.CreateRectangle(btn.GetPosition()), btn.Label)
}