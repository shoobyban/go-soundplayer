package main

import (
	"fmt"
	"os"
	"time"

	"github.com/shoobyban/go-soundplayer"
)

func main() {
	time.AfterFunc(time.Second*5, func() {
		soundplayer.Pause()
	})
	for _, arg := range os.Args[1:] {
		if err := soundplayer.Play(arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
