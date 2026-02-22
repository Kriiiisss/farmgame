package main

import rl "github.com/gen2brain/raylib-go/raylib"

func LoadPlaceables() []Placeable {
	var placeables = []Placeable{
		{
			Id:     0,
			AtlasX: 0,
			AtlasY: 0,
			Width:  0,
			Height: 0,
		},
		{
			Id:     TALL_GRASS,
			AtlasX: 160,
			AtlasY: 0,
			Width:  32,
			Height: 64,
		},
		{
			Id:     TREE,
			AtlasX: 0,
			AtlasY: 0,
			Width:  160,
			Height: 192,
		},
		{
			Id:     POTATO_CROP,
			AtlasX: 192,
			AtlasY: 0,
			Width:  32,
			Height: 32,
		},
		{
			Id:     CARROT_CROP,
			AtlasX: 224,
			AtlasY: 0,
			Width:  32,
			Height: 32,
		},
		{
			Id:     WHEAT_CROP,
			AtlasX: 256,
			AtlasY: 0,
			Width:  32,
			Height: 32,
		},
	}

	return placeables
}

func DrawPlaceablesAndPlayer(camera *rl.Camera2D) {
	var tint rl.Color
	var source rl.Rectangle
	var destination rl.Rectangle
	var origin = rl.Vector2{X: 0, Y: 0}
	cameraTopLeft := GetTilePos(rl.GetScreenToWorld2D(rl.Vector2{X: 0, Y: 0}, *camera))
	cameraBottomRight := GetTilePos(rl.GetScreenToWorld2D(rl.Vector2{X: float32(rl.GetRenderWidth() - 1), Y: float32(rl.GetRenderHeight() - 1)}, *camera))
	playerY := GetTilePos(player.WorldPosition).Y

	for y := Clamp(cameraTopLeft.Y-PLACEABLES_RENDER_TOLERANCE, 0, float32(gameMap.Height)); y < Clamp(cameraBottomRight.Y+PLACEABLES_RENDER_TOLERANCE, 0, float32(gameMap.Height)); y++ {
		if y == playerY {
			DrawPlayer(camera)
		}
		for x := Clamp(cameraTopLeft.X-PLACEABLES_RENDER_TOLERANCE, 0, float32(gameMap.Width)); x < Clamp(cameraBottomRight.X+PLACEABLES_RENDER_TOLERANCE, 0, float32(gameMap.Width)); x++ {
			placeable := gameMap.Placeables[int(y)][int(x)]
			if placeableNames[placeable.Id] != "" {
				switch placeable.Id {
				case TALL_GRASS:
					tint = rl.GetColor(GRASS_TINT)
				default:
					tint = rl.White
				}
				source = rl.Rectangle{X: float32(placeable.AtlasX) + EPSILON, Y: float32(placeable.AtlasY) + EPSILON, Width: float32(placeable.Width) - 2*EPSILON, Height: float32(placeable.Height) - 2*EPSILON}
				destination = rl.Rectangle{X: x*TILE_SIZE - (float32(placeable.Width)-TILE_SIZE)*0.5, Y: y*TILE_SIZE - float32(placeable.Height) + TILE_SIZE, Width: float32(placeable.Width), Height: float32(placeable.Height)}
				rl.DrawTexturePro(placeableAtlas, source, destination, origin, 0, tint)
			}
		}
	}
}
