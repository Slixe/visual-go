package fields

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type ValueField struct {
	structures.BaseInputField
	Value          int
	MinValue int
	MaxValue int
}

func CreateValueField(text string, value int, minValue int, maxValue int, editable bool, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f, onTextChanged func(text structures.IInputField)) *ValueField {
	return &ValueField{
		BaseInputField: structures.BaseInputField{
			BaseSelectableComponent: structures.BaseSelectableComponent{
				BaseComponent: structures.BaseComponent{
					Func: posFunc,
				},
				Selected:      false,
			},
			Text: text,
			Editable: editable,
			Callback: onTextChanged,
		},
		Value:	value,
		MinValue: minValue,
		MaxValue: maxValue,
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
		field.Callback(field)
	}
	field.Value = val
}