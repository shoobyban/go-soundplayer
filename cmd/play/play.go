package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/shoobyban/go-soundplayer"
)

func main() {
	time.AfterFunc(time.Second*3, func() {
		fmt.Println("seeking after 3 seconds into 5 seconds")
		fmt.Println(math.Floor(soundplayer.GetSecs()), math.Floor(soundplayer.GetDuration()), soundplayer.TimeScale())
		soundplayer.Seek(5)
		fmt.Println(math.Floor(soundplayer.GetSecs()), math.Floor(soundplayer.GetDuration()), soundplayer.TimeScale())
	})
	time.AfterFunc(time.Second*10, func() {
		fmt.Println(math.Floor(soundplayer.GetSecs()), math.Floor(soundplayer.GetDuration()), soundplayer.TimeScale())
		fmt.Println("pausing after 10 seconds")
		soundplayer.Pause()
	})
	time.AfterFunc(time.Second*12, func() {
		fmt.Println(math.Floor(soundplayer.GetSecs()), math.Floor(soundplayer.GetDuration()), soundplayer.TimeScale())
		fmt.Println("resuming after 12 seconds")
		soundplayer.Resume()
	})
	time.AfterFunc(time.Second*15, func() {
		fmt.Println("stopping after 13 seconds")
		soundplayer.Stop()
	})
	for _, arg := range os.Args[1:] {
		if err := soundplayer.Play(arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
