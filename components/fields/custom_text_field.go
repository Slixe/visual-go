package fields

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/CMD/globals"
	"github.com/Slixe/visual-go/structures"
)

type CustomTextField struct {
	structures.BaseInputField
	BackgroundColor rl.Color
	BorderColor rl.Color
	timer float32
	draw bool
}

func CreateCustomTextField(text string, editable bool, maxChars int, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.ComponentPos, backgroundColor rl.Color, borderColor rl.Color, onTextChanged func(text structures.IInputField)) *CustomTextField {
	return &CustomTextField{
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
			Callback: onTextChanged,
		},
		BackgroundColor: backgroundColor,
		BorderColor: borderColor,
		timer: 0,
		draw: true,
	}
}

func (field *CustomTextField) Show(graphics structures.IGraphics, app structures.IApp) {
	field.timer += rl.GetFrameTime()
	if field.timer >= 0.5 {
		field.timer = 0
		field.draw = !field.draw
	}

	field.HandleKey()
	graphics.DrawRectangle(int(field.GetPosition().PosX), int(field.GetPosition().PosY) - 10, int(field.GetPosition().Width), 50, field.BackgroundColor)
	y := int(field.GetPosition().PosY + field.GetPosition().Height)
	graphics.DrawRectangle(int(field.GetPosition().PosX), y, int(field.GetPosition().Width), 3, globals.GlobalColor)
	textSize := rl.MeasureTextEx(globals.Font, field.Text, 24, 0)
	if field.IsSelected() && field.draw {
		graphics.DrawLine(int(field.GetPosition().PosX + textSize.X), y - int(field.GetPosition().Height), int(field.GetPosition().PosX + textSize.X) + 1, y, field.BorderColor)
	}

	graphics.DrawText(globals.Font, field.Text, field.GetPosition().PosX, field.GetPosition().PosY, 24, 0, rl.White) //TODO, if textSize > Width then draw only the end of Text
}