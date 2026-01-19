package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func IsInRange(number int, min int, max int) bool {
	return (number >= min && number <= max)
}

func Clamp(number float32, min float32, max float32) float32 {
	if number < min {
		return min
	}
	if number > max {
		return max
	}
	return number
}

func Floor(number float32) int {
	return int(number)
}

func Ceil(number float32) int {
	return int(number) + 1
}

func HandleTextInput(text *string, maxLength int) {
	character := rl.GetCharPressed()
	if character != 0 && !slices.Contains(illegalRunes, rune(character)) && len(*text) < 48 {
		*text += string(character)
	}
	if rl.GetKeyPressed() == rl.KeyBackspace {
		_, lastRuneLength := utf8.DecodeLastRuneInString(*text)
		*text = (*text)[:len(*text)-lastRuneLength]
	}
}

func RefreshSaves() {
	var err error
	dirEntries, err = os.ReadDir("./saves")
	if err != nil {
		log.Fatal(err)
	}
	saves = []Save{}
	for dirEntry := range dirEntries {
		metadataBytes, err := os.ReadFile("./saves/" + dirEntries[dirEntry].Name() + "/metadata")
		if err != nil {
			log.Fatal(err)
		}
		metadata := string(metadataBytes)
		metadataLines := strings.Split(metadata, "\n")
		lastPlayedInt, err := strconv.ParseInt(metadataLines[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		lastPlayedTime := time.Unix(0, lastPlayedInt)

		saves = append(saves, Save{
			dirEntries[dirEntry].Name(),
			"map1",
			lastPlayedTime,
			lastPlayedTime.Format("2006-01-02 15:04:05"),
			Button{
				dirEntries[dirEntry].Name(),
				int32(rl.GetRenderHeight()) / 28,
				func() {},
				rl.Rectangle{},
				false,
				true,
			},
		})
	}
	slices.SortFunc(saves, func(a Save, b Save) int {
		if a.LastPlayed.After(b.LastPlayed) {
			return -1
		} else if b.LastPlayed.After(b.LastPlayed) {
			return 1
		} else {
			return 0
		}
	})
	for saveId := range saves {
		saves[saveId].MenuButton.Hitbox.Width = float32(0.6 * float32(rl.GetRenderWidth()))
		saves[saveId].MenuButton.Hitbox.Height = float32(rl.GetRenderHeight() / 28)
		saves[saveId].MenuButton.Hitbox.X = 0.15*float32(rl.GetRenderWidth()) + 0.05*float32(rl.GetRenderWidth())
		saves[saveId].MenuButton.Hitbox.Y = 0.1 * float32(rl.GetRenderHeight()) * float32(1+saveId)
	}

	selectedSaveId = -1
}

func UpdateSaveMetadata(save Save) {
	savePath := "./saves/" + save.Name
	metadataPath := savePath + "/metadata"

	metadataFile, err := os.Create(metadataPath)
	if err != nil {
		log.Fatal(err)
	}

	fileLines := []string{
		fmt.Sprintf("%d", time.Now().UTC().UnixNano()),
	}
	var fileContents strings.Builder
	for line := range fileLines {
		fileContents.WriteString(fileLines[line] + "\n")
	}
	fileContentsBytes := []byte(fileContents.String())

	err = os.WriteFile(metadataPath, fileContentsBytes, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	metadataFile.Close()
}
