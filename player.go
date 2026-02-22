package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func LoadPlayer() Player {
	var player = Player{
		Nickname:          "Freak",
		ModelName:         "player",
		TimePlayed:        0,
		MovementSpeed:     10,
		MovementState:     IDLE,
		Height:            0.8,
		Width:             0.8,
		WorldPosition:     rl.Vector2{X: float32(gameMap.Width) * TILE_SIZE * 0.5, Y: float32(gameMap.Height) * TILE_SIZE * 0.5},
		Inventory:         [MAX_INVENTORY_SIZE]Item{},
		AvailableInvSlots: MAX_INVENTORY_SIZE,
		SelectedSlot:      0,
		SelectedHotbar:    0,
		Direction:         rl.Vector2{X: 0, Y: 1},
		AnimationTimer:    0.0,
	}

	return player
}

func LoadPlayerAtlases() []rl.Texture2D {
	var playerAtlases = []rl.Texture2D{
		rl.LoadTexture("./assets/models/player/idle.png"),
		rl.LoadTexture("./assets/models/player/walk.png"),
		rl.LoadTexture("./assets/models/player/run.png"),
	}

	return playerAtlases
}

func DrawPlayer(camera *rl.Camera2D) {
	var atlasId int = player.MovementState
	var directionId int                     // vertical
	var frameId int = player.AnimationFrame // horizontal

	switch player.Direction.X {
	case -1:
		{
			directionId = LEFT
			break
		}
	case 1:
		{
			directionId = RIGHT
			break
		}
	default:
		{
			switch player.Direction.Y {
			case -1:
				{
					directionId = BACK
					break
				}
			case 1:
				{
					directionId = FRONT
					break
				}
			default:
				{
					directionId = FRONT
				}
			}
		}
	}

	var source = rl.Rectangle{X: float32(frameId*32) + EPSILON, Y: float32(directionId*32) + EPSILON, Width: 32 - 2*EPSILON, Height: 32 - 2*EPSILON}
	var destination = rl.Rectangle{X: player.WorldPosition.X - TILE_SIZE, Y: player.WorldPosition.Y - 2*TILE_SIZE, Width: 2 * TILE_SIZE, Height: 2 * TILE_SIZE}
	var origin = rl.Vector2{X: 0, Y: 0}

	rl.DrawTexturePro(playerAtlases[atlasId], source, destination, origin, 0.0, rl.White)
}

func HandlePlayerMovement() {
	var movement = rl.Vector2{X: 0, Y: 0}

	if rl.IsKeyDown(rl.KeyW) {
		movement.Y--
	}
	if rl.IsKeyDown(rl.KeyA) {
		movement.X--
	}
	if rl.IsKeyDown(rl.KeyS) {
		movement.Y++
	}
	if rl.IsKeyDown(rl.KeyD) {
		movement.X++
	}

	if movement.X != 0 || movement.Y != 0 {
		player.Direction = movement
		player.MovementState = WALK
	} else {
		player.MovementState = IDLE
	}

	movement = rl.Vector2{X: movement.X * player.MovementSpeed, Y: movement.Y * player.MovementSpeed}

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

	if tile.Id == SOIL && itemName == "Grass Seeds" {
		tileReplacement = tiles[GRASS]
		tileReplaced = true
	}
	if tile.Id == GRASS && itemCategory == HOE {
		tileReplacement = tiles[SOIL]
		tileReplaced = true
	}
	if tile.Id == GRASS && itemName == "Tall Grass Starter" {
		placeableReplacement = placeables[TALL_GRASS]
		placeableReplaced = true
	}
	if (tile.Id == GRASS || tile.Id == SOIL) && itemName == "Sapling" {
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

	if itemName == "Potato Crop" && tile.Id == SOIL {
		placeableReplacement = placeables[POTATO_CROP]
		placeableReplaced = true
	}

	if itemName == "Carrot Crop" && tile.Id == SOIL {
		placeableReplacement = placeables[CARROT_CROP]
		placeableReplaced = true
	}

	if itemName == "Wheat Crop" && tile.Id == SOIL {
		placeableReplacement = placeables[WHEAT_CROP]
		placeableReplaced = true
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
