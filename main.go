package main

import (
    "log"
    "github.com/hajimehoshi/ebiten/v2"
    //"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{
   
    inited bool
}

func (g *Game) Update() error {
    if !g.inited {
        g.init()
    }
    //fmt.Printf("TPS: %d \nFPS: %d\nVsync: %v \n", int(ebiten.ActualTPS()), int(ebiten.ActualFPS()), ebiten.IsVsyncEnabled())
    return nil
}

var img *ebiten.Image

func (g *Game) Draw(screen *ebiten.Image) {
    
    //ebitenutil.DebugPrint(screen, "Hello, World!")
    chip8Cycle()
    if chip8.draw {
        op := &ebiten.DrawImageOptions{}
        screen.DrawImage(img, op)
        chip8.draw = false
    }

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 64, 32
}

func (g *Game) init(){
   
}

func main() {

    ebiten.SetWindowSize(768, 576)
    ebiten.SetWindowTitle("Chip8 Emulator")
    ebiten.SetScreenClearedEveryFrame(false)
    
    initChip8()
    loadROM()
	
    if err := ebiten.RunGame(&Game{}); err != nil {
        log.Fatal(err)
    }
}
