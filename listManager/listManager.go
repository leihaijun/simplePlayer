package listManager

import (
	"fmt"
	"strconv"
)

//The Music Info
type MusicInfo struct{
	Id string
	Name string
	Author string
	Path string
	AudioFormat string
}

//[]musicInfo
type MusicListManager struct{
	musicList []MusicInfo
}

//new
func New() *MusicListManager{

	return &MusicListManager{make([]MusicInfo, 0)}
}

//Add
func (mp *MusicListManager)Add(	Name string,
								Author string,
								Path string,
								AudioFormat string){
	if mp.isNull() == false {
		musicInfo := MusicInfo{strconv.Itoa(len(mp.musicList)),
								Name, Author,
								Path, AudioFormat}
		mp.musicList = append(mp.musicList, musicInfo)
	}
}

//Print
func (mp *MusicListManager)Print() int{
	if nil == mp.musicList || len(mp.musicList) <= 0{
		fmt.Println("mp.musicList is nil")
		return -1
	}

	fmt.Printf("%-6s%-5s%-10s%-10s%-5s\n", "Index", "Id", "Name", "Author", "Format")

	for i, musicinfo := range(mp.musicList){
		fmt.Printf("%-6d%-5s%-10s%-10s%-5s\n", i, musicinfo.Id, musicinfo.Name, musicinfo.Author, musicinfo.AudioFormat)
	}

	return 0
}


//Remove
func (mp *MusicListManager)Remove(Id string) *MusicInfo {

	if nil == mp.musicList{
		fmt.Println("mp.musicList is nil")
		return nil
	}

	if len(Id) <= 0{
		fmt.Println("Id of passed is nil")
		return nil
	}

	for i,musicInfo := range(mp.musicList){
		if musicInfo.Id ==  Id{
			//remove the item
			return mp.remove(i + 1)
		}
	}

	return nil
}


func (mp *MusicListManager)Get(index int) *MusicInfo{

	//index取值從1開始，數字是從0開始，這裏需要做一個轉換
	if false == mp.isNull() {
		if index < mp.Len()  {
			return &mp.musicList[index-1]
		}
	}

	return nil
}

func (mp *MusicListManager)Find(Id string) *MusicInfo{
	if nil == mp.musicList || len(mp.musicList) <= 0{
		fmt.Println("mp.musicList is nil")
		return nil
	}

	if len(Id) <= 0{
		fmt.Println("Id of passed is nil")
		return nil
	}

	for _,musicInfo := range(mp.musicList){
		if musicInfo.Id ==  Id{
			return &musicInfo
		}
	}

	return nil
}

func (mp *MusicListManager)Len() int{
	if false == mp.isNull() {
		return len(mp.musicList)
	}

	return 0
}


func (mp *MusicListManager)isNull() bool{
	return 	nil == mp.musicList
}

//delete for specified index
func (mp *MusicListManager)remove(index int) *MusicInfo{

	if nil == mp.musicList{
		fmt.Println("mp.musicList is nil")
		return nil
	}

	if index <= 0 || index > len(mp.musicList){
		fmt.Printf("index out of rang, index(%d) listLen(%d)\n", index, len(mp.musicList))
		return nil
	}

	removedMusicInfo := mp.musicList[index - 1]

	//	slice = append(slice, elem1, elem2)
	//	slice = append(slice, anotherSlice...)
	if len(mp.musicList) == 1{
		//stop player
		//todo
		//...

		//empty mp.musicList
		mp.musicList = make([]MusicInfo, 0)
	} else {
		mp.musicList = append(mp.musicList[:index - 1], mp.musicList[index:]...)
	}

	return &removedMusicInfo
}