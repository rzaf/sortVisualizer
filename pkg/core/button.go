package core

import (
	ray "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	Texture   *Texture
	Boundary  ray.Rectangle
	OnClick   func()
	IsEnabled bool
}

func NewButton(texture *Texture, boundary ray.Rectangle) *Button {
	return &Button{texture, boundary, nil, true}
}

func (b *Button) SetTexture(texture *Texture) {
	b.Texture = texture
}

func (b *Button) Update() {
	if !b.IsEnabled {
		return
	}
	if b.OnClick != nil && ray.IsMouseButtonPressed(ray.MouseLeftButton) {
		if ray.CheckCollisionPointRec(ray.GetMousePosition(), b.Boundary) {
			b.OnClick()
		}
	}
}

func (b *Button) Draw() {
	b.Texture.Draw()
}

func (b *Button) DrawAt(dst ray.Rectangle) {
	b.Texture.DrawAt(dst)
}
