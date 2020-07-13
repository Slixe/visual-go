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

type panelInfo struct {
	layoutName string
	graphics   structures.IGraphics
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
	layouts 	 map[string]*flex.Node
	panels 		 map[structures.IPanel]panelInfo
	components   []structures.IComponent //globals components
}

func (app *App) Start() {
	app.layouts = make(map[string]*flex.Node)
	app.panels = make(map[structures.IPanel]panelInfo)

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
		var selectables = make(map[structures.IGraphics][]structures.ISelectableComponent)
		var clickables = make(map[structures.IGraphics][]structures.IClickable)
		var components = make(map[structures.IGraphics][]structures.IComponent)

		rl.BeginDrawing()
		rl.ClearBackground(app.DefaultColor)

		if lastHeight != app.GetHeight() || lastWidth != app.GetWidth() {
			lastWidth = app.GetWidth()
			lastHeight = app.GetHeight()

			app.CalculateLayout()
			globalLayout = graphics.CreateGraphics(*app.Flex.RootNode)
		}

		if app.globalPanel != nil {
			app.globalPanel.Show(globalLayout, app)
			components[globalLayout] = append(components[globalLayout], app.globalPanel.GetComponents()...)
		}

		for panel, g := range app.panels {
			panel.Show(g.graphics, app)
			components[g.graphics] = append(components[g.graphics], panel.GetComponents()...)
		}

		globalComponents := app.GetGlobalComponents()
		components[globalLayout] = append(components[globalLayout], globalComponents...)
		for graph, comps := range components {
			for _, component := range comps {
				if selectable, ok := component.(structures.ISelectableComponent); ok {
					selectables[graph] = append(selectables[graph], selectable)
					if currentSelected == nil && selectable.IsSelected() {
						currentSelected = selectable
					}
				} else if clickable, ok := component.(structures.IClickable); ok {
					clickables[graph] = append(clickables[graph], clickable)
				}

				component.Show(graph, app)
			}
		}

		if len(selectables) > 0 && currentSelected == nil {
			for _, values := range selectables {
				for _, selectable := range values {
					currentSelected = selectable
					currentSelected.SetSelected(true)
					break
				}
				break
			}
		}

		mousePos := rl.GetMousePosition()
		for graph, values := range selectables {
			for _, selectable := range values {
				pos := selectable.GetPosition()
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) && graph.GetPosX() + pos.PosX <= mousePos.X && graph.GetPosX() + pos.PosX + pos.Width >= mousePos.X && graph.GetPosY() + pos.PosY <= mousePos.Y && graph.GetPosY() + pos.PosY + pos.Height >= mousePos.Y {
					currentSelected.SetSelected(false)
					selectable.SetSelected(true)
					currentSelected = selectable
				}
			}
		}

		for graph, values := range clickables {
			for _, clickable := range values {
				pos := clickable.GetPosition()
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) && graph.GetPosX() + pos.PosX <= mousePos.X && graph.GetPosX() + pos.PosX + pos.Width >= mousePos.X && graph.GetPosY() + pos.PosY <= mousePos.Y && graph.GetPosY() + pos.PosY + pos.Height >= mousePos.Y {
					clickable.OnClicked()
				}
			}
		}
		rl.EndDrawing()
	}

	app.Close()
}

func (app App) Close() {
	rl.CloseWindow()
}

func (app *App) CalculateLayout() {
	if app.Flex.RootNode != nil {
		flex.CalculateLayout(app.Flex.RootNode, float32(app.GetWidth()), float32(app.GetHeight()), app.Flex.Direction)
	}

	for panel, info := range app.panels {
		graph := graphics.CreateGraphics(*app.layouts[info.layoutName])
		info.graphics = graph
		app.panels[panel] = info

		for _, comp := range panel.GetComponents() {
			comp.UpdatePosition(graph, app)
		}
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

func (app *App) SetPanel(layoutName string, panel structures.IPanel) bool {

	if _, ok := app.layouts[layoutName]; !ok {
		return false
	}

	app.panels[panel] = panelInfo{
		layoutName: layoutName,
		graphics:   graphics.CreateGraphics(*app.layouts[layoutName]),
	}

	app.CalculateLayout()

	return true
}

func (app *App) RemovePanel(panel structures.IPanel) {
	_, ok := app.panels[panel]
	if ok {
		delete(app.panels, panel)
	}
}

func (app App) GetPanels() []structures.IPanel {
	panels := make([]structures.IPanel, len(app.panels))
	i := 0
	for p := range app.panels {
		panels[i] = p
		i++
	}
	return panels
}

func (app *App) ClearPanels() {
	app.panels = make(map[structures.IPanel]panelInfo)
}

func (app *App) RegisterLayout(layoutName string, node *flex.Node) {
	app.layouts[layoutName] = node
}

func (app *App) RemoveLayout(layoutName string) *flex.Node {
	v, ok := app.layouts[layoutName]
	if ok {
		delete(app.layouts, layoutName)
	}

	return v
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