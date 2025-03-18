package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"math/rand"
	"strings"
	"time"
	"strconv"
)

var playlists map[string][]string
var currentPlaylist []string
var currentIndex int
var isPaused bool
var isShuffle bool
var isLooping bool
var currentVolume int
var isMuted bool

const reset = "\033[0m"
const bold = "\033[1m"
const red = "\033[31m"
const green = "\033[32m"
const yellow = "\033[33m"
const blue = "\033[34m"
const magenta = "\033[35m"
const cyan = "\033[36m"
const gray = "\033[90m"

func main() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)
	playlists = make(map[string][]string)
	currentPlaylist = make([]string, 0)
	currentVolume = 100

	for {
		printPrompt()
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		args := strings.Split(input, " ")

		if len(args) < 1 {
			continue
		}

		command := args[0]

		switch command {
		case "play":
			if len(args) < 2 {
				printError("Usage: play <filename>")
				continue
			}
			file := args[1]
			playMusic(file)

		case "list":
			listPlaylist()

		case "add":
			if len(args) < 2 {
				printError("Usage: add <filename>")
				continue
			}
			file := args[1]
			addToPlaylist(file)

		case "remove":
			if len(args) < 2 {
				printError("Usage: remove <filename>")
				continue
			}
			file := args[1]
			removeFromPlaylist(file)

		case "create":
			if len(args) < 2 {
				printError("Usage: create <playlist_name>")
				continue
			}
			playlistName := args[1]
			createPlaylist(playlistName)

		case "switch":
			if len(args) < 2 {
				printError("Usage: switch <playlist_name>")
				continue
			}
			playlistName := args[1]
			switchPlaylist(playlistName)

		case "shuffle":
			toggleShuffle()

		case "pause":
			pauseMusic()

		case "resume":
			resumeMusic()

		case "volume":
			if len(args) < 2 {
				printError("Usage: volume <level>")
				continue
			}
			volumeLevel := args[1]
			setVolume(volumeLevel)

		case "loop":
			toggleLoop()

		case "next":
			playNext()

		case "prev":
			playPrevious()

		case "mute":
			toggleMute()

		case "quit":
			printSuccess("Goodbye!")
			return

		default:
			printError("Unknown command. Please use play, list, add, remove, create, switch, shuffle, pause, resume, volume, loop, next, prev, mute, or quit.")
		}
	}
}

func playMusic(file string) {
	if !fileExists(file) {
		printError("Error: File does not exist: " + file)
		return
	}
	cmd := exec.Command("mpv", "--volume="+fmt.Sprintf("%d", currentVolume), file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		printError("Error playing file: " + err.Error())
		return
	}
	if !isPaused {
		currentPlaylist = append(currentPlaylist, file)
	}
	printSuccess("Now playing: " + file)
	cmd.Wait()
	if !isShuffle && !isLooping {
		currentIndex = len(currentPlaylist) - 1
	}
}

func listPlaylist() {
	if len(currentPlaylist) == 0 {
		printError("Playlist is empty.")
		return
	}
	printSuccess("Current Playlist:")
	for i, song := range currentPlaylist {
		fmt.Printf("%s%d. %s%s\n", cyan, i+1, song, reset)
	}
}

func addToPlaylist(file string) {
	if !fileExists(file) {
		printError("Error: File does not exist: " + file)
		return
	}
	currentPlaylist = append(currentPlaylist, file)
	printSuccess("Added to playlist: " + file)
}

func removeFromPlaylist(file string) {
	index := -1
	for i, song := range currentPlaylist {
		if song == file {
			index = i
			break
		}
	}
	if index == -1 {
		printError("Error: File not found in playlist: " + file)
		return
	}
	currentPlaylist = append(currentPlaylist[:index], currentPlaylist[index+1:]...)
	printSuccess("Removed from playlist: " + file)
}

func createPlaylist(playlistName string) {
	if _, exists := playlists[playlistName]; exists {
		printError("Playlist already exists: " + playlistName)
		return
	}
	playlists[playlistName] = make([]string, 0)
	printSuccess("Playlist created: " + playlistName)
}

func switchPlaylist(playlistName string) {
	if playlist, exists := playlists[playlistName]; exists {
		currentPlaylist = playlist
		printSuccess("Switched to playlist: " + playlistName)
	} else {
		printError("Error: Playlist does not exist: " + playlistName)
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func toggleShuffle() {
	isShuffle = !isShuffle
	if isShuffle {
		printSuccess("Shuffle mode enabled.")
		shufflePlaylist()
	} else {
		printSuccess("Shuffle mode disabled.")
	}
}

func shufflePlaylist() {
	rand.Shuffle(len(currentPlaylist), func(i, j int) {
		currentPlaylist[i], currentPlaylist[j] = currentPlaylist[j], currentPlaylist[i]
	})
}

func pauseMusic() {
	if isPaused {
		printError("Music is already paused.")
		return
	}
	isPaused = true
	printSuccess("Music paused.")
}

func resumeMusic() {
	if !isPaused {
		printError("Music is already playing.")
		return
	}
	isPaused = false
	if currentIndex < len(currentPlaylist) {
		playMusic(currentPlaylist[currentIndex])
	}
	printSuccess("Music resumed.")
}

func setVolume(level string) {
	volume, err := strconv.Atoi(level)
	if err != nil || volume < 0 || volume > 100 {
		printError("Error: Invalid volume level. Please enter a value between 0 and 100.")
		return
	}
	currentVolume = volume
	printSuccess("Volume set to: " + fmt.Sprintf("%d", volume))
}

func toggleLoop() {
	isLooping = !isLooping
	if isLooping {
		printSuccess("Loop mode enabled.")
	} else {
		printSuccess("Loop mode disabled.")
	}
}

func playNext() {
	if currentIndex+1 < len(currentPlaylist) {
		currentIndex++
		playMusic(currentPlaylist[currentIndex])
	} else if isLooping {
		currentIndex = 0
		playMusic(currentPlaylist[currentIndex])
	} else {
		printError("No next song in the playlist.")
	}
}

func playPrevious() {
	if currentIndex-1 >= 0 {
		currentIndex--
		playMusic(currentPlaylist[currentIndex])
	} else if isLooping {
		currentIndex = len(currentPlaylist) - 1
		playMusic(currentPlaylist[currentIndex])
	} else {
		printError("No previous song in the playlist.")
	}
}

func toggleMute() {
	isMuted = !isMuted
	if isMuted {
		printSuccess("Muted.")
		currentVolume = 0
	} else {
		printSuccess("Unmuted.")
		currentVolume = 100
	}
}

func printPrompt() {
	fmt.Print(blue + bold + "Enter command: " + reset)
}

func printError(message string) {
	fmt.Println(red + bold + message + reset)
}

func printSuccess(message string) {
	fmt.Println(green + bold + message + reset)
}
