package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	Id       int16
	HasGrass int16
	IsWet    int16
}

type Placeable struct {
	Id     int16
	AtlasX int16
	AtlasY int16
	Width  int16
	Height int16
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
	Name        string
	MapName     string
	LastPlayed  time.Time
	DisplayTime string
	MenuButton  Button
}
