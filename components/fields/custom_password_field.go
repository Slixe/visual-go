package fields

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/CMD/globals"
	"github.com/Slixe/visual-go/structures"
	"strings"
)

type CustomPasswordField struct {
	structures.BaseInputField
	BackgroundColor rl.Color
	BorderColor rl.Color
	timer float32
	draw bool
}

func CreateCustomPasswordField(text string, editable bool, maxChars int, posFunc func(graphics structures.IGraphics, app structures.IApp) structures.ComponentPos, backgroundColor rl.Color, borderColor rl.Color, onTextChanged func(text structures.IInputField)) *CustomPasswordField {
	return &CustomPasswordField{
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

func (field *CustomPasswordField) Show(graphics structures.IGraphics, app structures.IApp) {
	field.timer += rl.GetFrameTime()
	if field.timer >= 0.5 {
		field.timer = 0
		field.draw = !field.draw
	}

	field.HandleKey()
	mask := strings.Repeat("*", len(field.Text))
	graphics.DrawRectangle(int(field.GetPosition().PosX), int(field.GetPosition().PosY), int(field.GetPosition().Width), int(field.GetPosition().Height), field.BackgroundColor)
	y := int(field.GetPosition().PosY + field.GetPosition().Height)
	graphics.DrawRectangle(int(field.GetPosition().PosX), y, int(field.GetPosition().Width), 3, globals.GlobalColor)
	textSize := graphics.MeasureText(globals.Font, mask, 24, 0)

	for len(mask) > 0 && textSize.X > field.GetPosition().Width {
		mask = mask[1:]
		textSize = graphics.MeasureText(globals.Font, mask, 24, 0)
	}

	if field.IsSelected() && field.draw {
		graphics.DrawLine(int(field.GetPosition().PosX + textSize.X), y - int(field.GetPosition().Height), int(field.GetPosition().PosX + textSize.X) + 1, y, field.BorderColor)
	}

	graphics.DrawText(globals.Font, mask, field.GetPosition().PosX, field.GetPosition().PosY + field.GetPosition().Height / 4, 24, 0, rl.White)
}