package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func LoadTiles() []Tile {
	var tiles = []Tile{
		{
			0,
			0,
			0,
		},
		{
			WATER,
			0,
			0,
		},
		{
			SOIL,
			0,
			0,
		},
		{
			GRASS,
			1,
			0,
		},
		{
			STONE,
			0,
			0,
		},
		{
			BRIDGE,
			0,
			0,
		},
	}

	fmt.Printf("Loaded Tiles!\n")

	return tiles
}

func HighlightTile(currentCam *rl.Camera2D) {
	worldMousePos := rl.GetScreenToWorld2D(rl.GetMousePosition(), *currentCam)
	tileMousePos := GetTilePos(worldMousePos)

	offset := 2 * TILE_SIZE / 16
	var origin = rl.Vector2{X: 0, Y: 0}
	source := rl.Rectangle{X: 68, Y: 0, Width: 20, Height: 20}
	destination := rl.Rectangle{X: tileMousePos.X*TILE_SIZE - offset, Y: tileMousePos.Y*TILE_SIZE - offset, Width: float32(TILE_SIZE + 2*offset), Height: float32(TILE_SIZE + 2*offset)}
	rl.DrawTexturePro(HUDAtlas, source, destination, origin, 0, rl.White)
}
