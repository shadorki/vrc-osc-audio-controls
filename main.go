package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/hypebeast/go-osc/osc"
)

func main() {
	addr := "127.0.0.1:9001"
	d := osc.NewStandardDispatcher()
	d.AddMsgHandler("/avatar/parameters/OSC_AUDIO_CONTROLS_PLAY_PAUSE", handlePlayPauseSong)
	d.AddMsgHandler("/avatar/parameters/OSC_AUDIO_CONTROLS_NEXT", handleNextSong)
	d.AddMsgHandler("/avatar/parameters/OSC_AUDIO_CONTROLS_PREVIOUS", handlePreviousSong)
	d.AddMsgHandler("/avatar/parameters/OSC_AUDIO_CONTROLS_MUTE", handleMute)
	server := &osc.Server{
		Addr:       addr,
		Dispatcher: d,
	}
	server.ListenAndServe()
}

func handlePlayPauseSong(msg *osc.Message) {
	wasSelected, err := parseVRCBool(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !wasSelected {
		return
	}
	playPauseSong()
}

func handleNextSong(msg *osc.Message) {
	wasSelected, err := parseVRCBool(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !wasSelected {
		return
	}
	nextSong()
}

func handlePreviousSong(msg *osc.Message) {
	wasSelected, err := parseVRCBool(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !wasSelected {
		return
	}
	previousSong()
}

func handleMute(msg *osc.Message) {
	wasSelected, err := parseVRCBool(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !wasSelected {
		return
	}
	muteVolume()
}

func parseVRCBool(msg *osc.Message) (bool, error) {
	message := strings.Trim(msg.String(), " ")
	fmt.Println(message)
	if strings.HasSuffix(message, ",T true") {
		return true, nil
	} else if strings.HasSuffix(message, ",F false") {
		return false, nil
	}

	return false, fmt.Errorf("Unexpected value: %v\n", message)
}

func parseVRCFloat(msg *osc.Message) (float64, error) {
	message := strings.Trim(msg.String(), " ")
	messageSlice := strings.Split(message, " ")
	floatStr := messageSlice[len(messageSlice)-1]
	return strconv.ParseFloat(floatStr, 64)
}

func executePowershellCommand(command string) (string, error) {
	cmd := exec.Command("powershell", "-nologo", "-noprofile", command)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	defer stdin.Close()
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
func playPauseSong() (string, error) {
	return executePowershellCommand(`
		$wshell = New-Object -ComObject wscript.shell
		$wshell.SendKeys([Char]0xB3)
	`)
}
func nextSong() (string, error) {
	return executePowershellCommand(`
		$wshell = New-Object -ComObject wscript.shell
		$wshell.SendKeys([Char]0xB0)
	`)
}
func previousSong() (string, error) {
	return executePowershellCommand(`
		$wshell = New-Object -ComObject wscript.shell
		$wshell.SendKeys([Char]0xB1)
	`)
}
func muteVolume() (string, error) {
	return executePowershellCommand(`
		$wshell = New-Object -ComObject wscript.shell
		$wshell.SendKeys([Char]0xAD)
	`)
}
