package app

import (
	"fmt"
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type App struct {
	currentPanel structures.IPanel
	Title        string
	Width        int
	Height       int
	Components   []structures.IComponent //globals components
}

func (app App) Start() {
	rl.InitWindow(app.Width, app.Height, app.Title)
	rl.SetTargetFPS(60)
}

func (app *App) Render() {
	var currentSelected structures.ISelectableComponent

	for !rl.WindowShouldClose() {
		var selectables []structures.ISelectableComponent
		var components []structures.IComponent
		hasTextFieldSelected := false

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		if app.currentPanel != nil {
			app.currentPanel.Show(app)
			components = app.currentPanel.GetComponents()
		}
		components = append(components, app.GetGlobalComponents()...)
		for _, component := range components {
			if selectable, ok := component.(structures.ISelectableComponent); ok {
				if selectable.IsSelected() {
					if hasTextFieldSelected {
						selectable.SetSelected(false)
					}

					hasTextFieldSelected = true
				}
				selectables = append(selectables, selectable)
			}
			component.Show(app)
		}

		if len(selectables) > 0 && !hasTextFieldSelected {
			selectables[0].SetSelected(true) //we select first by default
			currentSelected = selectables[0]
		}

		mousePos := rl.GetMousePosition()
		for _, selectable := range selectables {
			base := selectable.GetBaseComponent()
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && base.PosX <= mousePos.X && base.PosX + base.Width >= mousePos.X && base.PosY <= mousePos.Y && base.PosY + base.Height >= mousePos.Y {
				currentSelected.SetSelected(false)
				selectable.SetSelected(true)
				currentSelected = selectable
				fmt.Println(selectable.IsSelected())
			}
		}
		rl.EndDrawing()
	}

	app.Close()
}

func (app App) Close() {
	rl.CloseWindow()
}
func (app App) GetCurrentPanel() structures.IPanel {
	return app.currentPanel
}

func (app *App) SetPanel(panel structures.IPanel) {
	app.currentPanel = panel
}

func (app App) GetWidth() int {
	return app.Width
}

func (app App) GetHeight() int {
	return app.Height
}

func (app *App) ClearGlobalComponents() {
	app.Components = []structures.IComponent{}
}

func (app *App) AddGlobalComponent(component structures.IComponent) {
	app.Components = append(app.Components, component)
}

func (app App) GetGlobalComponents() []structures.IComponent {
	return app.Components
}
