package fields

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/graphics"
	"github.com/Slixe/visual-go/structures"
)

type TextField struct {
	structures.BaseComponent
	Editable      bool
	Text          string
	MaxCharacters int
	Callback      func(text TextField)
}

func CreateTextField(text string, editable bool, maxChars int, posX float32, posY float32, width float32, height float32, onTextChanged func(text TextField)) *TextField {
	return &TextField{
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

func (field *TextField) Show(app structures.IApp) {
	_, str := rl.GuiTextBox(graphics.CreateRectangle(field.BaseComponent), field.Text, field.MaxCharacters, field.Editable)
	if field.Text != str {
		field.Callback(*field)
	}
	field.Text = str
}
