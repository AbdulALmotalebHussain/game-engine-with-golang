package engine

import "github.com/veandco/go-sdl2/sdl"

const speed = 5.0

// InputHandler interface
type InputHandler interface {
    HandleInput(player GameObject)
}

// inputHandler struct
type inputHandler struct{}

// NewInputHandler constructor
func NewInputHandler() InputHandler {
    return &inputHandler{}
}

// HandleInput method
func (ih *inputHandler) HandleInput(player GameObject) {
    keys := sdl.GetKeyboardState()
    p, ok := player.(*gameObject)
    if !ok {
        return
    }
    if keys[sdl.SCANCODE_W] != 0 {
        p.y -= speed
    }
    if keys[sdl.SCANCODE_S] != 0 {
        p.y += speed
    }
    if keys[sdl.SCANCODE_A] != 0 {
        p.x -= speed
    }
    if keys[sdl.SCANCODE_D] != 0 {
        p.x += speed
    }
}

