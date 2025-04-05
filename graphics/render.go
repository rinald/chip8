package graphics

import (
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

const SCALE_FACTOR = 15

type Renderer struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	texture  *sdl.Texture
}

func (r *Renderer) Init() {
	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_EVENTS); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("CHIP-8 Emulator", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WIDTH*SCALE_FACTOR, HEIGHT*SCALE_FACTOR, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	r.window = window

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	r.renderer = renderer

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STREAMING, WIDTH, HEIGHT)
	if err != nil {
		panic(err)
	}

	r.texture = texture
}

func (r *Renderer) Cleanup() {
	r.texture.Destroy()
	r.renderer.Destroy()
	r.window.Destroy()
	sdl.Quit()
}

func (r *Renderer) Update(g *Graphics) {
	r.texture.Update(nil, unsafe.Pointer(&g.Screen), WIDTH*4)
	r.renderer.Clear()
	r.renderer.Copy(r.texture, nil, nil)
	r.renderer.Present()
}

func (r *Renderer) ProcessInput(keypad *[16]bool) bool {
	quit := false

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			quit = true
		case *sdl.KeyboardEvent:
			key, state := event.(*sdl.KeyboardEvent).Keysym.Sym, event.(*sdl.KeyboardEvent).Type

			switch key {
			case sdl.K_ESCAPE:
				quit = true
			case sdl.K_1:
				keypad[0] = state == sdl.KEYDOWN
			case sdl.K_2:
				keypad[1] = state == sdl.KEYDOWN
			case sdl.K_3:
				keypad[2] = state == sdl.KEYDOWN
			case sdl.K_4:
				keypad[3] = state == sdl.KEYDOWN
			case sdl.K_q:
				keypad[4] = state == sdl.KEYDOWN
			case sdl.K_w:
				keypad[5] = state == sdl.KEYDOWN
			case sdl.K_e:
				keypad[6] = state == sdl.KEYDOWN
			case sdl.K_r:
				keypad[7] = state == sdl.KEYDOWN
			case sdl.K_a:
				keypad[8] = state == sdl.KEYDOWN
			case sdl.K_s:
				keypad[9] = state == sdl.KEYDOWN
			case sdl.K_d:
				keypad[10] = state == sdl.KEYDOWN
			case sdl.K_f:
				keypad[11] = state == sdl.KEYDOWN
			case sdl.K_z:
				keypad[12] = state == sdl.KEYDOWN
			case sdl.K_x:
				keypad[13] = state == sdl.KEYDOWN
			case sdl.K_c:
				keypad[14] = state == sdl.KEYDOWN
			case sdl.K_v:
				keypad[15] = state == sdl.KEYDOWN
			}
		}
	}

	return quit
}
