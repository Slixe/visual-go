package app

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

type App struct {
	Title        string
	Width        int
	Height       int
	Resizable 	 bool
	DefaultColor rl.Color
	OnSizeChanged func()
	font         *rl.Font
	currentPanel structures.IPanel
	components   []structures.IComponent //globals components
}

func (app App) Start() {
	if app.Resizable {
		rl.SetConfigFlags(rl.FlagWindowResizable)
	}
	rl.InitWindow(app.Width, app.Height, app.Title)
	rl.SetWindowMinSize(app.Width, app.Height)
	rl.SetTargetFPS(60)
	if app.font != nil {
		rl.GuiSetFont(*app.font)
	}
}

func (app *App) Render() {
	var currentSelected structures.ISelectableComponent

	lastWidth := app.GetWidth()
	lastHeight := app.GetHeight()

	for !rl.WindowShouldClose() {
		var selectables []structures.ISelectableComponent
		var components []structures.IComponent

		rl.BeginDrawing()
		rl.ClearBackground(app.DefaultColor)

		if app.OnSizeChanged != nil && (lastHeight != app.GetHeight() || lastWidth != app.GetWidth()) {
			app.OnSizeChanged()
			lastWidth = app.GetWidth()
			lastHeight = app.GetHeight()
		}

		if app.currentPanel != nil {
			app.currentPanel.Show(app)
			components = app.currentPanel.GetComponents()
		}
		components = append(components, app.GetGlobalComponents()...)
		for _, component := range components {
			if selectable, ok := component.(structures.ISelectableComponent); ok {
				selectables = append(selectables, selectable)
			}
			component.Show(app)
		}

		if len(selectables) > 0 && currentSelected == nil {
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
	return rl.GetScreenWidth()
}

func (app App) GetHeight() int {
	return rl.GetScreenHeight()
}

func (app App) GetMinWidth() int {
	return app.Width
}

func (app App) GetMinHeight() int {
	return app.Height
}

func (app *App) ClearGlobalComponents() {
	app.components = []structures.IComponent{}
}

func (app *App) AddGlobalComponent(component structures.IComponent) {
	app.components = append(app.components, component)
}

func (app App) GetGlobalComponents() []structures.IComponent {
	return app.components
}

func (app App) SetWindowTitle(title string) {
	rl.SetWindowTitle(title)
}

func (app App) SetWindowSize(width int, height int) {
	rl.SetWindowSize(width, height)
}

func (app App) SetWindowIcon(imagePath string) {
	img := *rl.LoadImage(imagePath)
	rl.SetWindowIcon(img)
}

func (app *App) SetGuiFont(font *rl.Font) {
	app.font = font
	rl.GuiSetFont(*app.font)
}