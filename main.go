package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var gamepadAxisCount int

var meshAtlases []rl.Texture2D
var atlasNames []string
var tileAtlas rl.Texture2D
var itemAtlas rl.Texture2D
var HUDAtlas rl.Texture2D
var placeableAtlas rl.Texture2D

var tiles []Tile
var items []Item
var placeables []Placeable

var gameMap Map
var meshTileMaps [][][]int
var seedMap [][]int
var mapName string

var player Player
var playerCam rl.Camera2D
var freeCam rl.Camera2D
var currentCam rl.Camera2D

var worldMousePos rl.Vector2
var playerCamOn bool
var debugMode bool
var mouseTilePos rl.Vector2

var clientState int
var menuSection int
var worldsSection int
var optionsSection int
var defaultMenuFontsize int32

var welcomeButtons []Button
var manageWorldsButtons []Button
var optionsButtons []Button
var saveButtons []Button
var deleteWorldButtons []Button
var createWorldButtons []Button

var selectedSaveId int
var loadedSaveId int

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint + rl.FlagVsyncHint + rl.FlagWindowResizable)
	rl.InitWindow(400, 300, "Farm Game")
	rl.MaximizeWindow()

	// Main menu
	LoadButtons()
	selectedSaveId = -1
	worldsSection = -1

	// Controller
	gamepadAxisCount = int(rl.GetGamepadAxisCount(0))

	for !rl.WindowShouldClose() {
		if clientState == MAIN_MENU {
			UpdateMainMenu()
			rl.BeginDrawing()

			rl.ClearBackground(rl.Black)
			DrawMainMenu()

			rl.EndDrawing()
		}
		if clientState == LOADING_TO_WORLD || clientState == IN_A_WORLD {
			// Load assets and map
			if clientState == LOADING_TO_WORLD {
				// Assets
				tileAtlas = rl.LoadTexture("./assets/textures/tile_atlas.png")
				itemAtlas = rl.LoadTexture("./assets/textures/item_atlas.png")
				HUDAtlas = rl.LoadTexture("./assets/textures/hud_atlas.png")
				placeableAtlas = rl.LoadTexture("./assets/textures/placeable_atlas.png")
				atlasNames = []string{"soil", "stone", "grass", "bridge"}
				meshAtlases = LoadMeshAtlases()
				tiles = LoadTiles()
				items = LoadItems()
				placeables = loadPlaceables()

				playerCamOn = true
				debugMode = false
				mapName = "map1"

				// Map
				// GenerateMapFromImage(mapName, saveName)
				gameMap = LoadMapFromJSON(saveButtons[loadedSaveId].Text)
				// GenerateSeedMap(gameMap.Width, gameMap.Height, 0, saveName)
				// seedMap = LoadSeedMapFromJSON(saveName)
				meshTileMaps = GenerateMeshTileMaps()

				// Player and Camera
				player = LoadPlayer()
				LoadInventory()
				playerCam = LoadPlayerCam()
				freeCam = LoadFreeCam()

				clientState = IN_A_WORLD
			}

			if clientState == IN_A_WORLD {
				worldMousePos = rl.GetScreenToWorld2D(rl.GetMousePosition(), currentCam)
				mouseTilePos = GetTilePos(worldMousePos)

				if rl.IsKeyPressed(rl.KeyF5) {
					playerCamOn = !playerCamOn
				}
				if rl.IsKeyPressed(rl.KeyF9) {
					debugMode = !debugMode
				}

				if playerCamOn {
					currentCam = playerCam
				} else {
					currentCam = freeCam
				}

				HandlePlayerMovement()
				HandleCamera(&playerCam)
				HandleInventory()

				if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
					InteractWithTile(int(mouseTilePos.Y), int(mouseTilePos.X))
					UpdateMeshTileMaps(rl.Rectangle{X: mouseTilePos.X, Y: mouseTilePos.Y, Width: 2, Height: 2})
				}

				rl.BeginDrawing()

				rl.BeginMode2D(currentCam)

				rl.ClearBackground(rl.Black)
				DrawMap(&playerCam)
				DrawMeshTileMaps(&playerCam)
				HighlightTile(&currentCam)
				DrawPlaceablesAndPlayer(&currentCam)

				rl.EndMode2D()

				DrawHUD()
				if debugMode {
					DrawDebug()
				}

				rl.EndDrawing()
			}
		}
	}
	if clientState == IN_A_WORLD {
		SaveMap(saveButtons[loadedSaveId].Text)
	}
	rl.CloseWindow()
}
