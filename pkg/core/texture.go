package core

import (
	ray "github.com/gen2brain/raylib-go/raylib"
)

type Texture struct {
	Texture  ray.Texture2D
	Dest     ray.Rectangle
	Src      ray.Rectangle
	Rotation float32
	Tint     ray.Color
	Filter   ray.TextureFilterMode
	OnDraw   func()
}

func NewTexture(path string, src ray.Rectangle) *Texture {
	var t Texture
	t.Texture = ray.LoadTexture(path)
	if !ray.IsTextureReady(t.Texture) {
		panic("texture loading failed")
	}
	t.Src = src
	t.Tint = ray.White
	t.Filter = ray.FilterTrilinear
	t.OnDraw = nil
	ray.SetTextureFilter(t.Texture, t.Filter)
	return &t
}

func (t *Texture) SetSrc(src ray.Rectangle) *Texture {
	t.Src = src
	return t
}

func (t *Texture) SetDst(dst ray.Rectangle) *Texture {
	t.Dest = dst
	return t
}

func (t *Texture) SetTextureFilter(f ray.TextureFilterMode) *Texture {
	t.Filter = f
	ray.SetTextureFilter(t.Texture, f)
	return t
}

func (t *Texture) Unload() {
	ray.UnloadTexture(t.Texture)
}

func (t *Texture) Draw() {
	if t.OnDraw != nil {
		t.OnDraw()
	}
	if t.Dest.Width == 0 || t.Dest.Height == 0 {
		return
	}
	ray.DrawTexturePro(t.Texture, t.Src, t.Dest, ray.NewVector2(0, 0), t.Rotation, t.Tint)
}

func (t *Texture) DrawAt(dst ray.Rectangle) {
	ray.DrawTexturePro(t.Texture, t.Src, dst, ray.NewVector2(0, 0), t.Rotation, t.Tint)
}
