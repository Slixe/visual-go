package app

import (
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
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		if app.currentPanel != nil {
			app.currentPanel.Show(app)
			for _, component := range app.currentPanel.GetComponents() {
				component.Show(app)
			}
		}
		for _, component := range app.GetGlobalComponents() {
			component.Show(app)
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
