package graphics

import (
	"github.com/Slixe/visual-go/structures"
	"github.com/kjk/flex"
	rl "github.com/lachee/raylib-goplus/raylib"
)

var textures = map[string]rl.Texture2D{}
var fonts = map[string]*rl.Font{}

func CreateGraphics(node flex.Node) structures.IGraphics {
	return &Graphics{
		posX:        node.LayoutGetLeft(),
		posY:        node.LayoutGetTop(),
		maxWidth:    node.LayoutGetWidth(),
		maxHeight:   node.LayoutGetHeight(),
	}
}

func UpdateGraphics(g Graphics, node flex.Node) structures.IGraphics {
	g.posX = node.LayoutGetLeft()
	g.posY = node.LayoutGetTop()
	g.maxWidth = node.LayoutGetWidth()
	g.maxHeight = node.LayoutGetHeight()
	g.verticalScroll.scrollable = false
	g.verticalScroll.maxValue = 0
	g.horizontalScroll.scrollable = false
	g.horizontalScroll.maxValue = 0

	return &g
}

func LoadFont(fontPath string, fontSize int) *rl.Font {
	font, ok := fonts[fontPath]
	if !ok {
		font = rl.LoadFont(fontPath)
		fonts[fontPath] = font
	}

	/*if font.BaseSize != int32(fontSize) {
		font, _ = rl.LoadFontEx(fontPath, fontSize, 0, 256)
		fonts[fontPath] = font
	}*/

	return font
}

func LoadTexture(texturePath string) rl.Texture2D {
	texture, ok := textures[texturePath]
	if !ok  {
		texture = rl.LoadTexture(texturePath)
		textures[texturePath] = texture
	}

	return texture
}

func AbsPos(posX float32, posY float32, width float32, height float32) func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f {
	return func(graphics structures.IGraphics, app structures.IApp) structures.Vector4f {
		return Position(posX, posY, width, height)
	}
}

func Position(posX float32, posY float32, width float32, height float32) structures.Vector4f {
	return structures.Vector4f{
		PosX:   posX,
		PosY:   posY,
		Width:  width,
		Height: height,
	}
}

func MeasureText(font rl.Font, text string, fontSize float32, spacing float32) rl.Vector2 {
	return rl.MeasureTextEx(font, text, fontSize, spacing)
}