package fields

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type TextField struct {
	structures.BaseInputField
}

func CreateTextField(text string, editable bool, maxChars int, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f, onTextChanged func(text structures.IInputField)) *TextField {
	return &TextField{
		BaseInputField: structures.BaseInputField{
			BaseSelectableComponent: structures.BaseSelectableComponent{
				BaseComponent: structures.BaseComponent{
					Func: posFunc,
				},
				Selected: false,
			},
			Text:          text,
			Editable:      editable,
			MaxCharacters: maxChars,
			Callback:      onTextChanged,
		},
	}
}

func (field *TextField) Show(graphics structures.IGraphics, app structures.IApp) {
	editable := field.Editable
	if editable {
		editable = field.Selected
	}

	_, str := rl.GuiTextBox(graphics.CreateRectangle(field.GetPosition()), field.Text, field.MaxCharacters, editable)

	if field.Text != str {
		field.Callback(field)
	}
	field.Text = str
}