package main

import "github.com/veandco/go-sdl2/sdl"

func SDLInit() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	SmartPanic(err)
}
func CreateWindow(title string, width int32, height int32) *sdl.Window {
	SDLInit()
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, width, height, sdl.WINDOW_SHOWN)
	window.SetResizable(false)
	SmartPanic(err)
	return window
}

func CreateRenderer(window *sdl.Window) *sdl.Renderer {
	renderer, err := sdl.CreateRenderer(window, -1, 0)
	SmartPanic(err)
	return renderer
}

func CreateSurface(size size, color color) *sdl.Surface {
	surface, err := sdl.CreateRGBSurface(0, int32(size.width), int32(size.height), 32, 0, 0, 0, 0)
	SmartPanic(err)
	return surface
}

func GetWindowSurface(window *sdl.Window) *sdl.Surface {
	surface, err := window.GetSurface()
	SmartPanic(err)
	return surface
}
