package fields

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type ValueField struct {
	structures.BaseComponent
	Editable      bool
	Text 		string
	Value          int
	MinValue int
	MaxValue int
	Callback      func(text ValueField)
	Selected bool
}

func CreateValueField(text string, value int, minValue int, maxValue int, editable bool, posX float32, posY float32, width float32, height float32, onTextChanged func(text ValueField)) *ValueField {
	return &ValueField{
		BaseComponent: structures.BaseComponent{
			PosX:   posX,
			PosY:   posY,
			Width:  width,
			Height: height,
		},
		Text: text,
		Value:          value,
		MinValue: minValue,
		MaxValue: maxValue,
		Editable:      editable,
		Callback:      onTextChanged,
		Selected: false,
	}
}

func (field *ValueField) Show(graphics structures.IGraphics, app structures.IApp) {
	editable := field.Editable
	if editable {
		editable = field.Selected
	}
	_, val := rl.GuiValueBox(graphics.CreateRectangle(field.BaseComponent), field.Text, field.Value, field.MinValue, field.MaxValue, editable) //min & max not working

	if val > field.MaxValue {
		val = field.MaxValue
	} else if val < field.MinValue {
		val = field.MinValue
	}

	if field.Value != val {
		field.Callback(*field)
	}
	field.Value = val
}

func (field ValueField) GetBaseComponent() structures.BaseComponent {
	return field.BaseComponent
}

func (field ValueField) IsSelected() bool {
	return field.Selected
}

func (field *ValueField) SetSelected(value bool) {
	field.Selected = value
}