package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func LoadPlayer() Player {
	var player = Player{
		"Freak",
		"player",
		0,
		10,
		1.8,
		0.8,
		rl.Vector2{X: float32(gameMap.Width) * TILE_SIZE * 0.5, Y: float32(gameMap.Height) * TILE_SIZE * 0.5},
		[MAX_INVENTORY_SIZE]Item{},
		MAX_INVENTORY_SIZE,
		0,
		0,
	}

	return player
}

func DrawPlaceablesAndPlayer(camera *rl.Camera2D) {
	var tint rl.Color
	var source rl.Rectangle
	var destination rl.Rectangle
	var origin = rl.Vector2{X: 0, Y: 0}
	cameraTopLeft := GetTilePos(rl.GetScreenToWorld2D(rl.Vector2{X: 0, Y: 0}, *camera))
	cameraBottomRight := GetTilePos(rl.GetScreenToWorld2D(rl.Vector2{X: float32(rl.GetRenderWidth() - 1), Y: float32(rl.GetRenderHeight() - 1)}, *camera))
	playerY := GetTilePos(player.WorldPosition).Y

	// Draw player as a black rectangle
	rl.DrawRectangle(int32(player.WorldPosition.X-player.Width*0.5*TILE_SIZE), int32(player.WorldPosition.Y-player.Height*TILE_SIZE), int32(player.Width*TILE_SIZE), int32(player.Height*TILE_SIZE), rl.Black)

	// Draw placeables
	for y := Clamp(cameraTopLeft.Y-PLACEABLES_RENDER_TOLERANCE, 0, float32(gameMap.Height)); y < Clamp(cameraBottomRight.Y+PLACEABLES_RENDER_TOLERANCE, 0, float32(gameMap.Height)); y++ {
		if y == playerY {
			rl.DrawRectangle(int32(player.WorldPosition.X-player.Width*0.5*TILE_SIZE), int32(player.WorldPosition.Y-player.Height*TILE_SIZE), int32(player.Width*TILE_SIZE), int32(player.Height*TILE_SIZE), rl.Black)
		}
		for x := Clamp(cameraTopLeft.X-PLACEABLES_RENDER_TOLERANCE, 0, float32(gameMap.Width)); x < Clamp(cameraBottomRight.X+PLACEABLES_RENDER_TOLERANCE, 0, float32(gameMap.Width)); x++ {
			placeable := gameMap.Placeables[int(y)][int(x)]
			if placeable.Name != "" {
				switch placeable.TextureId {
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

func HandlePlayerMovement() {
	var movement = rl.Vector2{X: 0, Y: 0}

	if rl.IsKeyDown(rl.KeyW) {
		movement.Y -= player.MovementSpeed
	}
	if rl.IsKeyDown(rl.KeyA) {
		movement.X -= player.MovementSpeed
	}
	if rl.IsKeyDown(rl.KeyS) {
		movement.Y += player.MovementSpeed
	}
	if rl.IsKeyDown(rl.KeyD) {
		movement.X += player.MovementSpeed
	}

	if rl.IsGamepadAvailable(0) {

	}

	if movement.X != 0 && movement.Y != 0 {
		movement.X = movement.X / SQRT2
		movement.Y = movement.Y / SQRT2
	}

	player.WorldPosition.X += movement.X * TILE_SIZE * rl.GetFrameTime()
	player.WorldPosition.Y += movement.Y * TILE_SIZE * rl.GetFrameTime()
}

func InteractWithTile(x, y int) {
	if player.SelectedSlot == -1 {
		return
	}

	var tileReplacement Tile
	var placeableReplacement Placeable
	var tile = &gameMap.Tiles[int(mouseTilePos.Y)][int(mouseTilePos.X)]
	var placeable = &gameMap.Placeables[int(mouseTilePos.Y)][int(mouseTilePos.X)]
	tileReplaced := false
	placeableReplaced := false
	itemCategory := player.Inventory[player.SelectedSlot].CategoryId
	itemName := player.Inventory[player.SelectedSlot].Name

	if itemName == "" {
		return
	}

	if tile.Name == "Soil" && itemName == "Grass Seeds" {
		tileReplacement = tiles[GRASS]
		tileReplaced = true
	}
	if tile.Name == "Grass" && itemCategory == HOE {
		tileReplacement = tiles[SOIL]
		tileReplaced = true
	}
	if tile.Name == "Grass" && itemName == "Tall Grass Starter" {
		placeableReplacement = placeables[TALL_GRASS]
		placeableReplaced = true
	}
	if (tile.Name == "Grass" || tile.Name == "Soil") && itemName == "Sapling" {
		placeableReplacement = placeables[TREE]
		placeableReplaced = true
	}
	if itemName == "Delete" {
		placeableReplacement = placeables[NONE]
		placeableReplaced = true
	}
	if itemName == "Grass Tile" {
		tileReplacement = tiles[GRASS]
		tileReplaced = true
	}
	if itemName == "Soil Tile" {
		tileReplacement = tiles[SOIL]
		tileReplaced = true
	}
	if itemName == "Water Tile" {
		tileReplacement = tiles[WATER]
		tileReplaced = true
	}
	if itemName == "Stone Tile" {
		tileReplacement = tiles[STONE]
		tileReplaced = true
	}
	if itemName == "Bridge Tile" {
		tileReplacement = tiles[BRIDGE]
		tileReplaced = true
	}

	if tileReplaced {
		*tile = tileReplacement
	}
	if placeableReplaced {
		*placeable = placeableReplacement
	}
}

func HandleInventory() {
	keyPressed := int(rl.GetKeyPressed())
	if keyPressed == rl.KeyQ || keyPressed == rl.KeyE {
		availableHotbars := int(math.Ceil(float64(player.AvailableInvSlots) / 9))
		hotbarSwitchDirection := 1
		if keyPressed == rl.KeyQ {
			hotbarSwitchDirection = -1
		}
		slotIndex := player.SelectedSlot % 9
		player.SelectedHotbar += hotbarSwitchDirection
		if player.SelectedHotbar < 0 {
			player.SelectedHotbar = availableHotbars - 1
		}
		if player.SelectedHotbar >= availableHotbars {
			player.SelectedHotbar = 0
		}
		if player.SelectedSlot != -1 {
			player.SelectedSlot = int(Clamp(float32(slotIndex+player.SelectedHotbar*9), 0, float32(player.AvailableInvSlots-1)))
		}
	}
	if keyPressed >= rl.KeyOne && keyPressed <= rl.KeyNine {
		if player.SelectedSlot == keyPressed-rl.KeyOne+9*player.SelectedHotbar {
			player.SelectedSlot = -1
		} else {
			player.SelectedSlot = int(Clamp(float32((keyPressed-rl.KeyOne)+9*player.SelectedHotbar), 0, float32(player.AvailableInvSlots-1)))
		}
	}
}

func DrawHUD() {
	renderWidth := float32(rl.GetRenderWidth())
	renderHeight := float32(rl.GetRenderHeight())
	hotbarTopLeft := rl.Vector2{X: renderWidth/2 - 4.5*HOTBAR_SLOT_SIZE*renderWidth - 4.0*HOTBAR_SLOT_PADDING*renderWidth, Y: renderHeight - HOTBAR_SLOT_SIZE*renderWidth - HOTBAR_SLOT_PADDING*renderWidth}
	var textureOffset float32
	var textureId int
	var source rl.Rectangle
	var destination rl.Rectangle
	var origin = rl.Vector2{X: 0, Y: 0}

	for i := range 9 {
		slotTexture := SLOT_TEXTURE_ID
		if player.SelectedHotbar*9+i == player.SelectedSlot {
			slotTexture = SLOT_SEL_TEXTURE_ID
		} else {
			if player.SelectedHotbar*9+i >= player.AvailableInvSlots {
				slotTexture = SLOT_UNAV_TEXTURE_ID
			}
		}
		source = rl.Rectangle{X: (float32(slotTexture-SLOT_TEXTURE_ID)*20 + 8), Y: 0, Width: 20, Height: 20}
		destination = rl.Rectangle{
			X:      hotbarTopLeft.X + float32(i)*renderWidth*(HOTBAR_SLOT_SIZE+HOTBAR_SLOT_PADDING),
			Y:      hotbarTopLeft.Y,
			Width:  HOTBAR_SLOT_SIZE * renderWidth,
			Height: HOTBAR_SLOT_SIZE * renderWidth,
		}
		rl.DrawTexturePro(HUDAtlas, source, destination, origin, 0, rl.White)
		if player.SelectedHotbar*9+i < player.AvailableInvSlots && player.Inventory[player.SelectedHotbar*9+i].Name != "" {
			textureId = player.Inventory[player.SelectedHotbar*9+i].TextureId
			if player.SelectedHotbar*9+i != player.SelectedSlot {
				textureOffset = 0.15 * HOTBAR_SLOT_SIZE * renderWidth
			} else {
				textureOffset = 0.05 * HOTBAR_SLOT_SIZE * renderWidth
			}
			source = rl.Rectangle{X: float32(textureId*16) + EPSILON, Y: EPSILON, Width: 16 - 2*EPSILON, Height: 16 - 2*EPSILON}
			destination = rl.Rectangle{
				X:      hotbarTopLeft.X + float32(i)*renderWidth*(HOTBAR_SLOT_SIZE+HOTBAR_SLOT_PADDING) + textureOffset,
				Y:      hotbarTopLeft.Y + textureOffset,
				Width:  HOTBAR_SLOT_SIZE*renderWidth - 2*textureOffset,
				Height: HOTBAR_SLOT_SIZE*renderWidth - 2*textureOffset,
			}
			rl.DrawTexturePro(itemAtlas, source, destination, origin, 0, rl.White)
		}
	}

	for i := range int(math.Ceil(float64(player.AvailableInvSlots) / 9)) {
		indexTexture := INDEX_TEXTURE_ID
		if i == player.SelectedHotbar {
			indexTexture = INDEX_SEL_TEXTURE_ID
		}
		x := hotbarTopLeft.X - 2*HOTBAR_INDEX_SIZE*renderWidth - float32(i/3*2)*HOTBAR_INDEX_SIZE*renderWidth
		y := hotbarTopLeft.Y + HOTBAR_INDEX_SIZE*renderWidth + float32(i%3*2)*HOTBAR_INDEX_SIZE*renderWidth
		source := rl.Rectangle{X: float32(indexTexture) * 4, Y: 0, Width: 4, Height: 4}
		destination = rl.Rectangle{X: x, Y: y, Width: HOTBAR_INDEX_SIZE * renderWidth, Height: HOTBAR_INDEX_SIZE * renderWidth}
		rl.DrawTexturePro(HUDAtlas, source, destination, origin, 0, rl.White)
	}
}
