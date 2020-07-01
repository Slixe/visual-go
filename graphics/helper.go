package graphics

import (
	rl "github.com/DankFC/raylib-goplus/raylib"
	"github.com/Slixe/visual-go/structures"
)

func DrawTexture(texture rl.Texture2D, posX float32, posY float32, width float32, height float32, rotation float32) {
	rect := rl.Rectangle{Y: posY, X: posX, Width: float32(texture.Width), Height: float32(texture.Height)}
	dst := rl.Rectangle{Y: posY, X: posX, Width: width, Height: height}
	rl.DrawTexturePro(texture, rect, dst, rl.Vector2{X: 0, Y: 0}, rotation, rl.RayWhite)
}

func CreateRectangle(component structures.BaseComponent) rl.Rectangle {
	return rl.NewRectangle(component.PosX, component.PosY, component.Width, component.Height)
}
