package main

import "fmt"


type MusicPlayer struct {

}
type TapePlayer struct {

}
type Tape struct {
	voice string
}
type Music struct {
	song string
}
func (MusicPlayer MusicPlayer) play(music Music)  {
	fmt.Print(music.song)

}

func (TapePlayer) play(tape Tape) string{
	return tape.voice
}
func tape2Music (tape Tape ) Music{
	return Music{tape.voice};
}
