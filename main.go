package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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
var deleteWorldButtons []Button
var createWorldButtons []Button

var saves []Save
var selectedSaveId int
var loadedSaveId int
var createdSave Save

var timeOfTheDay int64
var ambientPlaying bool

var ambient rl.Music
var click rl.Sound

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint + rl.FlagVsyncHint + rl.FlagWindowResizable)
	rl.InitWindow(400, 300, "Farm Game")
	rl.InitAudioDevice()
	rl.SetExitKey(0)
	rl.MaximizeWindow()

	// Main menu
	LoadButtons()
	selectedSaveId = -1
	worldsSection = -1

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

	background := rl.LoadTexture("./assets/textures/background.png")

	click = rl.LoadSound("./assets/sounds/ui/UIClick_UI Click 33_CB Sounddesign_ACTIVATION2.ogg")
	click.Stream.Buffer.Volume = 0.1

	ambient = rl.LoadMusicStream("./assets/sounds/ambient/AMBRurl_Meadow Open Plane Windy Deep Rumble_SYSO_SYSO011-1.ogg")
	ambient.Stream.Buffer.Looping = true
	ambient.Stream.Buffer.Volume = 0.4

	for !rl.WindowShouldClose() {
		createdSave.MapName = "map"
		if clientState == MAIN_MENU {
			UpdateMainMenu()
			background.Width = int32(rl.GetRenderWidth())
			background.Height = int32(rl.GetRenderHeight())
			rl.BeginDrawing()

			rl.ClearBackground(rl.Black)
			if menuSection == WELCOME {
				rl.DrawTexture(background, 0, 0, rl.White)
			}

			DrawMainMenu()

			rl.EndDrawing()
		}
		if clientState == LOADING_TO_WORLD || clientState == IN_A_WORLD {
			// Load assets and map
			if clientState == LOADING_TO_WORLD {

				playerCamOn = true
				debugMode = false
				mapName = "map1"

				// Map
				gameMap = LoadMapFromJSON(saves[loadedSaveId].Name)
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

				rl.UpdateMusicStream(ambient)
				if !ambientPlaying {
					rl.PlayMusicStream(ambient)
					ambientPlaying = true
				}

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
		SaveMapJSON(saves[loadedSaveId].Name)
		UpdateSaveMetadata(saves[loadedSaveId])
	}
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
