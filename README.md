# Visual GO

A simple Framework based on [Raylib](https://www.raylib.com/) using [GO Bindings](https://github.com/Lachee/raylib-goplus) and [Flex](https://github.com/kjk/flex), a CSS flexbox layout implementation in Go.

## How to start
Flex and raylib-goplus are required:
```
$ go get github.com/kjk/flex
$ go get github.com/lachee/raylib-goplus/raylib
```

On Linux, xorg-dev must be installed:
```
$ apt-get install xorg-dev
```

### Example:
```go
package main

import (
	"github.com/Slixe/visual-go/app"
	"github.com/Slixe/visual-go/components"
	"github.com/Slixe/visual-go/graphics"
	"github.com/Slixe/visual-go/panels"
	"github.com/Slixe/visual-go/structures"
	"github.com/kjk/flex"
	rl "github.com/lachee/raylib-goplus/raylib"
)

func main() {
	window := app.App{
		Title:        "Hello World",
		Width:        780,
		Height:       640,
		ScrollSpeed:  20,
		Resizable:    true,
		DrawFPS:      true,
		DefaultColor: rl.GopherBlue,
		Flex: app.Flex{
			Direction:  flex.DirectionLTR,
		},
	}
	window.Start() //start application

	font := *window.GetFont()//get default font
	mainLayout := window.NewLayout() //create global layout
	mainLayout.StyleSetFlexGrow(1)
	fontSize := float32(42)
	spacing := float32(0)

	panel := panels.CreateColoredPanel(rl.White) //create colored panel
	msg := "Hello World!"
	msgSize := graphics.MeasureText(font, msg, fontSize, spacing) //measure text size

	panel.AddComponent(components.CreateText(font, msg, rl.Black, func(g structures.IGraphics, app structures.IApp) structures.Vector4f {
		return graphics.Position((g.GetWidth() - msgSize.X) / 2, (g.GetHeight() - msgSize.Y) / 2, msgSize.X, msgSize.Y)
	}, fontSize, spacing))

	window.RegisterLayout("global", mainLayout) //register our layout
	window.SetPanel("global", panel) //set panel using global layout
	window.Render() //start rendering
}
```
### Scaling Example
![scale example](https://github.com/Slixe/visual-go/blob/master/scale_example.gif?raw=true)
