package main

import (
	"log"
	"os"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var dirEntries []os.DirEntry

func LoadButtons() {
	welcomeButtons = []Button{
		{
			"Manage Worlds",
			0,
			func() {
				rl.PlaySound(click)
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
				rl.PlaySound(click)
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
				rl.PlaySound(click)
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
					rl.PlaySound(click)
					clientState = LOADING_TO_WORLD
					loadedSaveId = selectedSaveId
					UpdateSaveMetadata(saves[loadedSaveId])
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
				rl.PlaySound(click)
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
				rl.PlaySound(click)
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
				rl.PlaySound(click)
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
				rl.PlaySound(click)
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
				rl.PlaySound(click)
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
				rl.PlaySound(click)
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
					GenerateSaveFilesJSON(createdSave)
					// GenerateSaveFilesBinary(createdSave)
					UpdateSaveMetadata(createdSave)
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
				rl.PlaySound(click)
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
