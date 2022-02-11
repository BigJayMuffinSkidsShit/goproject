package main

import (
	"github.com/nsf/termbox-go"
)

// use this for later ||

func main() {
	err := termbox.Init()

	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	eventQueue := make(chan termbox.Event)

	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	g := termgame.NewGame()

	//g.render()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowDown:
					//Move Down
				case ev.Key == termbox.KeyArrowUp:
					//Move Up
				case ev.Key == termbox.KeyArrowLeft:
					//Move Left
				case ev.Key == termbox.KeyArrowRight:
					//Move Right
				case ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC:
					return
				}
			}
		default:

		}
	}

}
