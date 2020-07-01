package buttons

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/graphics"
	"github.com/Slixe/visual-go/structures"
)

type LabelButton struct {
	structures.BaseComponent
	Label    string
	Callback func(label LabelButton)
}

func CreateLabelButton(label string, posX float32, posY float32, width float32, height float32, callback func(label LabelButton)) *LabelButton {
	return &LabelButton{
		BaseComponent: structures.BaseComponent{
			PosX:   posX,
			PosY:   posY,
			Width:  width,
			Height: height,
		},
		Label:    label,
		Callback: callback,
	}
}
func (btn LabelButton) Show(app structures.IApp) {
	if rl.GuiLabelButton(graphics.CreateRectangle(btn.BaseComponent), btn.Label) {
		btn.Callback(btn)
	}
}

func (btn LabelButton) GetBaseComponent() structures.BaseComponent {
	return btn.BaseComponent
}