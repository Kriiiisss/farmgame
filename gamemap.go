package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GetTilePos(worldPosition rl.Vector2) rl.Vector2 {
	return rl.Vector2{X: Clamp(float32(Floor(worldPosition.X/TILE_SIZE)), 0, float32(gameMap.Width-1)), Y: Clamp(float32(Floor(worldPosition.Y/TILE_SIZE)), 0, float32(gameMap.Height-1))}
}

func GenerateSaveFiles(save Save) {
	image := rl.LoadImage("./assets/map images/" + save.MapName + ".png")

	savePath := "./saves/" + save.Name
	tilesPath := savePath + "/map.json"
	placeablesPath := savePath + "/placeables.json"

	var loadedTiles [][]Tile
	var loadedPlaceables [][]Placeable

	for y := range image.Height {
		loadedTiles = append(loadedTiles, []Tile{})
		loadedPlaceables = append(loadedPlaceables, []Placeable{})
		for x := range image.Width {
			loadedPlaceables[y] = append(loadedPlaceables[y], placeables[NONE])
			loadedTiles[y] = append(loadedTiles[y], tiles[hexToTileId[uint32(rl.ColorToInt(rl.GetImageColor(*image, x, y)))]])
		}
	}

	jsonTiles, err := json.Marshal(loadedTiles)
	if err != nil {
		log.Fatal(err)
	}
	jsonPlaceables, err := json.Marshal(loadedPlaceables)
	if err != nil {
		log.Fatal(err)
	}

	f1, err := os.Create(tilesPath)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(tilesPath, jsonTiles, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	f2, err := os.Create(placeablesPath)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(placeablesPath, jsonPlaceables, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	UpdateSaveMetadata(save)

	fmt.Print("Generated save files for " + `"` + save.Name + `"` + ".\n")

	rl.UnloadImage(image)
	f1.Close()
	f2.Close()
}

func GenerateSeedMap(gameMapWidth, gameMapHeight, seed int, saveName string) {
	// ofstream seedMapFile("../saves/" + saveName + "/seed_map.json");
	// json seedMap;

	// for y := range gameMapHeight {
	// seedMap[y] = {};
	// for x := range gameMapWidth {
	// 	seedMap[y][x] = rand.Int() % 8
	// }
	// }

	fmt.Printf("Generated Seed Map!\n")

	// seedMapFile << seedMap;
	// seedMapFile.close();
}

func LoadMeshAtlases() []rl.Texture2D {
	for atlasId := range MESH_COUNT {
		meshAtlases = append(meshAtlases, rl.LoadTexture("./assets/textures/"+atlasNames[atlasId]+"_atlas.png"))
	}

	fmt.Printf("Loaded Mesh Atlases!\n")

	return meshAtlases
}

func LoadMapFromJSON(saveName string) Map {
	inputTilesPath := "./saves/" + saveName + "/map.json"
	inputPlaceablesPath := "./saves/" + saveName + "/placeables.json"

	var output Map

	jsonTiles, err := os.ReadFile(inputTilesPath)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(jsonTiles, &output.Tiles)
	if err != nil {
		fmt.Println(err)
	}

	jsonPlaceables, err := os.ReadFile(inputPlaceablesPath)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(jsonPlaceables, &output.Placeables)
	if err != nil {
		fmt.Println(err)
	}

	output.Height = len(output.Tiles)
	output.Width = len(output.Tiles[0])

	fmt.Printf("Loaded Tile Map!\n")

	return output
}

func LoadSeedMapFromJSON(saveName string) [][]int {
	var output [][]int
	// ifstream seedMapFile("../saves/" + saveName + "/seed_map.json");
	// json seedMap;
	// seedMapFile >> seedMap;
	// seedMapFile.close();

	// for y := range seedMap.size() {
	//     output.emplace_back();
	//     for x := range seedMap[0].size() {
	//         output[y].push_back(seedMap[y][x]);
	//     }
	// }

	fmt.Printf("Loaded Seed Map!\n")

	return output
}

func GenerateMeshTileMaps() [][][]int {
	var meshTileMaps [][][]int
	var data [MESH_COUNT][9]int
	atlasId := 0

	var tilePositions = [9][2]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{0, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}
	var meshTilePositions = [4][2]int{
		{0, 0},
		{1, 0},
		{1, 1},
		{0, 1},
	}
	var cornerPositions = [4][4]int{
		{0, 1, 4, 3},
		{1, 2, 5, 4},
		{4, 5, 8, 7},
		{3, 4, 7, 6},
	}
	for mapId := range MESH_COUNT {
		meshTileMaps = append(meshTileMaps, [][]int{})
		for y := range gameMap.Height + 1 {
			meshTileMaps[mapId] = append(meshTileMaps[mapId], []int{})
			for _ = range gameMap.Width + 1 {
				meshTileMaps[mapId][y] = append(meshTileMaps[mapId][y], 0)
			}
		}
	}

	for y := range gameMap.Height {
		for x := range gameMap.Width {
			for i := range 9 {
				if IsInRange(tilePositions[i][0]+x, 0, gameMap.Width-1) && IsInRange(tilePositions[i][1]+y, 0, gameMap.Height-1) {
					isStone := gameMap.Tiles[tilePositions[i][1]+y][tilePositions[i][0]+x].Name == "Stone"
					if isStone {
						data[STONE_MESH][i] = 1
					} else {
						data[STONE_MESH][i] = 0
					}

					isSoil := gameMap.Tiles[tilePositions[i][1]+y][tilePositions[i][0]+x].Name == "Soil" || gameMap.Tiles[tilePositions[i][1]+y][tilePositions[i][0]+x].Name == "Grass"
					if isSoil {
						data[SOIL_MESH][i] = 1
					} else {
						data[SOIL_MESH][i] = 0
					}

					isGrass := gameMap.Tiles[tilePositions[i][1]+y][tilePositions[i][0]+x].Name == "Grass"
					if isGrass {
						data[GRASS_MESH][i] = 1
					} else {
						data[GRASS_MESH][i] = 0
					}

					isBridge := gameMap.Tiles[tilePositions[i][1]+y][tilePositions[i][0]+x].Name == "Bridge"
					if isBridge {
						data[BRIDGE_MESH][i] = 1
					} else {
						data[BRIDGE_MESH][i] = 0
					}
				} else {
					data[STONE_MESH][i] = 0
					data[SOIL_MESH][i] = 0
					data[GRASS_MESH][i] = 0
					data[BRIDGE_MESH][i] = 0
				}
			}
			for mapId := range MESH_COUNT {
				for corner := range 4 {
					for cornerPos := range 4 {
						if data[mapId][cornerPositions[corner][cornerPos]] == 1 {
							atlasId += int(math.Pow(2, float64(cornerPos)))
						}
					}
					meshTileMaps[mapId][y+meshTilePositions[corner][1]][x+meshTilePositions[corner][0]] = atlasId
					atlasId = 0
				}
			}
		}
	}
	fmt.Printf("Generated Mesh Tile Maps!\n")

	return meshTileMaps
}

func UpdateMeshTileMaps(area rl.Rectangle) {
	var data [MESH_COUNT][9]int
	atlasId := 0

	var tilePositions = [9][2]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{0, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}
	var meshTilePositions = [4][2]int{
		{0, 0},
		{1, 0},
		{1, 1},
		{0, 1},
	}
	var cornerPositions = [4][4]int{
		{0, 1, 4, 3},
		{1, 2, 5, 4},
		{4, 5, 8, 7},
		{3, 4, 7, 6},
	}

	for y := area.Y; y < area.Y+area.Height-1; y++ {
		for x := area.X; x < area.X+area.Width-1; x++ {
			for i := range 9 {
				if IsInRange(tilePositions[i][0]+int(x), 0, gameMap.Width-1) && IsInRange(tilePositions[i][1]+int(y), 0, gameMap.Height-1) {
					isStone := (gameMap.Tiles[tilePositions[i][1]+int(y)][tilePositions[i][0]+int(x)].Name == "Stone")
					if isStone {
						data[STONE_MESH][i] = 1
					} else {
						data[STONE_MESH][i] = 0
					}

					isSoil := gameMap.Tiles[tilePositions[i][1]+int(y)][tilePositions[i][0]+int(x)].Name == "Soil" || gameMap.Tiles[tilePositions[i][1]+int(y)][tilePositions[i][0]+int(x)].Name == "Grass"
					if isSoil {
						data[SOIL_MESH][i] = 1
					} else {
						data[SOIL_MESH][i] = 0
					}

					isGrass := gameMap.Tiles[tilePositions[i][1]+int(y)][tilePositions[i][0]+int(x)].Name == "Grass"
					if isGrass {
						data[GRASS_MESH][i] = 1
					} else {
						data[GRASS_MESH][i] = 0
					}

					isBridge := gameMap.Tiles[tilePositions[i][1]+int(y)][tilePositions[i][0]+int(x)].Name == "Bridge"
					if isBridge {
						data[BRIDGE_MESH][i] = 1
					} else {
						data[BRIDGE_MESH][i] = 0
					}
				} else {
					data[STONE_MESH][i] = 0
					data[SOIL_MESH][i] = 0
					data[GRASS_MESH][i] = 0
					data[BRIDGE_MESH][i] = 0
				}
			}
			for mapId := range MESH_COUNT {
				for corner := range 4 {
					for cornerPos := range 4 {
						if data[mapId][cornerPositions[corner][cornerPos]] == 1 {
							atlasId += int(math.Pow(2, float64(cornerPos)))
						}
					}
					meshTileMaps[mapId][int(y)+meshTilePositions[corner][1]][int(x)+meshTilePositions[corner][0]] = atlasId
					atlasId = 0
				}
			}
		}
	}
}

func DrawMap(camera *rl.Camera2D) {
	tint := rl.White
	textureId := 0
	var source rl.Rectangle
	var destination rl.Rectangle
	origin := rl.Vector2{X: 0, Y: 0}
	cameraTopLeft := GetTilePos(rl.GetScreenToWorld2D(rl.Vector2{X: 0, Y: 0}, *camera))
	cameraBottomRight := GetTilePos(rl.GetScreenToWorld2D(rl.Vector2{X: float32(rl.GetRenderWidth() - 1), Y: float32(rl.GetRenderHeight() - 1)}, *camera))

	for y := Clamp(cameraTopLeft.Y-TILES_RENDER_TOLERANCE, 0, float32(gameMap.Height)); y < Clamp(cameraBottomRight.Y+TILES_RENDER_TOLERANCE, 0, float32(gameMap.Height)); y++ {
		for x := Clamp(cameraTopLeft.X-TILES_RENDER_TOLERANCE, 0, float32(gameMap.Width)); x < Clamp(cameraBottomRight.X+TILES_RENDER_TOLERANCE, 0, float32(gameMap.Width)); x++ {
			textureId = gameMap.Tiles[int(y)][int(x)].TextureId
			source = rl.Rectangle{X: float32(textureId*16) + EPSILON, Y: EPSILON, Width: 16 - 2*EPSILON, Height: 16 - 2*EPSILON}
			destination = rl.Rectangle{X: x * TILE_SIZE, Y: y * TILE_SIZE, Width: TILE_SIZE, Height: TILE_SIZE}
			rl.DrawTexturePro(tileAtlas, source, destination, origin, 0, tint)
		}
	}
}

func DrawMeshTileMaps(camera *rl.Camera2D) {
	// Positions of tilemap segments in a human-readable 4x4 atlas
	// Order of these is based on a clockwise binary addition (..._atlas_old.png)
	var atlasPositions = [16][2]int{
		{0, 3},
		{3, 3},
		{0, 2},
		{1, 2},
		{1, 3},
		{0, 1},
		{1, 0},
		{2, 2},
		{0, 0},
		{3, 2},
		{2, 3},
		{3, 1},
		{3, 0},
		{2, 0},
		{1, 1},
		{2, 1},
	}

	tint := rl.White
	var textureId rl.Vector2
	var source rl.Rectangle
	var destination rl.Rectangle
	origin := rl.Vector2{X: 0, Y: 0}
	cameraTopLeft := GetTilePos(rl.GetScreenToWorld2D(rl.Vector2{X: 0, Y: 0}, *camera))
	cameraBottomRight := GetTilePos(rl.GetScreenToWorld2D(rl.Vector2{X: float32(rl.GetRenderWidth() - 1), Y: float32(rl.GetRenderHeight() - 1)}, *camera))

	for y := Clamp(cameraTopLeft.Y-TILES_RENDER_TOLERANCE, 0, float32(gameMap.Height+1)); y < Clamp(cameraBottomRight.Y+TILES_RENDER_TOLERANCE, 0, float32(gameMap.Height+1)); y++ {
		for x := Clamp(cameraTopLeft.X-TILES_RENDER_TOLERANCE, 0, float32(gameMap.Width+1)); x < Clamp(cameraBottomRight.X+TILES_RENDER_TOLERANCE, 0, float32(gameMap.Width+1)); x++ {
			for mapId := range MESH_COUNT {
				textureId = rl.Vector2{X: float32(atlasPositions[meshTileMaps[mapId][int(y)][int(x)]][0]), Y: float32(atlasPositions[meshTileMaps[mapId][int(y)][int(x)]][1])}
				source = rl.Rectangle{X: float32(textureId.X*16) + EPSILON, Y: float32(textureId.Y*16) + EPSILON, Width: 16 - 2*EPSILON, Height: 16 - 2*EPSILON}
				destination = rl.Rectangle{X: (x - 0.5) * TILE_SIZE, Y: (y - 0.5) * TILE_SIZE, Width: TILE_SIZE, Height: TILE_SIZE}
				rl.DrawTexturePro(meshAtlases[mapId], source, destination, origin, 0, tint)
			}
		}
	}
}

func SaveMap(saveName string) {
	tilesPath := "./saves/" + saveName + "/map.json"
	placeablesPath := "./saves/" + saveName + "/placeables.json"

	jsonTiles, err := json.Marshal(gameMap.Tiles)
	if err != nil {
		log.Fatal(err)
	}
	jsonPlaceables, err := json.Marshal(gameMap.Placeables)
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Create(tilesPath)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(tilesPath, jsonTiles, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Create(placeablesPath)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(placeablesPath, jsonPlaceables, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Saved %s!\n", saveName)
}
