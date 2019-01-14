package player

import (
	"fmt"
	"time"
)

type WavPlayer struct {
	//player's status
	status string

	//player's progress
	progress int
}

func WavPlayerNew() *WavPlayer{
	return &WavPlayer{"stop", 0}
}

func (m *WavPlayer)Play(Path string) int{
	m.status = "playing"
	m.progress = 0

	fmt.Println("start play Wav file.")

	for i:=0; i <=100; i++{
		fmt.Printf("progress %d%%\n", i)
		time.Sleep(1*time.Second)
	}

	return 0
}