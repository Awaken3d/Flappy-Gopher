package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("cannot initialize sdl")
		os.Exit(2)
	}
}
	func run() error{
		err := sdl.Init(sdl.INIT_EVERYTHING)

		if err != nil{
		return fmt.Errorf("cannot initialize sdl")

	}

		defer sdl.Quit()

		if err:=ttf.Init(); err!=nil {
			return fmt.Errorf("cannot initialize ttf")
		}

		w, r, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)

		if err != nil {
			return fmt.Errorf("could not create window: %v", err)
		}

		defer w.Destroy()
		if err := drawTitle(r); err!= nil {
			return fmt.Errorf("could not draw title %v, err")
		}
		time.Sleep(5 * time.Second)
		if err:=drawBackground(r);err!=nil {
			return fmt.Errorf("could not draw background %v",err)
		}
		time.Sleep(5 * time.Second)
		return nil
}

func drawTitle(r *sdl.Renderer) error {
	r.Clear()
	f, err := ttf.OpenFont("res/fonts/Flappy.ttf",20)

	if err != nil {
		return fmt.Errorf("could not load font: %v", err)
	}

	defer f.Close()
	c := sdl.Color{R: 255, G: 100, B: 0, A:255}
	s, err := f.RenderUTF8Solid("Flappy Gopher", c)

	if err != nil {
		return fmt.Errorf("could not render title %v", err)
	}

	defer s.Free()
	t, err:= r.CreateTextureFromSurface(s)

	if err != nil {
		return fmt.Errorf("could not create texture")
	}

	defer t.Destroy()
	r.Copy(t,nil,nil)

	r.Present()

	return nil

}

func drawBackground(r *sdl.Renderer) error {
	r.Clear()

	t, err := img.LoadTexture(r,"res/imgs/background.png")

	if err != nil {
		return fmt.Errorf("could not load background %v",err)
	}

	defer t.Destroy()

	if err := r.Copy(t,nil,nil); err != nil {
		return fmt.Errorf("Could not copy background to renderer %v", err)
	}

	r.Present()

	return nil
}