package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Tile struct {
	Name      string
	TextureId int
	HasGrass  int
	IsWet     int
}

type Placeable struct {
	Name      string
	TextureId int
	AtlasX    int
	AtlasY    int
	Width     int
	Height    int
}

type Item struct {
	Name       string
	CategoryId int
	TextureId  int
}

type Player struct {
	Nickname          string
	ModelName         string
	TimePlayed        float64
	MovementSpeed     float32
	Height            float32
	Width             float32
	WorldPosition     rl.Vector2
	Inventory         [MAX_INVENTORY_SIZE]Item
	AvailableInvSlots int
	SelectedSlot      int
	SelectedHotbar    int
}

type Map struct {
	Tiles      [][]Tile
	Placeables [][]Placeable
	Width      int
	Height     int
}

type Entity struct {
	PosX   float32
	PosY   float32
	VelX   float32
	VelY   float32
	Width  float32
	Height float32
}

type Button struct {
	Text      string
	Fontsize  int32
	Click     func()
	Hitbox    rl.Rectangle
	Hovered   bool
	Available bool
}

type Save struct {
	Name    string
	MapName string
}
