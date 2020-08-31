package app

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/graphics"
	"github.com/Slixe/visual-go/structures"
	"github.com/kjk/flex"
	"os"
	"os/signal"
	"syscall"
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
	TargetFPS	 int
	DefaultColor rl.Color
	Flex         Flex
	font         *rl.Font
	layouts 	 map[string]*flex.Node
	panels 		 map[structures.IPanel]panelInfo
	shouldClose bool
}

func (app *App) Start() {
	app.layouts = make(map[string]*flex.Node)
	app.panels = make(map[structures.IPanel]panelInfo)

	if app.Resizable {
		rl.SetConfigFlags(rl.FlagWindowResizable)
	}

	if app.TargetFPS == 0 {
		app.TargetFPS = 60
	}

	if app.Flex.FlexConfig == nil {
		app.Flex.FlexConfig = flex.NewConfig()
	}

	rl.InitWindow(app.Width, app.Height, app.Title)
	rl.SetWindowMinSize(app.Width, app.Height)
	rl.SetTargetFPS(app.TargetFPS)
	rl.SetExitKey(0)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sigchan
		app.shouldClose = true
	}()
}

func (app *App) Render() {
	var currentSelected structures.ISelectable
	lastWidth := app.GetWidth()
	lastHeight := app.GetHeight()

	app.CalculateLayout()

	for !rl.WindowShouldClose() && !app.shouldClose {
		selectables := make(map[structures.IGraphics][]structures.ISelectable)
		clickables := make(map[structures.IGraphics][]structures.IClickable)
		wheelMove := rl.GetMouseWheelMove()
		mousePos := rl.GetMousePosition()

		if lastHeight != app.GetHeight() || lastWidth != app.GetWidth() {
			lastWidth = app.GetWidth()
			lastHeight = app.GetHeight()

			app.CalculateLayout()
		}

		rl.ClearBackground(app.DefaultColor)
		rl.BeginDrawing()

		for panel, g := range app.panels {
			rl.BeginScissorMode(int(g.graphics.GetPosX()), int(g.graphics.GetPosY()), int(g.graphics.GetWidth()), int(g.graphics.GetHeight())) //prevent drawing on each other
			temp := g.graphics.AllowScroll()
			g.graphics.ShouldAllowScroll(false)
			panel.Show(g.graphics, app)
			g.graphics.ShouldAllowScroll(temp)

			for _, component := range panel.GetComponents() {
				if selectable, ok := component.(structures.ISelectable); ok {
					selectables[g.graphics] = append(selectables[g.graphics], selectable)
					if currentSelected == nil && selectable.IsSelected() {
						currentSelected = selectable
					}
				} else if clickable, ok := component.(structures.IClickable); ok {
					clickables[g.graphics] = append(clickables[g.graphics], clickable)
				}

				component.Show(g.graphics, app)
			}

			if g.graphics.AllowScroll() && wheelMove != 0 && g.graphics.IsInArea2(mousePos.X, mousePos.Y) {
				switch wheelMove {
				case -1:
					g.graphics.UpdateScroll(-5)
				case 1:
					g.graphics.UpdateScroll(5)
				}
				wheelMove = 0 //we only update one scrollable panel
			}

			rl.EndScissorMode()
		}

		if currentSelected == nil && len(selectables) > 0 {
			for _, values := range selectables {
				for _, selectable := range values {
					currentSelected = selectable
					currentSelected.SetSelected(true)
					break
				}
				break
			}
		}

		tabKeyPressed := rl.IsKeyPressed(rl.KeyTab)
		currentSelectedPassed := false
		for graph, values := range selectables {
			for i, selectable := range values {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) && graph.IsInArea(selectable.GetPosition(), mousePos.X, mousePos.Y) {
					currentSelected.SetSelected(false)
					selectable.SetSelected(true)
					currentSelected = selectable
					break
				} else if tabKeyPressed && (currentSelectedPassed || len(values) == i + 1) {
					currentSelected.SetSelected(false)
					if currentSelectedPassed {
						selectable.SetSelected(true)
						currentSelected = selectable
					} else if len(values) == i + 1 {
						values[0].SetSelected(true)
						currentSelected = values[0]
					}
					tabKeyPressed = false
					break
				}

				if selectable == currentSelected {
					currentSelectedPassed = true
				}
			}
		}

		for graph, values := range clickables {
			for _, clickable := range values {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) && graph.IsInArea(clickable.GetPosition(), mousePos.X, mousePos.Y) {
					clickable.Callback()
					break //we can only click on one button at a time
				}
			}
			break
		}
		rl.EndDrawing()
	}

	app.Close()
}

func (app App) Close() {
	rl.CloseWindow()
}

func (app *App) CalculateLayout() {
	if app.Flex.RootNode == nil {
		app.Flex.RootNode = flex.NewNodeWithConfig(app.Flex.FlexConfig)
	}

	flex.CalculateLayout(app.Flex.RootNode, float32(app.GetWidth()), float32(app.GetHeight()), app.Flex.Direction)

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

func (app *App) SetPanel(layoutName string, panel structures.IPanel) {

	if _, ok := app.layouts[layoutName]; !ok {
		panic("Layout '" + layoutName + "' doesn't exist!")
	}

	app.panels[panel] = panelInfo{
		layoutName: layoutName,
		graphics:   graphics.CreateGraphics(*app.layouts[layoutName]),
	}

	app.CalculateLayout()
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
	if app.Flex.RootNode == nil {
		app.Flex.RootNode = flex.NewNodeWithConfig(app.Flex.FlexConfig)
	}

	app.RegisterLayoutChild(layoutName, node, app.Flex.RootNode)
}

func (app *App) RegisterLayoutChild(layoutName string, node *flex.Node, parent *flex.Node) {
	parent.InsertChild(node, len(parent.Children))
	app.layouts[layoutName] = node
}

func (app *App) RemoveLayout(layoutName string) *flex.Node {
	v, ok := app.layouts[layoutName]
	if ok {
		delete(app.layouts, layoutName)
	}

	return v
}

func (app App) NewLayout() *flex.Node {
	return flex.NewNodeWithConfig(app.Flex.FlexConfig)
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

func (app *App) SetTargetFPS(targetFPS int) {
	app.TargetFPS = targetFPS
	rl.SetTargetFPS(app.TargetFPS)
}

func (app App) GetFPS() int {
	return rl.GetFPS()
}