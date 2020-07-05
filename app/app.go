package app

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/graphics"
	"github.com/Slixe/visual-go/structures"
	"github.com/kjk/flex"
)

type Flex struct {
	FlexConfig *flex.Config
	RootNode *flex.Node
	Direction flex.Direction
}

type App struct {
	Title        string
	Width        int
	Height       int
	Resizable    bool
	DefaultColor rl.Color
	Flex         Flex
	font         *rl.Font
	globalPanel  structures.IPanel
	panels 		 map[string]structures.IPanel
	components   []structures.IComponent //globals components
}

func (app *App) Start() {
	app.panels = make(map[string]structures.IPanel)

	if app.Resizable {
		rl.SetConfigFlags(rl.FlagWindowResizable)
	}

	rl.InitWindow(app.Width, app.Height, app.Title)
	rl.SetWindowMinSize(app.Width, app.Height)
	rl.SetTargetFPS(60)
}

func (app *App) Render() {
	var currentSelected structures.ISelectableComponent
	lastWidth := app.GetWidth()
	lastHeight := app.GetHeight()

	app.CalculateLayout()
	globalLayout := graphics.CreateGraphics(*app.Flex.RootNode)
	for !rl.WindowShouldClose() {
		var selectables []structures.ISelectableComponent
		var components []structures.IComponent
		g := globalLayout

		rl.BeginDrawing()
		rl.ClearBackground(app.DefaultColor)

		if lastHeight != app.GetHeight() || lastWidth != app.GetWidth() {
			lastWidth = app.GetWidth()
			lastHeight = app.GetHeight()

			app.CalculateLayout()
		}

		for _, panel := range app.panels {
			if panel.GetLayout() != nil {
				panel.Show(graphics.CreateGraphics(*panel.GetLayout()), app)
			} else {
				panel.Show(globalLayout, app)
			}

			components = append(components, panel.GetComponents()...)
		}

		if app.globalPanel != nil {
			layout := app.globalPanel.GetLayout()
			if layout != nil {
				g = graphics.CreateGraphics(*layout)
			}
			app.globalPanel.Show(g, app)
			components = append(components, app.globalPanel.GetComponents()...)
		}

		globalComponents := app.GetGlobalComponents()
		components = append(components, globalComponents...)
		for i, component := range components {
			if selectable, ok := component.(structures.ISelectableComponent); ok {
				selectables = append(selectables, selectable)
			}

			if i > len(components) - len(globalComponents) {
				component.Show(globalLayout, app)
			} else {
				component.Show(g, app)
			}
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

func (app App) CalculateLayout() {
	if app.Flex.RootNode != nil {
		flex.CalculateLayout(app.Flex.RootNode, float32(app.GetWidth()), float32(app.GetHeight()), app.Flex.Direction)
	}
}

func (app *App) SetMainLayout(node *flex.Node) {
	app.Flex.RootNode = node
}

func (app App) GetGlobalPanel() structures.IPanel {
	return app.globalPanel
}

func (app *App) SetGlobalPanel(panel structures.IPanel) {
	app.globalPanel = panel
	app.CalculateLayout()
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

func (app *App) SetPanel(name string, panel structures.IPanel) {
	app.panels[name] = panel
}

func (app *App) RemovePanel(name string) {
	_, ok := app.panels[name]
	if ok {
		delete(app.panels, name)
	}
}

func (app App) GetPanel(name string) structures.IPanel {
	return app.panels[name]
}

func (app App) GetPanels() []structures.IPanel {
	panels := make([]structures.IPanel, len(app.panels))
	i := 0
	for _, v := range app.panels {
		panels[i] = v
		i++
	}
	return panels
}

func (app *App) ClearPanels() {
	app.panels = make(map[string]structures.IPanel)
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