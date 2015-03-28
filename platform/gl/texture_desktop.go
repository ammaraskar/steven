// +build !mobile

package gl

import (
	"github.com/go-gl/gl/v2.1/gl"
)

const (
	Texture2D TextureTarget = gl.TEXTURE_2D

	RGB  TextureFormat = gl.RGB
	RGBA TextureFormat = gl.RGBA

	TextureMinFilter TextureParameter = gl.TEXTURE_MIN_FILTER
	TextureMagFilter TextureParameter = gl.TEXTURE_MAG_FILTER
	TextureWrapS     TextureParameter = gl.TEXTURE_WRAP_S
	TextureWrapT     TextureParameter = gl.TEXTURE_WRAP_T

	Nearest              TextureValue = gl.NEAREST
	Linear               TextureValue = gl.LINEAR
	LinearMipmapLinear   TextureValue = gl.LINEAR_MIPMAP_LINEAR
	LinearMipmapNearest  TextureValue = gl.LINEAR_MIPMAP_NEAREST
	NearestMipmapNearest TextureValue = gl.NEAREST_MIPMAP_NEAREST
	NearestMipmapLinear  TextureValue = gl.NEAREST_MIPMAP_LINEAR
	ClampToEdge          TextureValue = gl.CLAMP_TO_EDGE
)

// State tracking
var (
	currentTexture       Texture
	currentTextureTarget TextureTarget
)

type TextureTarget uint32
type TextureFormat uint32
type TextureParameter uint32
type TextureValue int32

type Texture struct {
	internal uint32
}

func CreateTexture() Texture {
	var texture Texture
	gl.GenTextures(1, &texture.internal)
	return texture
}

func (t Texture) Bind(target TextureTarget) {
	if currentTexture == t && currentTextureTarget == target {
		return
	}
	gl.BindTexture(uint32(target), t.internal)
	currentTexture = t
	currentTextureTarget = target
}

func (t Texture) Image2D(level int, width, height int, format TextureFormat, ty Type, pix []byte) {
	if t != currentTexture {
		panic("texture not bound")
	}
	gl.TexImage2D(
		uint32(currentTextureTarget),
		int32(level),
		int32(format),
		int32(width),
		int32(height),
		0,
		uint32(format),
		uint32(ty),
		gl.Ptr(pix),
	)
}

func (t Texture) Parameter(param TextureParameter, val TextureValue) {
	if t != currentTexture {
		panic("texture not bound")
	}
	gl.TexParameteri(uint32(currentTextureTarget), uint32(param), int32(val))
}
