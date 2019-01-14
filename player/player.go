package player

import "fmt"

type Player interface {
	Play(Path string) int
}

func Play(Path string, AudioFormat string) int{

	var p Player

	switch AudioFormat{
	case "Mp3":fallthrough
	case "MP3":
		//todo 爲啥不能直接使用struct的new函數，而是需要struct之外的函數才行
		//todo 因爲struct的成員需要是對象的實例才能調用，需要傳入實例或者實例的引用或者指針，類似C++的this指針
		p = Mp3PlayerNew()
	case "Wav":fallthrough
	case "WAV":
		p = WavPlayerNew()
	default:
		fmt.Println("AudioFormat not correctly!")
	}

	return p.Play(Path)
}