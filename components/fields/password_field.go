package fields

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
	"strings"
)

type PasswordField struct {
	structures.BaseInputField
}

func CreatePasswordField(text string, editable bool, maxChars int, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f, onTextChanged func(password structures.IInputField)) *PasswordField {
	return &PasswordField{
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

func (field *PasswordField) Show(graphics structures.IGraphics, app structures.IApp) {
	editable := field.Editable
	if editable {
		editable = field.Selected
	}

	_, str := rl.GuiTextBox(graphics.CreateRectangle(field.GetPosition()), strings.Repeat("*", len(field.Text)), field.MaxCharacters, editable)

	if len(field.Text) != len(str) {
		if len(str) < len(field.Text) && len(field.Text) > 0 {
			field.Text = field.Text[:len(field.Text)-1]
		} else {
			field.Text += str[len(field.Text):]
		}

		field.Callback(field)
	}
}