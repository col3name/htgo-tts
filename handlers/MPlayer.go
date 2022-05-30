package handlers

import (
	"os/exec"
	"strconv"
)

type MPlayer struct {
	Volume int
}

func (MPlayer *MPlayer) Play(fileName string) error {
	mplayer := exec.Command("mplayer", "-cache", "106092", "-", fileName, "-af", "volume="+strconv.Itoa(MPlayer.Volume))
	return mplayer.Run()
}
