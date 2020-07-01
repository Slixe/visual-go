package fields

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/graphics"
	"github.com/Slixe/visual-go/structures"
	"strings"
)

type PasswordField struct {
	structures.BaseComponent
	Editable      bool
	Text          string
	MaxCharacters int
	Callback      func(password PasswordField)
}

func CreatePasswordField(text string, editable bool, maxChars int, posX float32, posY float32, width float32, height float32, onTextChanged func(password PasswordField)) *PasswordField {
	return &PasswordField{
		BaseComponent: structures.BaseComponent{
			PosX:   posX,
			PosY:   posY,
			Width:  width,
			Height: height,
		},
		Text:          text,
		Editable:      editable,
		MaxCharacters: maxChars,
		Callback:      onTextChanged,
	}
}

func (field *PasswordField) Show(app structures.IApp) {
	_, str := rl.GuiTextBox(graphics.CreateRectangle(field.BaseComponent), strings.Repeat("*", len(field.Text)), field.MaxCharacters, field.Editable)

	if len(field.Text) != len(str) {
		if len(str) < len(field.Text) && len(field.Text) > 0 {
			field.Text = field.Text[:len(field.Text)-1]
		} else {
			field.Text += str[len(field.Text):]
		}

		field.Callback(*field)
	}

}
