package engine

import "github.com/veandco/go-sdl2/sdl"

const (
    WindowWidth  = 800
    WindowHeight = 600
)

// Game struct
type Game struct {
    renderer     *sdl.Renderer
    player       *GameObject
    inputHandler InputHandler
    running      bool
}

// NewGame constructor
func NewGame(renderer *sdl.Renderer) *Game {
    return &Game{
        renderer:     renderer,
        running:      true,
        inputHandler: NewInputHandler(),
    }
}

// Initialize method
func (g *Game) Initialize() error {
    player, err := NewGameObject(g.renderer, 100, 100, 50, 50, "assets/images/istockphoto-174931919-1024x1024.jpg")
    if err != nil {
        return err
    }
    g.player = player
    return nil
}

// CleanUp method
func (g *Game) CleanUp() {
    g.player.Texture().Destroy()
}

// Run method
func (g *Game) Run() {
    for g.running {
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
            case *sdl.QuitEvent:
                g.running = false
            }
        }

        g.inputHandler.HandleInput(g.player)

        g.renderer.SetDrawColor(0, 0, 0, 255)
        g.renderer.Clear()
        g.player.Draw(g.renderer)
        g.renderer.Present()

        sdl.Delay(16)
    }
}

