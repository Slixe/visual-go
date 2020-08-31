package graphics

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
	"github.com/kjk/flex"
)

var textures = map[string]rl.Texture2D{}
var fonts = map[string]*rl.Font{}

func CreateGraphics(node flex.Node) structures.IGraphics {
	return &Graphics{
		posX:        node.LayoutGetLeft(),
		posY:        node.LayoutGetTop(),
		maxWidth:    node.LayoutGetWidth(),
		maxHeight:   node.LayoutGetHeight(),
		scrollX:     0,
		scrollY:     0,
		allowScroll: false,
	}
}

func LoadFont(fontPath string, fontSize int) *rl.Font {
	font, ok := fonts[fontPath]
	if !ok {
		font = rl.LoadFontEx(fontPath, fontSize, nil, 256)
		fonts[fontPath] = font
	}

	if font.BaseSize != int32(fontSize) {
		return rl.LoadFontEx(fontPath, fontSize, nil, 256)
	}

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