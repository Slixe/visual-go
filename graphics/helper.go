package graphics

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/kjk/flex"
)

var textures = map[string]rl.Texture2D{}
var fonts = map[string]*rl.Font{}

func CreateGraphics(node flex.Node) Graphics {
	return Graphics{
		posX: node.LayoutGetLeft(),
		posY: node.LayoutGetTop(),
		maxWidth: node.LayoutGetWidth(),
		maxHeight: node.LayoutGetHeight(),
	}
}

func LoadFont(fontPath string, fontSize int) *rl.Font {
	font, ok := fonts[fontPath]
	if !ok {
		font = rl.LoadFontEx(fontPath, fontSize, nil, 256)
		fonts[fontPath] = font
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