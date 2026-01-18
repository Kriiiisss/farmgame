package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawDebug() {
	lines := 7

	debugMenu := rl.Rectangle{X: DEBUG_PADDING, Y: DEBUG_PADDING, Width: DEBUG_WIDTH, Height: float32(lines*(DEBUG_PADDING+1) + lines*int(DEBUG_FONTSIZE_SCREEN))}
	rl.DrawRectangleRec(debugMenu, rl.Black)

	rl.DrawText(fmt.Sprintf("%d FPS", rl.GetFPS()), 2*DEBUG_PADDING, 2*DEBUG_PADDING, DEBUG_FONTSIZE_SCREEN, rl.White)
	if playerCamOn {
		rl.DrawText("Cam.: Player", 2*DEBUG_PADDING, 3*DEBUG_PADDING+DEBUG_FONTSIZE_SCREEN, DEBUG_FONTSIZE_SCREEN, rl.White)
	} else {
		rl.DrawText("Cam.: Free", 2*DEBUG_PADDING, 3*DEBUG_PADDING+DEBUG_FONTSIZE_SCREEN, DEBUG_FONTSIZE_SCREEN, rl.White)
	}
	rl.DrawText(fmt.Sprintf("Zoom: %f", currentCam.Zoom), 2*DEBUG_PADDING, 4*DEBUG_PADDING+2*DEBUG_FONTSIZE_SCREEN, DEBUG_FONTSIZE_SCREEN, rl.White)
	rl.DrawText("Mouse Position:", 2*DEBUG_PADDING, 5*DEBUG_PADDING+3*DEBUG_FONTSIZE_SCREEN, DEBUG_FONTSIZE_SCREEN, rl.White)
	rl.DrawText(fmt.Sprintf("> World %.0f %.0f", worldMousePos.X, worldMousePos.Y), 2*DEBUG_PADDING, 6*DEBUG_PADDING+4*DEBUG_FONTSIZE_SCREEN, DEBUG_FONTSIZE_SCREEN, rl.White)
	rl.DrawText(fmt.Sprintf("> Tile %.0f %.0f", mouseTilePos.X, mouseTilePos.Y), 2*DEBUG_PADDING, 7*DEBUG_PADDING+5*DEBUG_FONTSIZE_SCREEN, DEBUG_FONTSIZE_SCREEN, rl.White)
	rl.DrawText(fmt.Sprintf("> %s, %s", gameMap.Tiles[int(mouseTilePos.Y)][int(mouseTilePos.X)].Name, gameMap.Placeables[int(mouseTilePos.Y)][int(mouseTilePos.X)].Name), 2*DEBUG_PADDING, 8*DEBUG_PADDING+6*DEBUG_FONTSIZE_SCREEN, DEBUG_FONTSIZE_SCREEN, rl.White)

	rl.BeginMode2D(currentCam)

	rl.DrawRectangleLines(int32(mouseTilePos.X*TILE_SIZE-0.5*TILE_SIZE), int32(mouseTilePos.Y*TILE_SIZE-0.5*TILE_SIZE), int32(TILE_SIZE), int32(TILE_SIZE), rl.Blue)
	rl.DrawText(fmt.Sprintf("%.0f, %.0f", mouseTilePos.X, mouseTilePos.Y), int32(mouseTilePos.X*TILE_SIZE-0.5*TILE_SIZE), int32(mouseTilePos.Y*TILE_SIZE-0.5*TILE_SIZE), DEBUG_FONTSIZE_TILE, rl.White)
	rl.DrawRectangleLines(int32(mouseTilePos.X*TILE_SIZE+0.5*TILE_SIZE), int32(mouseTilePos.Y*TILE_SIZE-0.5*TILE_SIZE), int32(TILE_SIZE), int32(TILE_SIZE), rl.Blue)
	rl.DrawText(fmt.Sprintf("%.0f, %.0f", mouseTilePos.X+1, mouseTilePos.Y), int32(mouseTilePos.X*TILE_SIZE+0.5*TILE_SIZE), int32(mouseTilePos.Y*TILE_SIZE-0.5*TILE_SIZE), DEBUG_FONTSIZE_TILE, rl.White)
	rl.DrawRectangleLines(int32(mouseTilePos.X*TILE_SIZE-0.5*TILE_SIZE), int32(mouseTilePos.Y*TILE_SIZE+0.5*TILE_SIZE), int32(TILE_SIZE), int32(TILE_SIZE), rl.Blue)
	rl.DrawText(fmt.Sprintf("%.0f, %.0f", mouseTilePos.X, mouseTilePos.Y+1), int32(mouseTilePos.X*TILE_SIZE-0.5*TILE_SIZE), int32(mouseTilePos.Y*TILE_SIZE+0.5*TILE_SIZE), DEBUG_FONTSIZE_TILE, rl.White)
	rl.DrawRectangleLines(int32(mouseTilePos.X*TILE_SIZE+0.5*TILE_SIZE), int32(mouseTilePos.Y*TILE_SIZE+0.5*TILE_SIZE), int32(TILE_SIZE), int32(TILE_SIZE), rl.Blue)
	rl.DrawText(fmt.Sprintf("%.0f, %.0f", mouseTilePos.X+1, mouseTilePos.Y+1), int32(mouseTilePos.X*TILE_SIZE+0.5*TILE_SIZE), int32(mouseTilePos.Y*TILE_SIZE+0.5*TILE_SIZE), DEBUG_FONTSIZE_TILE, rl.White)

	rl.DrawRectangleLines(int32(mouseTilePos.X*TILE_SIZE), int32(mouseTilePos.Y*TILE_SIZE), int32(TILE_SIZE), int32(TILE_SIZE), rl.Red)

	rl.EndMode2D()
}

func UpdateMainMenu() {
	defaultMenuFontsize = int32(rl.GetRenderHeight()) / 20
	mousePos := rl.GetMousePosition()

	if rl.IsWindowResized() {
		for button := range saveButtons {
			saveButtons[button].Fontsize = defaultMenuFontsize * 20 / 28
		}
		for button := range welcomeButtons {
			welcomeButtons[button].Fontsize = defaultMenuFontsize
		}
		for button := range manageWorldsButtons {
			manageWorldsButtons[button].Fontsize = defaultMenuFontsize
		}
		for button := range optionsButtons {
			optionsButtons[button].Fontsize = defaultMenuFontsize
		}
		for button := range deleteWorldButtons {
			deleteWorldButtons[button].Fontsize = defaultMenuFontsize
		}
		for button := range createWorldButtons {
			createWorldButtons[button].Fontsize = defaultMenuFontsize
		}
	}

	for button := range welcomeButtons {
		if rl.IsWindowResized() {
			length := rl.MeasureText(welcomeButtons[button].Text, welcomeButtons[button].Fontsize)
			welcomeButtons[button].Hitbox.Width = float32(length)
			welcomeButtons[button].Hitbox.Height = float32(welcomeButtons[button].Fontsize)
			welcomeButtons[button].Hitbox.X = float32(int32(rl.GetRenderWidth()/2) - length/2)
			welcomeButtons[button].Hitbox.Y = float32(rl.GetRenderHeight()/3 + button*rl.GetRenderHeight()/12)
		}

		if rl.CheckCollisionPointRec(mousePos, welcomeButtons[button].Hitbox) {
			welcomeButtons[button].Hovered = true
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && welcomeButtons[button].Available {
				welcomeButtons[button].Click()
			}
		} else {
			welcomeButtons[button].Hovered = false
		}
	}

	for button := range manageWorldsButtons {
		if rl.IsWindowResized() {
			length := rl.MeasureText(manageWorldsButtons[button].Text, manageWorldsButtons[button].Fontsize)
			manageWorldsButtons[button].Hitbox.Width = float32(length)
			manageWorldsButtons[button].Hitbox.Height = float32(manageWorldsButtons[button].Fontsize)
			manageWorldsButtons[button].Hitbox.X = 0.15*float32(rl.GetRenderWidth())*float32(1+button%2) + 0.275*float32(button%2*rl.GetRenderWidth())
			manageWorldsButtons[button].Hitbox.Y = 0.65*float32(rl.GetRenderHeight()) + 0.05*float32(rl.GetRenderHeight())*float32(1+button/2) + float32(button/2*int(defaultMenuFontsize))
		}

		if rl.CheckCollisionPointRec(mousePos, manageWorldsButtons[button].Hitbox) {
			manageWorldsButtons[button].Hovered = true
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && manageWorldsButtons[button].Available {
				manageWorldsButtons[button].Click()
			}
		} else {
			manageWorldsButtons[button].Hovered = false
		}
	}

	for button := range saveButtons {
		if rl.IsWindowResized() {
			saveButtons[button].Hitbox.Width = float32(0.6 * float32(rl.GetRenderWidth()))
			saveButtons[button].Hitbox.Height = float32(rl.GetRenderHeight() / 28)
			saveButtons[button].Hitbox.X = 0.15*float32(rl.GetRenderWidth()) + 0.05*float32(rl.GetRenderWidth())
			saveButtons[button].Hitbox.Y = 0.1 * float32(rl.GetRenderHeight()) * float32(1+button)
		}

		if rl.CheckCollisionPointRec(mousePos, saveButtons[button].Hitbox) {
			saveButtons[button].Hovered = true
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && saveButtons[button].Available {
				saveButtons[button].Click()
				selectedSaveId = button
			}
		} else {
			saveButtons[button].Hovered = false
		}
	}

	for button := range createWorldButtons {
		length := rl.MeasureText(createWorldButtons[button].Text, createWorldButtons[button].Fontsize)
		if rl.IsWindowResized() {
			createWorldButtons[button].Hitbox.Width = float32(length)
			createWorldButtons[button].Hitbox.Height = float32(createWorldButtons[button].Fontsize)
			createWorldButtons[button].Hitbox.X = 0.3 * float32(rl.GetRenderWidth()) * float32(1+button%2)
			createWorldButtons[button].Hitbox.Y = 0.5*float32(rl.GetRenderHeight()) + 0.05*float32(rl.GetRenderHeight())*float32(1+button/2) + float32(button/2*int(defaultMenuFontsize))
		}

		if rl.CheckCollisionPointRec(mousePos, createWorldButtons[button].Hitbox) {
			createWorldButtons[button].Hovered = true
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && createWorldButtons[button].Available {
				createWorldButtons[button].Click()
				selectedSaveId = button
			}
		} else {
			createWorldButtons[button].Hovered = false
		}
	}

	for button := range deleteWorldButtons {
		length := rl.MeasureText(deleteWorldButtons[button].Text, deleteWorldButtons[button].Fontsize)
		if rl.IsWindowResized() {
			deleteWorldButtons[button].Hitbox.Width = float32(length)
			deleteWorldButtons[button].Hitbox.Height = float32(deleteWorldButtons[button].Fontsize)
			deleteWorldButtons[button].Hitbox.X = 0.3 * float32(rl.GetRenderWidth()) * float32(1+button%2)
			deleteWorldButtons[button].Hitbox.Y = 0.45*float32(rl.GetRenderHeight()) + 0.05*float32(rl.GetRenderHeight())*float32(1+button/2) + float32(button/2*int(defaultMenuFontsize))
		}

		if rl.CheckCollisionPointRec(mousePos, deleteWorldButtons[button].Hitbox) {
			deleteWorldButtons[button].Hovered = true
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && deleteWorldButtons[button].Available {
				deleteWorldButtons[button].Click()
				selectedSaveId = button
			}
		} else {
			deleteWorldButtons[button].Hovered = false
		}
	}

	if selectedSaveId == -1 {
		manageWorldsButtons[BTN_LOADWORLD].Available = false
		manageWorldsButtons[BTN_DELETEWORLD].Available = false
	} else {
		manageWorldsButtons[BTN_LOADWORLD].Available = true
		manageWorldsButtons[BTN_DELETEWORLD].Available = true
	}

	if menuSection == MANAGE_WORLDS && worldsSection == CREATE_NEW_WORLD {
		HandleTextInput(&createdSave.Name)
	}
}

func DrawMainMenu() {
	switch menuSection {
	case WELCOME:
		{
			rl.DrawText("Farmgame", int32(rl.GetRenderWidth()/2)-rl.MeasureText("Farmgame", int32(rl.GetRenderHeight())/8)/2, int32(rl.GetRenderHeight())/12, int32(float32(defaultMenuFontsize)*2.5), rl.White)
			for button := range welcomeButtons {
				DrawButton(welcomeButtons[button])
			}
		}
	case MANAGE_WORLDS:
		{
			padding := 0.0075 * float32(rl.GetRenderWidth())
			switch worldsSection {
			case MNGWORLD_WELCOME:
				{
					for button := range manageWorldsButtons {
						DrawButton(manageWorldsButtons[button])
					}
					for button := range saveButtons {
						rl.DrawText(saveButtons[button].Text, int32(saveButtons[button].Hitbox.X), int32(saveButtons[button].Hitbox.Y), int32(saveButtons[button].Fontsize), rl.White)
						if button == selectedSaveId {
							rl.DrawRectangleLines(
								int32(saveButtons[button].Hitbox.X-padding),
								int32(saveButtons[button].Hitbox.Y-padding),
								int32(saveButtons[button].Hitbox.Width+2*padding),
								int32(saveButtons[button].Hitbox.Height+2*padding),
								rl.White,
							)
						}
					}
				}
			case CREATE_NEW_WORLD:
				{
					text := "Creating a new world"
					rl.DrawText(text, int32(rl.GetRenderWidth())/2-rl.MeasureText(text, defaultMenuFontsize)/2, int32(0.3*float32(rl.GetRenderHeight())), defaultMenuFontsize, rl.White)
					textBox := rl.Rectangle{}
					textBox.X = 0.3*float32(rl.GetRenderWidth()) - padding
					textBox.Y = 0.425*float32(rl.GetRenderHeight()) - padding
					textBox.Width = createWorldButtons[1].Hitbox.X + createWorldButtons[1].Hitbox.Width - createWorldButtons[0].Hitbox.X + 2*padding
					textBox.Height = float32(defaultMenuFontsize) + 2*padding
					rl.DrawRectangleLinesEx(textBox, 1, rl.White)
					rl.DrawText(createdSave.Name, int32(textBox.X+padding), int32(textBox.Y+padding), defaultMenuFontsize, rl.White)
					for button := range createWorldButtons {
						DrawButton(createWorldButtons[button])
					}
				}
			case DELETE_WORLD:
				{
					text := fmt.Sprintf(`Delete "%s"?`, saves[loadedSaveId].Name)
					rl.DrawText(text, int32(rl.GetRenderWidth())/2-rl.MeasureText(text, defaultMenuFontsize)/2, int32(rl.GetRenderHeight()/3), defaultMenuFontsize, rl.White)
					for button := range deleteWorldButtons {
						DrawButton(deleteWorldButtons[button])
					}
				}
			}
		}
	}
}
