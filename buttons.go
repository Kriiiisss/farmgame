package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var dirEntries []os.DirEntry

func LoadButtons() {
	welcomeButtons = []Button{
		{
			"Manage Worlds",
			0,
			func() {
				menuSection = MANAGE_WORLDS
				worldsSection = MNGWORLD_WELCOME
				RefreshSaves()
			},
			rl.Rectangle{},
			false,
			true,
		},
		{
			"Options",
			0,
			func() {
				menuSection = OPTIONS
			},
			rl.Rectangle{},
			false,
			false,
		},
	}
	manageWorldsButtons = []Button{
		{
			"Create New World",
			0,
			func() {
				worldsSection = CREATE_NEW_WORLD
			},
			rl.Rectangle{},
			false,
			true,
		},
		{
			"Load World",
			0,
			func() {
				if selectedSaveId != -1 {
					clientState = LOADING_TO_WORLD
					loadedSaveId = selectedSaveId
				}
			},
			rl.Rectangle{},
			false,
			true,
		},
		{
			"Delete World",
			0,
			func() {
				worldsSection = DELETE_WORLD
			},
			rl.Rectangle{},
			false,
			true,
		},
		{
			"Refresh",
			0,
			func() {
				RefreshSaves()
			},
			rl.Rectangle{},
			false,
			true,
		},
	}
	optionsButtons = []Button{
		{
			"Volume",
			0,
			func() {
				optionsSection = VOLUME
			},
			rl.Rectangle{},
			false,
			true,
		},
	}
	deleteWorldButtons = []Button{
		{
			"Delete",
			0,
			func() {
				err := os.RemoveAll("./saves/" + saves[selectedSaveId].Name)
				if err != nil {
					log.Fatal(err)
				}
				RefreshSaves()
				worldsSection = MNGWORLD_WELCOME
			},
			rl.Rectangle{},
			false,
			true,
		},
		{
			"Cancel",
			0,
			func() {
				worldsSection = MNGWORLD_WELCOME
			},
			rl.Rectangle{},
			false,
			true,
		},
	}
	createWorldButtons = []Button{
		{
			"Create",
			0,
			func() {
				nameIllegal := slices.Contains(illegalNames, createdSave.Name)
				if len(saves) > 0 {
					for saveId := range saves {
						if saves[saveId].Name == createdSave.Name {
							nameIllegal = true
							break
						}
					}
				}
				if !nameIllegal {
					err := os.Mkdir("./saves/"+createdSave.Name, os.ModeDir)
					if err != nil {
						log.Fatal(err)
					}
					GenerateSaveFiles(createdSave.Name, createdSave.MapName)
					createdSave.Name = ""
					createdSave.MapName = ""
					RefreshSaves()
					worldsSection = MNGWORLD_WELCOME
				}
			},
			rl.Rectangle{},
			false,
			true,
		},
		{
			"Cancel",
			0,
			func() {
				createdSave.Name = ""
				createdSave.MapName = ""
				worldsSection = MNGWORLD_WELCOME
			},
			rl.Rectangle{},
			false,
			true,
		},
	}
}

func DrawButton(button Button) {
	var color rl.Color
	if button.Available {
		if button.Hovered {
			color = rl.SkyBlue
		} else {
			color = rl.White
		}
	} else {
		color = rl.Gray
	}

	rl.DrawText(button.Text, button.Hitbox.ToInt32().X, button.Hitbox.ToInt32().Y, button.Fontsize, color)
}

func RefreshSaves() {
	var err error
	dirEntries, err = os.ReadDir("./saves")
	if err != nil {
		log.Fatal(err)
	}
	saveButtons = []Button{}
	saves = []Save{}
	for dirEntry := range dirEntries {
		saveButtons = append(saveButtons, Button{
			dirEntries[dirEntry].Name(),
			int32(rl.GetRenderHeight()) / 28,
			func() {},
			rl.Rectangle{
				X:      0.15*float32(rl.GetRenderWidth()) + 0.05*float32(rl.GetRenderWidth()),
				Y:      0.1 * float32(rl.GetRenderHeight()) * float32(1+dirEntry),
				Width:  float32(0.6 * float32(rl.GetRenderWidth())),
				Height: float32(rl.GetRenderHeight() / 28),
			},
			false,
			true,
		})
		metadataBytes, err := os.ReadFile("./saves/" + dirEntries[dirEntry].Name() + "/metadata")
		if err != nil {
			log.Fatal(err)
		}
		metadata := string(metadataBytes)
		metadataLines := strings.Split(metadata, "\n")
		lastLaunchTimeInt, err := strconv.ParseInt(metadataLines[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		lastLaunchTime := time.Unix(0, lastLaunchTimeInt)

		saves = append(saves, Save{
			dirEntries[dirEntry].Name(),
			"map1",
			lastLaunchTime,
			lastLaunchTime.Format("2006-01-02 15:04:05"),
		})
	}
	selectedSaveId = -1
}
