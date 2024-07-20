package engine

import (
    "github.com/veandco/go-sdl2/sdl"
    "image"
    "image/jpeg"
    "os"
)

// GameObject represents an object in the game
type GameObject struct {
    texture *sdl.Texture
    x, y    int32
    width, height int32
}

// NewGameObject creates a new game object
func NewGameObject(renderer *sdl.Renderer, x, y, width, height int32, imagePath string) (*GameObject, error) {
    imgFile, err := os.Open(imagePath)
    if err != nil {
        return nil, err
    }
    defer imgFile.Close()

    img, err := jpeg.Decode(imgFile)
    if err != nil {
        return nil, err
    }

    rgba := image.NewRGBA(img.Bounds())
    if rgba.Stride != rgba.Rect.Size().X*4 {
        return nil, fmt.Errorf("unsupported stride")
    }

    for y := 0; y < rgba.Rect.Size().Y; y++ {
        for x := 0; x < rgba.Rect.Size().X; x++ {
            r, g, b, a := img.At(x, y).RGBA()
            rgba.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
        }
    }

    texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STATIC, rgba.Rect.Size().X, rgba.Rect.Size().Y)
    if err != nil {
        return nil, err
    }
    texture.Update(nil, unsafe.Pointer(&rgba.Pix[0]), rgba.Stride)

    return &GameObject{texture: texture, x: x, y: y, width: width, height: height}, nil
}

// Draw method
func (obj *GameObject) Draw(renderer *sdl.Renderer) {
    rect := sdl.Rect{X: obj.x, Y: obj.y, W: obj.width, H: obj.height}
    renderer.Copy(obj.texture, nil, &rect)
}

// Texture method
func (obj *GameObject) Texture() *sdl.Texture {
    return obj.texture
}

