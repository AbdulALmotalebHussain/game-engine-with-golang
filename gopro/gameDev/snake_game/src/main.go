package main

import (
    "fmt"
    "os"

    "snake_game/src/engine"
    "github.com/veandco/go-sdl2/sdl"
)

func main() {
    if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
        fmt.Fprintf(os.Stderr, "Failed to initialize SDL: %v\n", err)
        os.Exit(1)
    }
    defer sdl.Quit()

    window, err := sdl.CreateWindow("Go Game Engine", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, engine.WindowWidth, engine.WindowHeight, sdl.WINDOW_SHOWN)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create window: %v\n", err)
        os.Exit(1)
    }
    defer window.Destroy()

    renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create renderer: %v\n", err)
        os.Exit(1)
    }
    defer renderer.Destroy()

    game := engine.NewGame(renderer)
    if err := game.Initialize(); err != nil {
        fmt.Fprintf(os.Stderr, "Failed to initialize game: %v\n", err)
        os.Exit(1)
    }

    game.Run()

    game.CleanUp()
}

