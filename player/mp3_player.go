package player

import (
	"fmt"
	"time"
)

type Mp3Player struct {
	//player's status
	status string

	//player's progress
	progress int
}

func (m *Mp3Player)New() *Mp3Player{
	return &Mp3Player{"stop", 0}
}

func Mp3PlayerNew() *Mp3Player{
	return &Mp3Player{"stop", 0}
}

func (m *Mp3Player)Play(Path string) int{
	m.status = "playing"
	m.progress = 0

	fmt.Println("start play mp3 file.")

	for i:=0; i <=100; i++{
		fmt.Printf("progress %d%%\n", i)
		time.Sleep(1*time.Second)
	}

	return 0
}