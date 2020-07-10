package fields

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type ValueField struct {
	structures.BaseSelectableComponent
	Editable      bool
	Text 		string
	Value          int
	MinValue int
	MaxValue int
	Callback      func(text ValueField)
}

func CreateValueField(text string, value int, minValue int, maxValue int, editable bool, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.ComponentPos, onTextChanged func(text ValueField)) *ValueField {
	return &ValueField{
		BaseSelectableComponent: structures.BaseSelectableComponent{
			BaseComponent: structures.BaseComponent{
				Func: posFunc,
			},
			Selected:      false,
		},
		Text: text,
		Value:	value,
		MinValue: minValue,
		MaxValue: maxValue,
		Editable: editable,
		Callback: onTextChanged,
	}
}

func (field *ValueField) Show(graphics structures.IGraphics, app structures.IApp) {
	editable := field.Editable
	if editable {
		editable = field.Selected
	}
	_, val := rl.GuiValueBox(graphics.CreateRectangle(field.GetPosition()), field.Text, field.Value, field.MinValue, field.MaxValue, editable) //min & max not working

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